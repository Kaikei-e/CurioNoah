package com.example

import com.example.plugins.configureHTTP
import com.example.plugins.configureSecurity
import com.example.plugins.configureSerialization
import com.example.rest.OllamaRequest
import com.example.rest.SummarizerResponse
import com.example.usecase.summarize_today.SummarizeTodayUseCase
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
import org.jetbrains.exposed.sql.Database
import org.jsoup.Jsoup
import java.io.File


fun main() {
    embeddedServer(Netty, port = 5555, host = "0.0.0.0", module = Application::module).start(wait = true)
}

fun Application.module() {
    val env = loadEnv()
    configureSecurity()
    configureHTTP()
    configureSerialization()
    configureRouting()
    configureDatabases(env)
}


fun Application.configureRouting() {
    val env = loadEnv()
    val summarizerAPIEndpoint = env["SUMMARIZER_API_ENDPOINT"] ?: throw Exception("SUMMARIZER_API_ENDPOINT is not set")
    val summarizerAPIEndpointPort =
        env["SUMMARIZER_API_ENDPOINT_PORT"] ?: throw Exception("SUMMARIZER_API_ENDPOINT_PORT is not set")

    val client = HttpClient(CIO) {
        install(HttpTimeout) {
            requestTimeoutMillis = 80_000 // 80 seconds
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
                    // will implement the logic to fetch articles from database
                    val testURL1 = "https://zenn.dev/e_kaikei/articles/tauri-rust-react-setup-plan-1"

                    val summarizeTodayUseCase = SummarizeTodayUseCase()
                    summarizeTodayUseCase.execute()

                    val articleList = listOf(testURL1)

                    var combinedHtmlContent = ""
                    for ((index, url) in articleList.withIndex()) {
                        val res: HttpResponse = client.get(url)

                        val contentType = res.headers[HttpHeaders.ContentType]
                        val charset = contentType?.let { ContentType.parse(it).charset() } ?: Charsets.UTF_8

                        val htmlContent = res.bodyAsText(charset)
                        val parsedHtml = Jsoup.parse(htmlContent)
                        val body = parsedHtml.body().text()

                        combinedHtmlContent += "\n\narticle No.$index: \n-------------------------\n content\n$body"

                        // sleep for 1 second
                        Thread.sleep(1000)
                    }

                    val apiUrl = URLBuilder().apply {
                        protocol = URLProtocol.HTTP
                        host = summarizerAPIEndpoint
                        port = summarizerAPIEndpointPort.toInt()
                        path("api", "generate")
                    }.build()

                    val prompt = """
                        Summarize bellow article. 
                        First give the title of the article in Japanese.
                        Then, summarize the article in Japanese.
                        In doing so, respond in bullet points with key points and crucial information in Japanese.
                        
                        ---------------------------------------
                        """.trimMargin()
                    val combinedPrompt = "$prompt\n\n$combinedHtmlContent"

                    val summarizerRequest = OllamaRequest(
                        model = "gemma2-lite:latest", prompt = combinedPrompt, stream = false
                    )

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

fun Application.configureDatabases(env: Map<String, String>) {
    val host = env["DATABASE_HOST"] ?: throw Exception("DATABASE_HOST is not set")
    val port = env["DATABASE_PORT"] ?: throw Exception("DATABASE_PORT is not set")
    val databaseName = env["DATABASE_NAME"] ?: throw Exception("DATABASE_NAME is not set")
    val user = env["DATABASE_USER"] ?: throw Exception("DATABASE_USER is not set")
    val password = env["DATABASE_PASSWORD"] ?: throw Exception("DATABASE_PASSWORD is not set")

    Database.connect(
        url = "jdbc:postgresql://$host:$port/$databaseName",
        driver = "org.postgresql.Driver",
        user = user,
        password = password
    )
}
