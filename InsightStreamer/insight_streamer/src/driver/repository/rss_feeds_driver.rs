use crate::domain::feed_source::{FeedElement, FollowList};
use axum::async_trait;
use sqlx::{Error, MySql, Pool};

#[derive(Clone)]
pub struct DatabasePool {
    pub(crate) pool: Pool<MySql>,
}

#[derive(Debug, Clone)]
pub struct RSSFeedRepository {
    pub pool: Pool<MySql>,
}

impl RSSFeedRepository {
    pub fn new(pool: DatabasePool) -> Self {
        RSSFeedRepository { pool: pool.pool }
    }
}

#[async_trait]
pub trait RSSFeedRepositoryTrait {
    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, sqlx::Error>;
    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
}

#[async_trait]
impl RSSFeedRepositoryTrait for RSSFeedRepository {
    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, Error> {
        todo!()
    }

    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, Error> {
        todo!()
    }
}
