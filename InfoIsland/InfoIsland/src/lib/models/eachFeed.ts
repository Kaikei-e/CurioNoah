export interface Feeds {
  feeds: EachFeed[];
  hadExceeded: boolean;
}

export interface EachFeed {
  item_id: string;
  site_url: string;
  title: string;
  description: string;
  feed_url: string;
  language?: Language;
  dt_created: string;
  dt_updated: string;
  favorites: number;
}

export enum Language {
  En = "en",
  EnUS = "en-US",
  EnUs = "en-us",
  Ja = "ja",
}
