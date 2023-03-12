use axum::http::StatusCode;
use axum::Json;
use axum::response::IntoResponse;
use serde::{Serialize, Deserialize};
use crate::fiber::server::summarizer::domain::domain::summarize;

#[derive(Deserialize)]
pub struct SummarizerRequest {
    pub text: String,
    pub summary_length: usize,
}

#[derive(Serialize)]
pub struct SummarizedResponse {
    pub summary: String,
}

pub async fn summarizer(Json(payload): Json<SummarizerRequest>) -> impl IntoResponse {
    let summarizer_request = SummarizerRequest {
        text: payload.text,
        summary_length: payload.summary_length,
    };
    let text = summarizer_request.text;
    let summary_length = summarizer_request.summary_length;
    let summary = summarize(text, summary_length);
    let response = SummarizedResponse {
        summary,
    };

    (StatusCode::OK, Json(response))
}
