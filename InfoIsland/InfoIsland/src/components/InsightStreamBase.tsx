import React, { useEffect, useState } from "react";
import { Flex } from "@chakra-ui/react";
import Timeline from "./timeline/Timeline";
import { Feed } from "../lib/models/feedModel";

const InsightStreamBase = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [data, setData] = React.useState<Feed[]>([]);
  const [isLoading, setIsLoading] = React.useState<boolean>(true);
  const [error, setError] = React.useState<Error>();
  const [page, setPage] = useState(1);

  const fetchMoreFeeds = async () => {
    setIsLoading(true);
    const response = await fetch(`${apiURL}/fetch-feed/stored-all?page=${page}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Origin: origin,
      },
    });

    if (response.ok) {
      const result = await response.json();
      if (result.feeds && Array.isArray(result.feeds) && result.feeds.length > 0) {
        setData((prevData) => [...prevData, ...result.feeds]);
        setPage(page + 1);
      } else {
        setError(new Error("No more feeds available"));
      }
    } else {
      setError(new Error("Failed to fetch data"));
    }

    setIsLoading(false);
  };

  useEffect(() => {
    fetchMoreFeeds();
  }, []);

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
        <Timeline
            data={data}
            isLoading={isLoading}
            fetchMoreFeeds={fetchMoreFeeds}
            error={error}
        />
      </Flex>
  );
};

export default InsightStreamBase;
