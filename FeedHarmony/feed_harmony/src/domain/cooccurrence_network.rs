use serde::{Deserialize, Serialize};
use serde_json::Value;

#[derive(sqlx::Encode, Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct CooccurrenceNetworkPoolForRead {
    pub site_url: String,
    pub titles: Value,
    pub descriptions: Value,
}

#[derive(sqlx::Encode, Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
pub struct CooccurrenceNetworkPoolForWrite {
    pub id: String,
    pub site_url: String,
    pub titles: Value,
    pub descriptions: Value,
}
