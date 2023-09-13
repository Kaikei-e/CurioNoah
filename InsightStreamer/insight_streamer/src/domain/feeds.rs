use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct Feed {
    pub id: String,
    pub site_url: String,
    pub feed_url: String,
    pub title: String,
    pub language: String,
    pub description: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
    pub favorites: i32,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct FollowList {
    pub id: i32,
    pub uuid: Uuid,
    pub xml_version: i8,
    pub rss_version: i8,
    pub url: String,
    pub title: String,
    pub description: String,
    pub link: String,
    pub links: FeedLinks,
    pub item_description: OneFeed,
    pub language: String,
    pub dt_created: DateTime<Utc>,
    pub dt_updated: DateTime<Utc>,
    pub dt_last_inserted: DateTime<Utc>,
    pub feed_category: i32,
    pub is_favorite: i8,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct FeedLinks {
    pub links: Vec<String>,
}

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct OneFeed {
    pub guid: String,
    pub item_id: String,
    pub updated: String,
    pub item_link: String,
    pub published: String,
    pub item_title: String,
    pub item_description: String,
    pub published_parsed: String,
    pub categories: Option<Vec<String>>,
}
