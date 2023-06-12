use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::FeedElement;
use crate::domain::Feed;
use crate::driver::repository::feed_db::connect::{FeedConnection, FeedRepository};
use crate::usecase::fetch_all_follow_list::fetch_all_follow_list;
use anyhow::{Error, Result};
use chrono::Utc;

pub async fn parse_and_store_feeds(pool: DatabasePool) -> Result<(), Error> {
    let follow_list = fetch_all_follow_list(pool.clone()).await;
    if let Err(e) = follow_list {
        println!("Failed to fetch all follow list: {}", e);
        return Err(e);
    }

    let mut feed_list = Vec::<Feed>::new();

    for feed in follow_list.unwrap() {
        feed.item_description.iter().for_each(|item| {
            let feed_element = FeedElement {
                guid: item.guid.clone(),
                item_id: item.clone().item_id,
                updated: item.clone().updated,
                item_link: item.clone().item_link,
                published: item.clone().published,
                item_title: item.clone().item_title,
                item_description: item.clone().item_description,
                published_parsed: item.clone().published_parsed,
                categories: item.clone().categories,
            };
            let one_feed = Feed {
                id: uuid::Uuid::new_v4().to_string(),
                site_url: feed.clone().url,
                title: feed_element.clone().item_title,
                description: feed_element.clone().item_description.to_string(),
                created_at: Utc::now(),
                updated_at: Utc::now(),
                language: feed.clone().language,
                feed_url: feed_element.item_link,
                favorites: 0,
            };

            feed_list.push(one_feed);
        });
    }

    let feed_repository = FeedRepository::new(pool.clone());
    let result = feed_repository.insert_all_feeds(feed_list).await;
    if let Err(e) = result {
        println!("Failed to insert all feeds: {:?}", e);
        return Err(e.into());
    }

    Ok(())
}