import React from "react";
import InfiniteScroll from "react-infinite-scroller";
import {CircularProgress} from "@chakra-ui/react";
import {Feed} from "../../../lib/models/feedModel";

type Props = {};

const InfiniteFeeds: React.FC<Props> = () => {
    const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
    const origin = import.meta.env.VITE_ORIGIN;

    const [items, setItems] = React.useState([]);
    const [hasMoreItems, setHasMoreItems] = React.useState(true);

    const loadMore = (page: number) => {
        setItems(...items, ...page);
    }

    const loader = <CircularProgress isIndeterminate color="green.300" />;

  return (
    <div>
      <InfiniteScroll
        loadMore={loadMore}
        hasMore={hasMoreItems}
        loader={loader}
      >
        {items}
      </InfiniteScroll>
    </div>
  );
};
