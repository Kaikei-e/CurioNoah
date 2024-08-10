package com.example.usecase.summarize_today

import com.example.domain.SummarizedFeed
import com.example.gateway.SummarizeTodayGateway
import io.ktor.http.*

class SummarizeTodayUseCase {
    fun execute(): List<SummarizedFeed> {
        val summarizeTodayGateway = SummarizeTodayGateway()
        summarizeTodayGateway.summarizeTodayFeed()

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
}
