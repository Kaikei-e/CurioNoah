use crate::api_handler::handler::DatabasePool;
use crate::domain::cooccurrence_network::CooccurrenceNetworkPoolForRead;
use axum::async_trait;
use serde_json::json;
use sqlx::{Error, MySql, Pool};

#[derive(Debug, Clone)]
pub struct WordNetworkRepository {
    pub pool: Pool<MySql>,
}

impl WordNetworkRepository {
    pub fn new(pool: DatabasePool) -> Self {
        WordNetworkRepository { pool: pool.pool }
    }
}

#[async_trait]
pub trait FetchSiteUrlGroup {
    async fn fetch_site_url_group(&self) -> Result<Vec<CooccurrenceNetworkPoolForRead>, Error>;
    async fn insert_site_url_group(
        &self,
        grouped_urls: Vec<CooccurrenceNetworkPoolForRead>,
    ) -> Result<(), Error>;
}

#[async_trait]
impl FetchSiteUrlGroup for WordNetworkRepository {
    async fn fetch_site_url_group(&self) -> Result<Vec<CooccurrenceNetworkPoolForRead>, Error> {
        let rows = sqlx::query!(
            "SELECT site_url,
        JSON_ARRAYAGG(title) as titles,
        JSON_ARRAYAGG(description) as descriptions
        FROM feeds GROUP BY site_url;"
        )
        .fetch_all(&self.pool)
        .await
        .map_err(|e| anyhow::anyhow!(e));

        if let Err(e) = rows {
            println!("Failed to fetch all feeds: {}", e);
            return Err(Error::RowNotFound);
        }

        let grouped_urls = rows
            .unwrap()
            .iter()
            .map(|row| CooccurrenceNetworkPoolForRead {
                site_url: row.site_url.clone(),
                titles: json!(row.titles),
                descriptions: json!(row.descriptions),
            })
            .collect();

        Ok(grouped_urls)
    }

    async fn insert_site_url_group(
        &self,
        grouped_urls: Vec<CooccurrenceNetworkPoolForRead>,
    ) -> Result<(), Error> {
        self.pool.begin().await?;

        for one_site in grouped_urls {
            let uuid = uuid::Uuid::new_v4();
            let mut tx = self.pool.begin().await?;
            let result = sqlx::query!(
                "INSERT INTO cooccurrence_network_pools (id, site_url, titles, descriptions)
                 VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE site_url = ?;
                 ",
                uuid.to_string(),
                one_site.site_url,
                one_site.titles,
                one_site.descriptions,
                one_site.site_url
            )
            .execute(&mut tx)
            .await;

            if let Err(e) = result {
                println!("Failed to insert all feeds: {:?}", e);
                return Err(e);
            }

            tx.commit().await?;
        }

        Ok(())
    }
}
