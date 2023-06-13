use crate::api_handler::handler::DatabasePool;
use crate::driver::repository::data_pool::group_site_url;
use crate::driver::repository::data_pool::group_site_url::FetchSiteUrlGroup;
use anyhow::Error;

pub async fn store_word_net_usecase(pool: DatabasePool) -> Result<(), Error> {
    let word_network_repository = group_site_url::WordNetworkRepository::new(pool);
    let grouped_urls = word_network_repository.fetch_site_url_group().await?;

    let result = word_network_repository
        .insert_site_url_group(grouped_urls)
        .await;
    if let Err(e) = result {
        println!("Failed to store word network: {:?}", e);
        return Err(e.into());
    }

    Ok(())
}
