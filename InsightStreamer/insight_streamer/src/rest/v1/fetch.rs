use crate::rest::server::DatabasePool;
use axum::extract::State;
use axum::{response::IntoResponse, Json};
use crate::domain::feed_source::FollowList;

pub async fn fetch(
    State(pool): State<DatabasePool>,
    Json(payload): Json<FollowList>,
) -> impl IntoResponse {
    Json(payload)
}
