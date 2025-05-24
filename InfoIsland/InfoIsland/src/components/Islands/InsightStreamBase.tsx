import React, { useEffect, useState } from "react";
import { Flex } from "@chakra-ui/react";
import { CircularProgress } from "@chakra-ui/progress";
import { Timeline } from "./timeline/Timeline";
import type {
  FollowingSiteFeeds,
  FollowingSiteFeedsList,
} from "../../lib/models/feedModel";
import { Feeds } from "../../lib/models/eachFeed";

export const InsightStreamBase = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [followListFeeds, setFollowListFeeds] = React.useState<
    FollowingSiteFeeds[]
  >([]);
  const [error, setError] = React.useState<Error>();
  const [hasMoreItems, setHasMoreItems] = React.useState(true);

  const fetchMoreFeeds = async (
    page: number,
  ): Promise<FollowingSiteFeedsList | null> => {
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
        },
      );

      if (response.ok) {
        return await response.json();
      }
      throw new Error("Failed to fetch data");
    } catch (error) {
      console.error(error);
      return null;
    }
  };

  const loadMore = async (page: number) => {
    const result = await fetchMoreFeeds(page);

    if (result) {
      if (result.feeds && result.feeds.length > 0) {
        setFollowListFeeds((prevItems) => [...prevItems, ...result.feeds]);
        setHasMoreItems(!result.hadExceeded);
      } else {
        setHasMoreItems(false);
      }
    }
  };

  const loader = <CircularProgress isIndeterminate color="green.300" />;
  const noMoreItems = "No more items. Add more feeds to your list.";

  return (
    <Flex
      flexDirection={"column"}
      w={"100%"}
      h={"100%"}
      bgColor={"#EAF2F8"}
      rounded={"xl"}
    >
      <Timeline
        data={followListFeeds}
        isLoading={hasMoreItems}
        loadMore={loadMore}
        loader={loader}
        hadExceeded={hasMoreItems}
      />
    </Flex>
  );
};

// function arrayEquals(old: Feed[], newF: Feed[]) {
//   return (
//     Array.isArray(old) &&
//     Array.isArray(newF) &&
//     old.length === newF.length &&
//     old.every((val, index) => val === newF[index])
//   );
// }
