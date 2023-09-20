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

#[cfg(test)]
mod tests {
    use super::*;
    use crate::domain::feeds::FeedElement;
    use async_trait::async_trait;
    use mockall::{automock, mock, predicate::*};

    mock! {
        pub RSSFeedRepository {}

        #[async_trait]
        impl RSSFeedRepositoryTrait for RSSFeedRepository {
            async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, sqlx::Error>;
            async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, sqlx::Error>;
            async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
            async fn fetch_follow_list_with_offset(
                &self,
                offset: i32,
            ) -> Result<Vec<FollowList>, sqlx::Error>;
        }
    }

    #[tokio::test]
    async fn test_success_follow_list_by_twenty() {
        let mock_driver = MockRSSFeedRepository::new();
        let result = mock_driver.fetch_follow_list_by_twenty();

        todo!("implement test")
    }
}
