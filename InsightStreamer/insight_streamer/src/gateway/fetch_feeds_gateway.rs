use crate::domain::feeds::{FeedElement, FollowList};
use axum::async_trait;

#[async_trait]
pub trait RSSFeedRepositoryTrait {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_feeds(&self, limit: i32, offset: i32) -> Result<Vec<FeedElement>, sqlx::Error>;
    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_follow_list_with_offset(
        &self,
        offset: i32,
    ) -> Result<Vec<FollowList>, sqlx::Error>;
}
