use std::fmt::Debug;
use std::str::FromStr;

use anyhow::Result;
use axum::async_trait;
use chrono::{DateTime, Utc};
use serde_json::Value;
use sqlx::mysql::MySqlPoolOptions;
use sqlx::Error as SqlxError;
use sqlx::{MySql, Pool, Row};

use crate::api_handler::handler::DatabasePool;
use crate::domain::audit_log_action::{AuditLog, AuditLogAction};
use crate::domain::feed::{FollowLists, OneFeed};
use crate::domain::Feed;

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
    async fn get_all_follow_list(&self) -> Result<Vec<FollowLists>, SqlxError>;
    async fn get_follow_list_by_time(&self) -> Result<Vec<FollowLists>, SqlxError>;
    async fn insert_all_feeds(&self, feed: Vec<Feed>) -> Result<(), SqlxError>;
    async fn insert_latest_feeds(
        &self,
        feed: Vec<Feed>,
        audit_log: AuditLog,
    ) -> Result<(), SqlxError>;
    async fn insert_action_to_feed_audit_log_table(
        &self,
        action: AuditLog,
    ) -> Result<(), SqlxError>;
    async fn update_all_follow_list(&self, follow_lists: Vec<FollowLists>) -> Result<(), SqlxError>;
    async fn update_follow_list_by_using_uuid(
        &self,
        follow_lists: Vec<FollowLists>,
    ) -> Result<(), SqlxError>;
}

