use crate::domain::feed_source::{FeedElement, FollowList};
use axum::async_trait;

#[async_trait]
pub trait RSSFeedRepositoryTrait {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, sqlx::Error>;
    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
}
