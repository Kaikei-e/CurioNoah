use crate::domain::feed::FollowList;
use axum::async_trait;
use sqlx::mysql::MySqlPool;
use sqlx::{MySql, Pool};

#[derive(Debug, Clone)]
pub struct FeedRepository {
    pub pool: Pool<MySql>,
}

impl FeedRepository {
    pub fn new(pool: Pool<MySql>) -> Self {
        FeedRepository { pool }
    }
}

#[async_trait]
pub trait FeedConnection {
    async fn get_all_feeds(&self) -> anyhow::Result<Vec<FollowList>>;
}

// TODO: need to think about using query builder
#[async_trait]
impl FeedConnection for FeedRepository {
    async fn get_all_feeds(&self) -> anyhow::Result<Vec<FollowList>> {
        let follow_list: FollowList;

        let follow_list = sqlx::query_as::<_, FollowList>("SELECT * FROM follow_list")
            .bind(follow_list)
            .fetch_all(&self.pool)
            .await?;
        Ok(follow_list);
        todo!("Implement this function")
    }
}

#[tokio::main]
pub async fn initialize_connection(var: String) -> anyhow::Result<Pool<MySql>, sqlx::Error> {
    let pool = MySqlPool::connect(&var).await;
    match pool {
        Ok(_) => {
            println!("Connected to database");
            pool
        }
        Err(e) => {
            panic!("Failed to connect to database: {}", e)
        }
    }
}