#[async_trait]
impl FeedConnection for FeedRepository {
    async fn get_all_follow_list(&self) -> Result<Vec<FollowLists>, SqlxError> {
        let maybe_rows = sqlx::query(
            "SELECT id, uuid, xml_version, \
        rss_version, url, title, \
        description, link, links, \
        item_description, language, dt_created, \
        dt_updated, dt_last_inserted, feed_category, \
        is_favorite, is_active, is_read, \
        is_updated FROM follow_lists",
        )
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
            .map(|row| FollowLists {
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

    async fn get_follow_list_by_time(&self) -> Result<Vec<FollowLists>, SqlxError> {
        let row = sqlx::query(
            "SELECT id, updated_at, action_id \
            FROM feed_audit_trail_logs ORDER BY updated_at DESC LIMIT 1",
        )
        .fetch_optional(&self.pool)
        .await?;

        let latest_updated_at: DateTime<Utc> = if let Some(row) = row {
            row.get("updated_at")
        } else {
            Utc::now()
        };

        println!("latest_updated_at: {:?}", latest_updated_at);

        let interval_days = 90;

        let maybe_rows = sqlx::query(
            "SELECT * FROM \
        follow_lists WHERE dt_updated > DATE_SUB(?, INTERVAL ? DAY)",
        )
        .bind(latest_updated_at)
        .bind(interval_days)
        .fetch_all(&self.pool)
        .await;

        if let Err(e) = maybe_rows {
            println!("Failed to fetch all feeds: {}", e);
            return Err(SqlxError::RowNotFound);
        }

        let follow_list: Vec<FollowLists> = maybe_rows
            .unwrap()
            .iter()
            .map(|row| FollowLists {
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

        println!("follow_list length: {}", follow_list.len());

        Ok(follow_list)
    }

    async fn insert_all_feeds(&self, feeds: Vec<Feed>) -> Result<(), SqlxError> {
        self.pool.begin().await?;

        let mut result: Result<(), SqlxError> = Ok(());

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
                .execute(&mut *tx)
                .await?;

            result = tx.commit().await;
        }

        match result {
            Ok(_) => Ok(()),
            //TODO error handling
            Err(e) => Err(SqlxError::Database(e.into_database_error().unwrap())),
        }
    }

    async fn insert_latest_feeds(
        &self,
        feeds: Vec<Feed>,
        _audit_log: AuditLog,
    ) -> Result<(), SqlxError> {
        let maybe_action_row =
            sqlx::query("SELECT id FROM feed_audit_trail_actions WHERE action = ?")
                .bind(AuditLogAction::Upsert.convert_to_string())
                .fetch_one(&self.pool)
                .await;

        let action_row = match maybe_action_row {
            Ok(row) => row,
            Err(_) => {
                let mut tx = self.pool.begin().await?;

                let row = sqlx::query("INSERT INTO feed_audit_trail_actions (action) VALUES (?)")
                    .bind(AuditLogAction::Upsert.convert_to_string())
                    .execute(&mut *tx)
                    .await;

                if let Err(e) = row {
                    return Err(e.into());
                }

                match row {
                    Ok(_) => {
                        let row =
                            sqlx::query("SELECT id FROM feed_audit_trail_actions WHERE action = ?")
                                .bind(AuditLogAction::Upsert.convert_to_string())
                                .fetch_one(&mut *tx)
                                .await?;

                        tx.commit().await?;

                        row
                    }
                    Err(e) => {
                        println!("Failed to insert action: {}", e);
                        return Err(SqlxError::RowNotFound);
                    }
                }
            }
        };

        let action_id = action_row.get::<i64, _>("id");

        let result = sqlx::query("SELECT COUNT(*) as count FROM feed_audit_trail_logs")
            .fetch_one(&self.pool)
            .await;

        let count = match result {
            Ok(row) => row.get("count"),
            Err(_) => 0,
        };

        let now = Utc::now();

        let result: Result<(), SqlxError>;

        if count == 0 {
            let mut tx = self.pool.begin().await?;

            let _row = sqlx::query(
                "INSERT INTO feed_audit_trail_logs (updated_at, action_id) VALUES (?, ?)",
            )
            .bind(now)
            .bind(action_id)
            .execute(&mut *tx)
            .await?;

            for one_feed in feeds {
                let upserting_url = one_feed.feed_url.clone();

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
                    .execute(&mut *tx)
                    .await?;
            }
            result = tx.commit().await;
        } else {
            let mut tx = self.pool.begin().await?;

            for one_feed in feeds {
                let upserting_url = one_feed.feed_url.clone();

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
                    .execute(&mut *tx)
                    .await?;
            }

            let _row = sqlx::query(
                "INSERT INTO feed_audit_trail_logs (updated_at, action_id) VALUES (?, ?)",
            )
            .bind(now)
            .bind(action_id)
            .execute(&mut *tx)
            .await?;
            result = tx.commit().await;
        }

        match result {
            Ok(_) => Ok(()),
            //TODO error handling
            Err(e) => Err(SqlxError::Database(e.into_database_error().unwrap())),
        }
    }

    async fn insert_action_to_feed_audit_log_table(
        &self,
        audit_log: AuditLog,
    ) -> Result<(), SqlxError> {
        let mut tx = self.pool.begin().await?;

        let row = match audit_log.action {
            AuditLogAction::Upsert => sqlx::query(
                "INSERT INTO feed_audit_trail (updated_at, action)
                    VALUES (?, ?);",
            )
            .bind(audit_log.updated_at)
            .bind(audit_log.action.convert_to_string()),
            AuditLogAction::Delete => sqlx::query(
                "INSERT INTO feed_audit_trail (updated_at, action)
                    VALUES (?, ?);",
            )
            .bind(audit_log.updated_at)
            .bind(audit_log.action.convert_to_string()),
            AuditLogAction::Fail => sqlx::query(
                "INSERT INTO feed_audit_trail (updated_at, action)
                    VALUES (?, ?);",
            )
            .bind(audit_log.updated_at)
            .bind(audit_log.action.convert_to_string()),
        };

        let _row = row.execute(&mut *tx).await?;
        let result = tx.commit().await;
        match result {
            Ok(_) => Ok(()),
            //TODO error handling
            Err(e) => Err(SqlxError::Database(e.into_database_error().unwrap())),
        }
    }

    async fn update_all_follow_list(&self, follow_lists: Vec<FollowLists>) -> Result<(), SqlxError> {
        let mut tx = self.pool.begin().await?;
        let now = Utc::now().format("%Y-%m-%dT%H:%M:%S").to_string();

        println!("now: {:?}", now);

        let mut row_count = 0;

        for follow_list in follow_lists {
            match sqlx::query("UPDATE follow_lists SET dt_updated = ? WHERE uuid = ?")
                .bind(now.clone())
                .bind(follow_list.uuid.to_string())
                .execute(&mut *tx)
                .await
            {
                Ok(res) => {
                    if res.rows_affected() > 0 {
                        println!("Update was successful for uuid: {}", follow_list.uuid);
                    } else {
                        println!(
                            "Update didn't affect any rows for uuid: {}",
                            follow_list.uuid
                        );
                    }
                    row_count += res.rows_affected();
                }
                Err(e) => {
                    println!("Error: {}", e);
                }
            }
        }

        println!("row_count: {}", row_count);

        if row_count > 0 {
            tx.commit().await?;
        }

        Ok(())
    }

    async fn update_follow_list_by_using_uuid(
        &self,
        follow_lists: Vec<FollowLists>,
    ) -> Result<(), SqlxError> {
        let mut tx = self.pool.begin().await?;
        let now = Utc::now().format("%Y-%m-%dT%H:%M:%S").to_string();

        println!("now: {:?}", now);

        let mut row_count = 0;

        for follow_list in follow_lists {
            match sqlx::query("UPDATE follow_lists SET dt_updated = ? WHERE uuid = ?")
                .bind(now.clone())
                .bind(follow_list.uuid.to_string())
                .execute(&mut *tx)
                .await
            {
                Ok(res) => {
                    if res.rows_affected() > 0 {
                        println!(
                            "Update was successful for uuid: {:?}",
                            follow_list
                                .uuid
                                .to_string()
                                .chars()
                                .enumerate()
                                .filter(|&(i, _)| i > 0 && i < 7)
                                .fold("".to_string(), |s, (_, c)| format!("{}{}", s, c))
                        );
                    } else {
                        println!(
                            "Update didn't affect any rows for uuid: {}",
                            follow_list.uuid
                        );
                    }
                    row_count += res.rows_affected() as usize;
                }
                Err(e) => {
                    println!(
                        "Update failed for uuid: {}. Error: {:?}",
                        follow_list.uuid, e
                    );
                    // Depending on your use case you may want to return here
                }
            };
        }

        println!("row_count: {:?}", row_count);

        let result = tx.commit().await;
        match result {
            Ok(_) => Ok(()),
            Err(e) => Err(SqlxError::Database(e.into_database_error().unwrap())),
        }
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

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn success_feed_repository() {
        // let expected = MockFeedRepository::new();

        let var = dotenvy::var("DATABASE_URL").unwrap();
        let pool = initialize_connection(var).await.unwrap();
        let feed_repository = FeedRepository::new(DatabasePool { pool });
        let result = feed_repository.get_all_follow_list().await;

        assert!(result.is_ok());
    }

    // #[tokio::test]
    // async fn test_initialize_connection() {
    //     let var = dotenvy::var("DATABASE_URL").unwrap();
    //
    //     let pool = initialize_connection(var);
    //
    //
    //
    //     todo!()
    // }
}
