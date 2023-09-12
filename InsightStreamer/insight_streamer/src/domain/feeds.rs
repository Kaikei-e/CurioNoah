use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(sqlx::Encode, Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
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
