package com.example

import com.example.plugins.configureHTTP
import com.example.plugins.configureMonitoring
import com.example.plugins.configureSecurity
import com.example.plugins.configureSerialization
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
            val forSummarizerResponseJson = Json {
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

                    println(res)
                    // for japanese text

                    val htmlContent = res.bodyAsText(charset)
                    println(htmlContent)

                    val doc = Jsoup.parse(htmlContent)
                    val textBody = doc.body().text().trim()

                    //Print the document and text body for debugging
                    println(doc)
                    println(textBody)

                    val apiUrl = URLBuilder().apply {
                        protocol = URLProtocol.HTTP
                        host = summarizerAPIEndpoint
                        port = summarizerAPIEndpointPort.toInt()
                        path("api", "generate")
                    }.build()

                    val summarizedResponse = client.post(apiUrl) {
                        contentType(ContentType.Application.Json)
                        setBody(
                            body = """{"model": "llama3-elyza:8b", "prompt": "次の文章を要約して。また、キーとなる概念についても詳述して。> ($textBody)", stream: false}"""
//                            body = """{"model": "llama3-elyza:8b", "prompt": "why sky is blue?", "stream": false}"""
                        )
                    }.bodyAsText()

                    println(summarizedResponse)

                    val escapedResponse = summarizedResponse.escapeSpecialCharacters()

                    try {
                        val response = forSummarizerResponseJson.decodeFromString<SummarizerResponse>(escapedResponse)
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
