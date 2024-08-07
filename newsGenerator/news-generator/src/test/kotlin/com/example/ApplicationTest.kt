package com.example

import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.server.testing.*
import kotlin.test.Test
import kotlin.test.assertEquals

class ApplicationTest {
    @Test
    fun testRoot() = testApplication {
        application {
            configureRouting()
        }
        client.get("/api/v1/systems/alive").apply {
            assertEquals(HttpStatusCode.OK, status)
            assertEquals("keep alive", bodyAsText())
        }
    }
}
