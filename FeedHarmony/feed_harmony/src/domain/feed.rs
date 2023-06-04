use axum::http;
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

pub struct FeedList {
    pub uuid: Uuid,
    pub guid: String,
    pub item_link: http::Uri,
    pub item_title: String,
    pub item_description: String,
    pub published_parsed: DateTime<Utc>,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

