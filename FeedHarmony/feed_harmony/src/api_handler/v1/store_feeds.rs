use axum::{response::IntoResponse, Json};
use serde::Deserialize;
use uuid::Uuid;

pub async fn store_feeds(Json(payload): Json<AllFeeds>) -> impl IntoResponse {
    // Payload is a just boolean. This is bad.
    let will_launch = payload.store;
    if will_launch {}

    todo!()
}

#[derive(Clone, Debug, PartialEq, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct AllFeeds {
    store: bool,
}

#[derive(Clone, Debug, PartialEq, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct FollowListById {
    user_id: Uuid,
}
