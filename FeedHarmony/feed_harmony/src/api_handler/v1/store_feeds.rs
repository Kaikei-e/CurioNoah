use anyhow::Error;
use axum::{http::StatusCode, response::IntoResponse, Json};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

pub async fn store_feeds(Json(payload): Json<all_feeds>) -> impl IntoResponse {
    // Payload is a just boolean. This is bad.
    

    todo!()
}

#[derive(Clone, Debug, PartialEq, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct all_feeds {
    store: bool,
}

#[derive(Clone, Debug, PartialEq, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct follow_list_by_id {
    user_id: Uuid,
}
