use crate::domain::feed_source::{FeedElement, FollowList};
use axum::async_trait;
use sqlx::{Error, MySql, Pool};
use std::env;

#[derive(Clone)]
pub struct DatabasePool {
    pub(crate) pool: Pool<MySql>,
}

#[derive(Debug, Clone)]
pub struct RSSFeedRepository {
    pub pool: Pool<MySql>,
}

impl RSSFeedRepository {
    pub async fn new() -> Pool<MySql> {
        let loaded = dotenvy::dotenv();

        match loaded {
            Ok(_) => {
                println!("Loaded .env file");
            }
            Err(e) => {
                panic!("Failed to load .env file: {}", e)
            }
        }
        let var = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
        let pool_future = sqlx::mysql::MySqlPool::connect(&var);
        let pool = match pool_future.await {
            Ok(p) => p,
            Err(e) => panic!("Failed to connect to database: {}", e),
        };
        pool
    }
}

#[async_trait]
pub trait RSSFeedRepositoryTrait {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, sqlx::Error>;
    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, sqlx::Error>;
    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, sqlx::Error>;
}

#[async_trait]
impl RSSFeedRepositoryTrait for RSSFeedRepository {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, Error> {
        todo!()
    }

    async fn fetch_feeds(&self) -> Result<Vec<FeedElement>, Error> {
        todo!()
    }

    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, Error> {
        todo!()
    }
}
