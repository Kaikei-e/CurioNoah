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
    use crate::domain::feeds::{FeedElement, FeedLinks, OneFeed};
    use async_trait::async_trait;
    use chrono::{DateTime, Utc};
    use mockall::{automock, mock, predicate::*};
    use uuid::Uuid;

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
        let uid = Uuid::new_v4();
        let feed_links = FeedLinks {
            links: vec![
                "link1".to_string(),
                "link2".to_string(),
                "link3".to_string(),
            ],
        };
        let one_feed = vec![FeedElement {
            guid: "id".to_string(),
            item_id: "id".to_string(),
            updated: "date".to_string(),
            item_link: "item_link".to_string(),
            published: "date".to_string(),
            item_title: "title".to_string(),
            item_description: Default::default(),
            published_parsed: "date".to_string(),
            categories: None,
        }];
        let dt_created = DateTime::<Utc>::from_naive_utc_and_offset(Default::default(), Utc);
        let dt_updated = DateTime::<Utc>::from_naive_utc_and_offset(Default::default(), Utc);
        let dt_last_inserted = DateTime::<Utc>::from_naive_utc_and_offset(Default::default(), Utc);

        let mut mock_driver = MockRSSFeedRepository::new();
        let result = mock_driver.expect_fetch_follow_list_by_twenty();

        let expected: Result<Vec<FollowList>, ()> = Ok(vec![FollowList {
            id: 1,
            uuid: uid,
            xml_version: 1,
            rss_version: 1,
            url: "url".to_string(),
            title: "title".to_string(),
            description: "description".to_string(),
            link: "link".to_string(),
            links: feed_links,
            item_description: one_feed,
            language: "language".to_string(),
            dt_created,
            dt_updated,
            dt_last_inserted,
            feed_category: 1,
            is_favorite: 1,
            is_active: 1,
            is_read: 1,
            is_updated: 1,
        }]);

        assert_eq!(expected, result);
        todo!("implement test")
    }
}