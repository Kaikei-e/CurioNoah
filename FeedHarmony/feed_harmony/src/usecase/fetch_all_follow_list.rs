use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::FollowLists;
use crate::driver::repository::feed_db::connect::{FeedConnection, FeedRepository};

pub async fn fetch_all_follow_list(pool: DatabasePool) -> anyhow::Result<Vec<FollowLists>> {
    let feed_repository = FeedRepository::new(pool.clone());

    match feed_repository.get_all_follow_list().await {
        Ok(feeds) => {
            println!("Fetched all feeds");
            Ok(feeds)
        }
        Err(e) => {
            println!("Failed to fetch all feeds: {:?}", e);
            Err(e.into())
        }
    }
}
