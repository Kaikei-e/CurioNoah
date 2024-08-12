use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::FollowLists;
use crate::driver::repository::feed_db::connect::{FeedConnection, FeedRepository};

pub async fn fetch_latest_follow_list(pool: DatabasePool) -> anyhow::Result<Vec<FollowLists>> {
    let feed_repository = FeedRepository::new(pool.clone());

    match feed_repository.get_follow_list_by_time().await {
        Ok(follow_list) => {
            println!("Fetched all feeds");
            Ok(follow_list)
        }
        Err(e) => {
            println!("Failed to fetch all feeds: {:?}", e);
            Err(e.into())
        }
    }
}
