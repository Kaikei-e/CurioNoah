package com.example.rest

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

@Serializable
data class SummarizerResponse(
    val model: String = "",
    @SerialName("created_at") val createdAt: String = "",
    val response: String = "",
    val done: Boolean = false,
    @SerialName("done_reason") val doneReason: String? = null,
    val context: List<Int> = emptyList(),
    @SerialName("total_duration") val totalDuration: Long = 0,
    @SerialName("load_duration") val loadDuration: Long = 0,
    @SerialName("prompt_eval_count") val promptEvalCount: Int = 0,
    @SerialName("prompt_eval_duration") val promptEvalDuration: Long = 0,
    @SerialName("eval_count") val evalCount: Int = 0,
    @SerialName("eval_duration") val evalDuration: Long = 0,
)