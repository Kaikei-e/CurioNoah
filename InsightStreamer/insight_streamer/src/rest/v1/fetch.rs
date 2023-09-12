use crate::rest::server::DatabasePool;
use axum::extract::State;
use axum::{response::IntoResponse, Json};
use serde::Deserialize;
use uuid::Uuid;

pub async fn fetch(
    State(pool): State<DatabasePool>,
    Json(payload): Json<String>,
) -> impl IntoResponse {
}
