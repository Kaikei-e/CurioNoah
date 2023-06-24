import React, { useEffect, useState } from "react";
import { Flex, Spinner } from "@chakra-ui/react";
import Timeline from "./timeline/Timeline";
import { Feed } from "../lib/models/feedModel";
import EachFeed from "./timeline/eachFeed";

const InsightStreamBase = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [data, setData] = React.useState<Feed[]>([]);
  const [isLoading, setIsLoading] = React.useState<boolean>(true);
  const [error, setError] = React.useState<Error>();
  const [hasMore, setHasMore] = React.useState(true);
  const [page, setPage] = useState(0);

  const fetchMoreFeeds = async (page: number) => {
    setIsLoading(true);
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
      const result = await response.json();
      const feeds = result.feeds;

      if (!feeds || !Array.isArray(feeds) || feeds.length === 0) {
        setError(new Error("No feeds found"));
        setHasMore(false);
      } else {
        setData((prevData) => [...prevData, ...feeds]);
        setPage(page + 1);
      }
    } else {
      setError(new Error("Failed to fetch data"));
      setHasMore(false);
    }
    setIsLoading(false);
  };

  useEffect(() => {
    fetchMoreFeeds(page).catch((error) => {
      setError(error);
      setIsLoading(false);
    });
  }, [page]);

  let displayData;


  if (isLoading && data.length === 0) {
    displayData = <FetchingFeeds />;
  } else if (error) {
    displayData = <div>Something went wrong ...</div>;
  } else {
    displayData = (
        <Timeline
            data={data}
            hasMore={hasMore}
            isLoading={isLoading}
            fetchMoreFeeds={() => fetchMoreFeeds(page)}
        />
    );
  }

  return (
    <Flex
      flexDirection={"column"}
      w={"100%"}
      h={"100%"}
      bgColor={"#EAF2F8"}
      rounded={"xl"}
      overflow={"scroll"}
      overflowX={"hidden"}
    >
      {displayData}
    </Flex>
  );
};

function FetchingFeeds() {
  return (
    <Flex flexDirection={"column"} w={"100%"} h={"100%"}>
      <Spinner
        thickness="4px"
        speed="0.65s"
        emptyColor="gray.200"
        color="blue.500"
        size="xl"
      />
    </Flex>
  );
}

export default InsightStreamBase;
