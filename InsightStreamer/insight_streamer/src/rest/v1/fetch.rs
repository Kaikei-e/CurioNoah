use crate::domain::feed_source::FollowList;
use crate::rest::server::DatabasePool;
use axum::extract::State;
use axum::{response::IntoResponse, Json};

pub async fn fetch(
    State(pool): State<DatabasePool>,
    Json(payload): Json<FollowList>,
) -> impl IntoResponse {
    Json(payload)
}
