import { Button, Flex, Link, Text } from "@chakra-ui/react";

import { Alert, AlertIcon, AlertTitle } from "@chakra-ui/alert";
import { CircularProgress } from "@chakra-ui/progress";
import type React from "react";
import type { FollowingSiteFeeds } from "../../../lib/models/feedModel";
import { EachFeed } from "./eachFeed";
import InfiniteScroll from "react-infinite-scroller";
import { Simulate } from "react-dom/test-utils";
import { useNavigate } from "react-router-dom";

type Props = {
  data: FollowingSiteFeeds[];
  isLoading: boolean;
  loadMore: (page: number) => Promise<void>;
  hadExceeded: boolean;
  loader: React.ReactNode;
};

export const Timeline: React.FC<Props> = ({
  data,
  isLoading,
  loadMore,
  hadExceeded,
}) => {
  const navigate = useNavigate();
  const loader = (
    <Flex flexDirection={"column"} h={"85%"} w={"100%"}>
      <CircularProgress isIndeterminate color="green.300" />
    </Flex>
  );

  return (
    <Flex flexDirection={"column"} h={"100%"} w={"100%"} fontFamily="Jost">
      <Flex p={"2%"}>
        <Text
          fontSize={{
            base: "2xl",
            md: "xl",
            lg: "2xl",
          }}
        >
          You can find the latest feeds that you have collected here. Read more
          at{" "}
          <Text as="u">
            <Button onClick={() => navigate("/insight-stream")}>
              InsightStream
            </Button>
          </Text>{" "}
          page.
        </Text>
      </Flex>
      <Flex
        flexDirection={"column"}
        h={"85%"}
        w={"100%"}
        overflowY={"scroll"}
        borderY={"2px"}
        sx={{
          "&::-webkit-scrollbar": {
            width: "10px",
          },
          "&::-webkit-scrollbar-track": {
            background: "#f1f1f1",
          },
          "&::-webkit-scrollbar-thumb": {
            background: "#888",
            borderRadius: "5px",
          },
          "&::-webkit-scrollbar-thumb:hover": {
            background: "#555",
          },
        }}
      >
        <InfiniteScroll
          pageStart={0}
          loadMore={loadMore}
          hasMore={hadExceeded}
          loader={loader}
          useWindow={false}
        >
          {data.map((feed: FollowingSiteFeeds, index: number) => {
            return <EachFeed feed={feed} index={index} />;
          })}
        </InfiniteScroll>
      </Flex>
      {!hadExceeded && (
        <Alert status="error">
          <AlertIcon />
          {hadExceeded
            ? "Scroll down to load more feeds"
            : "No more feeds to load"}
        </Alert>
      )}
      {isLoading && (
        <Text
          textAlign={"center"}
          fontSize={{
            base: "2xl",
            md: "xl",
            lg: "2xl",
          }}
        >
          Scroll down to load more feeds
        </Text>
      )}
    </Flex>
  );
};
