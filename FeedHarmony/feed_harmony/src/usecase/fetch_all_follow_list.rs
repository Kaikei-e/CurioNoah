use crate::domain::feed::FeedList;
use crate::driver::repository::feed_db::connect::{initialize_connection, FeedConnection};

pub async fn fetch_all_follow_list(conn_info: String) -> anyhow::Result<Vec<FeedList>> {
    let feed_repository =  <dyn FeedConnection>::new(initialize_connection(conn_info)?);

    feed_repository.get_all_feeds().await;

    todo!()
}
