use crate::domain::feed::FeedList;
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
    async fn get_all_feeds(&self) -> anyhow::Result<Vec<FeedList>>;
}

// TODO: need to think about using query builder
#[async_trait]
impl FeedConnection for FeedRepository {
    async fn get_all_feeds(&self) -> anyhow::Result<Vec<FeedList>> {
        const feed_list: FeedList = FeedList {
            id: 0,
            uuid: uuid::Uuid::new_v4(),
            xml_version: 0,
            rss_version: 0,
            url: Default::default(),
            title: "".to_string(),
            description: "".to_string(),
            link: Default::default(),
            links: Default::default(),
            item_description: Default::default(),
            language: "".to_string(),
            dt_created: Default::default(),
            dt_updated: Default::default(),
            dt_last_inserted: Default::default(),
            feed_category: 0,
            is_favorite: 0,
            is_active: 0,
            is_read: 0,
            is_updated: 0,
        };

        let feeds = sqlx::query_as::<_, FeedList>("SELECT * FROM follow_list")
            .fetch_all(&self.pool)
            .await?;
        Ok(feeds);
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
