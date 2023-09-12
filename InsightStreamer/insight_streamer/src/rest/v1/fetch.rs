use crate::domain::feed_source::FollowList;
use crate::driver::repository::rss_feeds_driver::DatabasePool;
use axum::extract::State;
use axum::{response::IntoResponse, Json};

pub async fn fetch(
    State(pool): State<DatabasePool>,
    Json(payload): Json<FollowList>,
) -> impl IntoResponse {
    Json(payload)
}
