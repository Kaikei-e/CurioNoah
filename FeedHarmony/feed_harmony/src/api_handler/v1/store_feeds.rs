use crate::api_handler::handler::DatabasePool;
use crate::usecase;
use axum::extract::State;
use axum::{response::IntoResponse, Json};
use serde::Deserialize;
use uuid::Uuid;

pub async fn store_feeds(
    State(pool): State<DatabasePool>,
    Json(payload): Json<AllFeeds>,
) -> impl IntoResponse {
    // Payload is a just boolean. This is bad.
    let will_launch = payload.store;
    if will_launch {
        let follow_list = usecase::fetch_all_follow_list::fetch_all_follow_list(pool).await;
        match follow_list {
            Ok(feeds) => {
                println!("Fetched all feeds");
                Json(feeds)
            }
            Err(e) => {
                println!("Failed to fetch all feeds: {}", e);
                todo!("Return error response")
            }
        }
    } else {
        todo!("Return error response")
    }
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
