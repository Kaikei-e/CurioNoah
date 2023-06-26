import React from "react";
import InfiniteScroll from "react-infinite-scroller";
import { CircularProgress } from "@chakra-ui/react";
import { EachFeed, Feeds } from "../../../lib/models/eachFeed";

type Props = {};

const InfiniteFeeds: React.FC<Props> = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [items, setItems] = React.useState<EachFeed[]>([]);
  const [hasMoreItems, setHasMoreItems] = React.useState(true);

  const fetchFeeds = async (page: number): Promise<Feeds | null> => {
    try {
      const response = await fetch(
        `${apiURL}/fetch-feed/stored-all?page=${page}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
            Origin: origin,
          },
        }
      );

      if (response.ok) {
        return await response.json();
      } else {
        throw new Error("Failed to fetch data");
      }
    } catch (error) {
      console.error(error);
      return null;
    }
  };

  const loadMore = async (page: number) => {
    const result = await fetchFeeds(page);

    if (result) {
      if (result.feeds && result.feeds.length > 0) {
        setItems((prevItems) => [...prevItems, ...result.feeds]);
        setHasMoreItems(!result.hadExceeded);
      } else {
        setHasMoreItems(false);
      }
    } else {
      setHasMoreItems(false);
    }
  };

  const loader = <CircularProgress isIndeterminate color="green.300" />;

  return (
    <div>
      <InfiniteScroll
        pageStart={0}
        loadMore={loadMore}
        hasMore={hasMoreItems}
        loader={loader}
      >
        {items.map((item, index) => (
          <div key={index}>
            {/* Here you should use your component to display each feed */}
            {/* <FeedItem feed={item} /> */}
          </div>
        ))}
      </InfiniteScroll>
    </div>
  );
};
