use crate::api_handler::handler::DatabasePool;
use crate::domain::feed::FeedElement;
use crate::domain::Feed;
use crate::usecase::fetch_all_follow_list::fetch_all_follow_list;
use anyhow::{Error, Result};

pub async fn parse_and_store_feeds(pool: DatabasePool) -> Result<(), Error> {
    let follow_list = fetch_all_follow_list(pool).await;
    if let Err(e) = follow_list {
        println!("Failed to fetch all follow list: {}", e);
        return Err(e);
    }

    for feed in follow_list.unwrap() {
        println!("feed: {:?}", feed.links);
        // let one_feed: FeedElement = FeedElement{
        //     guid: feed.,
        //     item_id: "".to_string(),
        //     updated: "".to_string(),
        //     item_link: "".to_string(),
        //     published: "".to_string(),
        //     item_title: "".to_string(),
        //     item_description: "".to_string(),
        //     published_parsed: "".to_string(),
        //     categories: None,
        // };
    }

    Ok(())
}
