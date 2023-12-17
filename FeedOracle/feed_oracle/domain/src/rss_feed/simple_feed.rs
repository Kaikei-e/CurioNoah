struct SimpleFeed {
    pub title: String,
    pub description: String,
    pub link: String,
    pub items: Vec<SimpleFeedItem>,
}

struct SimpleFeedItem {
    pub title: String,
    pub description: String,
    pub link: String,
}
