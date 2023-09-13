use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

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

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct FeedElement {
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
