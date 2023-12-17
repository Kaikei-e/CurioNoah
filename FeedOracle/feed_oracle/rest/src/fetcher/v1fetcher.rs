use axum::http::Error;

trait Fetcher {
    fn fetch_feed_link(&self, url: &str) -> Result<String, Error>;
}

struct V1Fetcher;

impl Fetcher for V1Fetcher {
    fn fetch_feed_link(&self, url: &str) -> Result<String, Error> {
        todo!("fetch feed link")
    }
}
