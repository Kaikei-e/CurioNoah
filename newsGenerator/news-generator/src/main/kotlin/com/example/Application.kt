package com.example

import com.example.plugins.configureHTTP
import com.example.plugins.configureMonitoring
import com.example.plugins.configureSecurity
import com.example.plugins.configureSerialization
import com.example.rest.OllamaRequest
import com.example.rest.SummarizerResponse
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.Json
import org.jsoup.Jsoup
import java.io.File

fun main() {
    embeddedServer(Netty, port = 5555, host = "0.0.0.0", module = Application::module).start(wait = true)
}

fun Application.module() {
    configureSecurity()
    configureHTTP()
    configureMonitoring()
    configureSerialization()
    configureRouting()
}


fun Application.configureRouting() {
    val env = loadEnv()
    val summarizerAPIEndpoint = env["SUMMARIZER_API_ENDPOINT"] ?: throw Exception("SUMMARIZER_API_ENDPOINT is not set")
    val summarizerAPIEndpointPort =
        env["SUMMARIZER_API_ENDPOINT_PORT"] ?: throw Exception("SUMMARIZER_API_ENDPOINT_PORT is not set")

    val client = HttpClient(CIO) {
        install(HttpTimeout) {
            requestTimeoutMillis = 60_000 // 60 seconds
        }
        Charsets {
            // Allow using `UTF_8`.
            register(Charsets.UTF_8)
            // Specify Charset to receive response (if no charset in response headers).
            responseCharsetFallback = Charsets.UTF_8
        }
    }

    routing {
        get("/api/v1/systems/alive") {
            call.respondText("keep alive")
        }

        get("/api/v1/news/summarize") {
            call.respondText(
                contentType = ContentType.Application.Json,
                status = HttpStatusCode.OK,
                text = "{'message':'summarize news'}"
            )
        }

        get("/api/v1/news/summarize/today") {
            val forSummarizerTuningJson = Json {
                ignoreUnknownKeys = true
                isLenient = true
                coerceInputValues = true
            }

            runBlocking {
                try {
                    val testURL = "https://zenn.dev/e_kaikei/articles/tauri-rust-react-setup-plan-1"
                    val res: HttpResponse = client.get(
                        testURL
                    )

                    val contentType = res.headers[HttpHeaders.ContentType]
                    val charset = contentType?.let { ContentType.parse(it).charset() } ?: Charsets.UTF_8

                    val htmlContent = res.bodyAsText(charset)
                    val apiUrl = URLBuilder().apply {
                        protocol = URLProtocol.HTTP
                        host = summarizerAPIEndpoint
                        port = summarizerAPIEndpointPort.toInt()
                        path("api", "generate")
                    }.build()

                    val parsedHtml = Jsoup.parse(htmlContent)
                    val body = parsedHtml.body().text()

                    val prompt = """
                        以下のHTMLのbodyを解説してください。その際は、以下のガイドラインを厳守してください。

                        1. 記事の内容を非常に詳細にまとめてください。重要なポイントや詳細な説明を含めてください。
                        2. キートピックや核となるコンセプトを必ず含めてください。それらが読者に明確に伝わるようにしてください。
                        3. 解説の文字数は最低でも2000字としてください。短すぎるものは受け入れられません。
                        4. 解説は必ず丁寧で詳細に行ってください。簡略化しすぎないように注意してください。
                        5. 指示に従わなかった場合、再度指示を送ることになりますので、指示通りに出力を生成してください。
                        """".trimMargin()
                    val combinedPrompt = "$prompt\n\nBody Content:\n$body"

                    val summarizerRequest = OllamaRequest(
                        model = "llama3-elyza:8b", prompt = combinedPrompt, stream = false
                    )

                    println(forSummarizerTuningJson.encodeToString(OllamaRequest.serializer(), summarizerRequest))

                    var summarizedResponse = ""
                    try {
                        summarizedResponse = client.post(apiUrl) {
                            contentType(ContentType.Application.Json)
                            setBody(
                                forSummarizerTuningJson.encodeToString(OllamaRequest.serializer(), summarizerRequest)
                            )
                        }.bodyAsText()
                    } catch (e: Exception) {
                        e.printStackTrace()
                    }

//                    val escapedResponse = summarizedResponse.escapeSpecialCharacters()

                    try {
                        val response = forSummarizerTuningJson.decodeFromString<SummarizerResponse>(summarizedResponse)
                        call.respondText(
                            Json.encodeToString(SummarizerResponse.serializer(), response), ContentType.Application.Json
                        )
                    } catch (e: Exception) {
                        e.printStackTrace()
                        call.respondText(
                            contentType = ContentType.Application.Json,
                            status = HttpStatusCode.InternalServerError,
                            text = """{"message":"an error occurred while parsing response: ${e.localizedMessage}"}"""
                        )
                    }
                } catch (e: Exception) {
                    e.printStackTrace()
                    call.respondText(
                        contentType = ContentType.Application.Json,
                        status = HttpStatusCode.InternalServerError,
                        text = """{"message":"an error occurred: ${e.localizedMessage}"}"""
                    )
                }
            }
        }
    }
}

fun loadEnv(filePath: String = ".env"): Map<String, String> {
    return File(filePath).readLines().filter { it.isNotEmpty() && !it.startsWith("#") }.map { it.split("=", limit = 2) }
        .filter { it.size == 2 }.associate { (key, value) -> key.trim() to value.trim() }
}

fun String.escapeSpecialCharacters(): String {
    return this.replace("\\", "").replace("\b", "").replace("\n", "").replace("\r", "").replace("\t", "")
}
