use crate::domain::feed::FollowList;
use crate::driver::repository::feed_db::connect::{
    initialize_connection, FeedConnection, FeedRepository,
};

pub async fn fetch_all_follow_list(conn_info: String) -> anyhow::Result<Vec<FollowList>> {
    let feed_repository = FeedRepository::new(initialize_connection(conn_info)?);

    match feed_repository.get_all_feeds().await {
        Ok(feeds) => {
            println!("Fetched all feeds");
            Ok(feeds)
        }
        Err(e) => {
            println!("Failed to fetch all feeds: {}", e);
            Err(e)
        }
    }
}
