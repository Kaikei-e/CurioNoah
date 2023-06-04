use chrono::prelude::*;
use uuid::Uuid;

pub struct Feed {
    pub id: Uuid,
    pub title: String,
    pub url: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}
