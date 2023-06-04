use anyhow::Error;
use axum::{http::StatusCode, response::IntoResponse, Json};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

pub async fn store_feeds(Json(payload): Json<store_all>) -> impl IntoResponse {
    todo!()
}

#[derive(Clone, Debug, PartialEq, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct store_all {
    user_id: Uuid,
    store_all: bool,
}
