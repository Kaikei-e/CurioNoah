use crate::domain::feeds::{FeedElement, FollowList, OneFeed};
use crate::gateway::fetch_feeds_gateway::RSSFeedRepositoryTrait;
use axum::async_trait;
use serde_json::Value;
use sqlx::{Error, MySql, Pool, Row};
use std::env;
use std::str::FromStr;

#[derive(Clone)]
pub struct DatabasePool {
    pub(crate) pool: Pool<MySql>,
}

#[derive(Debug, Clone)]
pub struct RSSFeedRepository {
    pub pool: Pool<MySql>,
}

impl RSSFeedRepository {
    pub fn new(pool: Pool<MySql>) -> Self {
        Self { pool }
    }
    pub async fn initialize_connection() -> Pool<MySql> {
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

        match pool_future.await {
            Ok(p) => p,
            Err(e) => panic!("Failed to connect to database: {}", e),
        }
    }
}

#[async_trait]
impl RSSFeedRepositoryTrait for RSSFeedRepository {
    async fn fetch_follow_list_by_twenty(&self) -> Result<Vec<FollowList>, Error> {
        // This is bad practice, I need to fix this to use parameters.
        // Because this function has a side effect.
        let query_limit = 20;
        let mut conn = self.pool.acquire().await?;

        let maybe_rows = sqlx::query(
            "SELECT id, uuid, xml_version, \
            rss_version, url, title, \
            description, link, links, \
            item_description, language, dt_created, \
            dt_updated, dt_last_inserted, feed_category, \
            is_favorite, is_active, is_read, \
            is_updated FROM follow_list LIMIT ?",
        )
        .bind(query_limit)
        .fetch_all(&mut conn)
        .await;

        if let Err(e) = maybe_rows {
            println!("Error: {}", e);
            return Err(sqlx::Error::RowNotFound);
        }

        let follow_lists: Vec<FollowList> = maybe_rows
            .unwrap()
            .iter()
            .map(|row| FollowList {
                id: row.get("id"),
                uuid: uuid::Uuid::from_str(&row.get::<String, _>("uuid")).unwrap(),
                xml_version: row.get("xml_version"),
                rss_version: row.get("rss_version"),
                url: row.get("url"),
                title: row.get("title"),
                description: row.get("description"),
                link: row.get("link"),
                links: serde_json::from_str(&row.get::<Value, _>("links").to_string()).unwrap(),
                item_description: serde_json::from_str::<OneFeed>(
                    &row.get::<Value, _>("item_description").to_string(),
                )
                .unwrap(),
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

    async fn fetch_feeds(&self, limit: i32, offset: i32) -> Result<Vec<FeedElement>, Error> {
        let mut conn = self.pool.acquire().await?;

        let maybe_rows = sqlx::query(
            "SELECT guid, item_id, updated, \
            item_link, published, item_title, \
            item_description, published_parsed, \
            categories FROM feeds LIMIT ? OFFSET ?",
        )
        .bind(limit)
        .bind(offset)
        .fetch_all(&mut conn)
        .await;

        if let Err(e) = maybe_rows {
            println!("Error: {}", e);
            return Err(sqlx::Error::RowNotFound);
        };

        let feeds: Vec<FeedElement> = maybe_rows
            .unwrap()
            .iter()
            .map(|row| FeedElement {
                guid: row.get("guid"),
                item_id: row.get("item_id"),
                updated: row.get("updated"),
                item_link: row.get("item_link"),
                published: row.get("published"),
                item_title: row.get("item_title"),
                item_description: serde_json::from_str(
                    &row.get::<Value, _>("item_description").to_string(),
                )
                .unwrap(),
                published_parsed: row.get("published_parsed"),
                categories: serde_json::from_str(&row.get::<Value, _>("categories").to_string())
                    .unwrap(),
            })
            .collect();

        Ok(feeds)
    }

    async fn fetch_follow_list(&self) -> Result<Vec<FollowList>, Error> {
        todo!()
    }

    async fn fetch_follow_list_with_offset(&self, _offset: i32) -> Result<Vec<FollowList>, Error> {
        todo!()
    }
}
