package com.example.port

import com.example.domain.SummarizedFeed

interface SummarizeTodayFeedPort {
    fun summarizeTodayFeed()
}

interface FetchTodayFeedsPort {
    fun fetchTodayFeeds(): List<SummarizedFeed>
}