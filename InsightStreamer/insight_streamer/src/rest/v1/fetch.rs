use crate::domain::follow_list::FollowList;
use crate::driver::repository::rss_feeds_driver::DatabasePool;
use axum::extract::State;
use axum::{response::IntoResponse, Json};

pub async fn fetch(
    State(_pool): State<DatabasePool>,
    Json(payload): Json<FollowList>,
) -> impl IntoResponse {
    Json(payload)
}
