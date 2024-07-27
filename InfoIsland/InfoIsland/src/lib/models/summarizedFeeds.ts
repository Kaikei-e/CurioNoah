export interface SummarizedFeeds {
  summarizedFeeds: SummarizedFeed[];
  hasExceeded: boolean;
}

export interface SummarizedFeed {
  feed_url: string;
  title: string;
  description: string;
}
