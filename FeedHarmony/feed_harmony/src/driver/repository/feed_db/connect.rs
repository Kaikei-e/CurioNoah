use crate::domain::feed::FollowList;
use axum::async_trait;
use sqlx::mysql::MySqlPool;
use sqlx::{MySql, Pool};
use std::env;

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
        let follow_list = sqlx::query_as::<_, FollowList>("SELECT * FROM follow_list")
            .fetch_all(&self.pool)
            .await?;

        Ok(follow_list)
    }
}

pub async fn initialize_connection() -> anyhow::Result<Pool<MySql>> {
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

    let pool = MySqlPool::connect(&var).await;
    // match pool {
    //     Ok(_) => {
    //         println!("Connected to database");
    //         Ok(pool?)
    //     }
    //     Err(e) => {
    //         panic!("Failed to connect to database: {}", e)
    //     }
    // }
    match pool {
        Ok(_) => {
            println!("Connected to database");
            Ok(pool?)
        }
        Err(e) => {
            panic!("Failed to connect to database: {}", e)
        }
    }
}
