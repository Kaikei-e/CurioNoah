package com.example.rest

import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*


fun Application.configureRouting() {
    routing {
        get("/v1/news/summarize/today") {
            call.respondText("Today's news summary")
        }
    }
}