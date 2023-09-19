use crate::domain::feeds::FollowList;
use crate::driver::repository::rss_feeds_driver::{DatabasePool, RSSFeedRepository};
use crate::gateway::fetch_feeds_gateway::RSSFeedRepositoryTrait;
use anyhow::{anyhow, Error};

pub async fn follow_list_by_twenty(pool: DatabasePool) -> Result<Vec<FollowList>, Error> {
    let repository = RSSFeedRepository::new(pool.pool);
    let result = repository.fetch_follow_list_by_twenty().await;

    match result {
        Ok(result) => Ok(result),
        Err(e) => Err(anyhow!(e)),
    }
}
