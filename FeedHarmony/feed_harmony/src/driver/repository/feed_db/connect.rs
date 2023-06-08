use std::str::FromStr;
use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::FollowList;
use axum::async_trait;
use sqlx::mysql::{MySqlPool, MySqlRow};
use sqlx::{MySql, Pool, Row};

#[derive(Debug, Clone)]
pub struct FeedRepository {
    pub pool: Pool<MySql>,
}

impl FeedRepository {
    pub fn new(pool: DatabasePool) -> Self {
        FeedRepository { pool: pool.pool }
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
        let rows = sqlx::query("SELECT * FROM follow_lists")
            .fetch_all(&self.pool)
            .await
            .map_err(|e| anyhow::anyhow!(e))?;

        let follow_lists = rows
            .iter()
            .map(|row| FollowList {
                id: row.get("id"),
                uuid: uuid::Uuid::from_str(row.get("uuid")).map_err(|e| anyhow::anyhow!(e)).unwrap(),
                xml_version: row.get("xml_version"),
                rss_version: row.get("rss_version"),
                url: row.get("url"),
                title: row.get("title"),
                description: row.get("description"),
                link: row.get("link"),
                links: row.get("links"),
                item_description: row.get("item_description"),
                language: row.get("language"),
                dt_created: row.get("dt_created"),
                dt_updated: row.get("dt_updated"),
                dt_last_inserted: row.get("dt_last_inserted"),
                feed_category: row.get("feed_category"),
                is_favorite: row.get("is_favorite"),
                is_active: row.get("is_active"),
                is_read: row.get("is_read"),
                is_updated: row.get("is_updated"),
            })
            .collect();

        Ok(follow_lists)
    }
}

pub async fn initialize_connection(var: String) -> anyhow::Result<Pool<MySql>> {
    let pool = MySqlPool::connect(&var).await;
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
