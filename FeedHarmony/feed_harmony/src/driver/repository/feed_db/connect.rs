use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::{FollowList, OneFeed};
use crate::domain::Feed;
use anyhow::Result;
use axum::async_trait;
use serde_json::Value;
use sqlx::mysql::MySqlPoolOptions;
use sqlx::Error as SqlxError;
use sqlx::{MySql, Pool, Row};
use std::fmt::Debug;
use std::str::FromStr;

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
    async fn get_all_feeds(&self) -> Result<Vec<FollowList>, SqlxError>;
    async fn insert_all_feeds(&self, feed: Vec<Feed>) -> Result<(), SqlxError>;
}

// TODO: need to think about using query builder
#[async_trait]
impl FeedConnection for FeedRepository {
    async fn get_all_feeds(&self) -> anyhow::Result<Vec<FollowList>, SqlxError> {
        let maybe_rows = sqlx::query("SELECT id, uuid, xml_version, rss_version, url, title, description, link, links, item_description, language, dt_created, dt_updated, dt_last_inserted, feed_category, is_favorite, is_active, is_read, is_updated FROM follow_lists")
            .fetch_all(&self.pool)
            .await
            .map_err(|e| anyhow::anyhow!(e));

        if let Err(e) = maybe_rows {
            println!("Failed to fetch all feeds: {}", e);
            return Err(SqlxError::RowNotFound);
        }

        let follow_lists = maybe_rows
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

    async fn insert_all_feeds(&self, feeds: Vec<Feed>) -> Result<(), SqlxError> {
        self.pool.begin().await?;

        for one_feed in feeds {
            let upserting_url = one_feed.feed_url.clone();

            let mut tx = self.pool.begin().await?;
            let _row = sqlx::query(
                "INSERT INTO feeds (id, site_url, title, description, feed_url, language, favorites, dt_created, dt_updated)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
                    ON DUPLICATE KEY UPDATE feed_url = ?;",   
            )
            .bind(one_feed.id)
                .bind(one_feed.site_url)
            .bind(one_feed.title)
                .bind(one_feed.description)
            .bind(one_feed.feed_url)
                .bind(one_feed.language)
            .bind(one_feed.favorites)
            .bind(one_feed.created_at)
            .bind(one_feed.updated_at)
                .bind(upserting_url)
            .execute(&mut tx)
            .await?;

            tx.commit().await?;
        }

        Ok(())
    }
}

pub async fn initialize_connection(var: String) -> Result<Pool<MySql>, SqlxError> {
    let pool = MySqlPoolOptions::new()
        .max_connections(10)
        .min_connections(4)
        .idle_timeout(std::time::Duration::from_secs(60 * 60 * 12))
        .max_lifetime(std::time::Duration::from_secs(60 * 60 * 12))
        .acquire_timeout(std::time::Duration::from_secs(10))
        .connect(&var)
        .await;

    match pool {
        Ok(_) => {
            println!("Connected to database");
            Ok(pool?)
        }
        Err(e) => {
            println!("Failed to connect to database:",);
            Err(SqlxError::Database(e.into_database_error().unwrap()))
        }
    }
}