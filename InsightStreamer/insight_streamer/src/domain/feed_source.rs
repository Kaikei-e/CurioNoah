use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(sqlx::Encode, Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct FeedSource {
    pub id: Uuid,
    pub site_url: Option<String>,
    pub title: String,
    pub description: String,
    pub language: Option<String>,
    pub dt_created: DateTime<Utc>,
    pub dt_updated: DateTime<Utc>,
    pub favorites: i32,
}
