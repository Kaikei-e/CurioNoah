use sqlx::{Error, MySql, Pool};
use crate::domain::cooccurrence_network::CooccurrenceNetworkPoolForRead;
use serde_json::{
    json
};

pub async fn fetch_site_url_group(pool: Pool<MySql>) -> Result<Vec<CooccurrenceNetworkPoolForRead>, Error> {
    let rows = sqlx::query!(
        "SELECT site_url,
        JSON_ARRAYAGG(title) as titles,
        JSON_ARRAYAGG(description) as descriptions
        FROM feeds GROUP BY site_url;"
    )
    .fetch_all(&pool)
    .await
    .map_err(|e| anyhow::anyhow!(e));

    if let Err(e) = rows {
        println!("Failed to fetch all feeds: {}", e);
        return Err(Error::RowNotFound);
    }

    let grouped_urls = rows.unwrap()
        .iter()
            .map(|row|
                CooccurrenceNetworkPoolForRead {
                    site_url: row.site_url.clone(),
                    titles: json!(row.titles),
                    descriptions: json!(row.descriptions),

            })
            .collect();

    Ok(grouped_urls)
}
