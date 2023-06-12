use axum::http;
use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use serde_json::Value;
use uuid::Uuid;

pub struct Feed {
    pub id: Uuid,
    pub site_url: String,
    pub feed_url: String,
    pub title: String,
    pub language: String,
    pub description: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
    pub favorites: i32,
}

// TODO: just copied from db schema, need to update and refactor
#[derive(sqlx::FromRow, Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct FollowList {
    pub id: i32,
    pub uuid: Uuid,
    pub xml_version: i8,
    pub rss_version: i8,
    pub url: String,
    pub title: String,
    pub description: String,
    pub link: String,
    pub links: Value,
    pub item_description: Value,
    pub language: String,
    pub dt_created: DateTime<Utc>,
    pub dt_updated: DateTime<Utc>,
    pub dt_last_inserted: DateTime<Utc>,
    pub feed_category: i32,
    pub is_favorite: i8,
    pub is_active: i8,
    pub is_read: i8,
    pub is_updated: i8,
}
