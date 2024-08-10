package com.example.gateway

import com.example.domain.SummarizedFeed
import com.example.port.SummarizeTodayFeedPort
import io.ktor.http.*

class SummarizeTodayGateway : SummarizeTodayFeedPort {
    override fun summarizeTodayFeed(): List<SummarizedFeed> {
        return listOf(
            SummarizedFeed(
                title = "title",
                description = "description",
                siteUrl = Url("https://example.com"),
            )
        )
    }
}