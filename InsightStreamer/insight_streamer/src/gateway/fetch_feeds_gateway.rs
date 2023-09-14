use crate::domain::feeds::FeedElement;
use crate::domain::follow_list::FollowList;
use axum::async_trait;

#[async_trait]
pub trait RSSFeedRepositoryTrait {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, sqlx::Error>;
    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_follow_list_with_offset(
        &self,
        offset: i32,
    ) -> Result<Vec<FollowList>, sqlx::Error>;
}
