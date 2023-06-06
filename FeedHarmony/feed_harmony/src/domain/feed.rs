use axum::{http, Json};
use chrono::prelude::*;
use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use uuid::Uuid;

pub struct Feed {
    pub uuid: Uuid,
    pub guid: String,
    pub item_link: http::Uri,
    pub item_title: String,
    pub item_description: String,
    pub published_parsed: NaiveDateTime,
    pub created_at: NaiveDateTime,
    pub updated_at: NaiveDateTime,
}

// TODO: just copied from db schema, need to update and refactor
#[derive(Debug, sqlx::FromRow, Serialize, Deserialize, Clone, PartialEq, Eq)]
pub struct FollowList {
    pub id: i32,
    pub uuid: Uuid,
    pub xml_version: i8,
    pub rss_version: i8,
    pub url: http::Uri,
    pub title: String,
    pub description: String,
    pub link: http::Uri,
    pub links: Json<String>,
    pub item_description: Json<String>,
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
