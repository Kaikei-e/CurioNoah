package com.example.domain

import io.ktor.http.*

data class SummarizedFeed(
    val title: String = "failed to summarize title",
    val description: String = "failed to summarize description",
    val siteUrl: Url = Url("https://example.com"),
)
