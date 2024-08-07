package com.example.usecase.summarize_today

import com.example.domain.SummarizedFeed
import io.ktor.http.*

fun execute(): List<SummarizedFeed> {
    val feeds = listOf(
        SummarizedFeed(
            title = "Title 1",
            description = "Description 1",
            siteUrl = Url("https://www.google.com")
        ),
        SummarizedFeed(
            title = "Title 2",
            description = "Description 2",
            siteUrl = Url("https://www.google.com")
        )
    )

    return feeds
}