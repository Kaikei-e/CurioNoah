use axum::{http, Json};
use chrono::prelude::*;
use uuid::Uuid;

pub struct Feed {
    pub uuid: Uuid,
    pub guid: String,
    pub item_link: http::Uri,
    pub item_title: String,
    pub item_description: String,
    pub published_parsed: DateTime<Utc>,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

// TODO: just copied from db schema, need to update
pub struct FeedList {
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
