package com.example

import com.example.plugins.configureHTTP
import com.example.plugins.configureMonitoring
import com.example.plugins.configureSecurity
import com.example.plugins.configureSerialization
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import org.jsoup.Jsoup

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
    val client = HttpClient(CIO)

    routing {
        get("/v1/systems/alive") {
            call.respondText("keep alive")
        }

        get("/v1/news/summarize") {
            call.respondText(
                contentType = ContentType.Application.Json,
                status = HttpStatusCode.OK,
                text = "{'message':'summarize news'}"
            )
        }

        get("/v1/news/summarize/today") {
            val testURL = "https://zenn.dev/e_kaikei/articles/tauri-rust-react-setup-plan-1"
            val res = client.get(
                Url(testURL)
            )

            val htmlContent: String = res.bodyAsText()
            val document = Jsoup.parse(htmlContent)
            val textBody = document.body().text()

            call.respondText(textBody)
        }
    }
}