import {
  CircularProgress,
  Container,
  Flex,
  Link,
  Text,
} from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import React from "react";
import { Feed, Item } from "../../lib/models/feedModel";
import EachFeed from "./eachFeed";
import InfiniteScroll from "react-infinite-scroller";

function Timeline(props: {
  data: Feed[];
  loadMoreFeeds: (page: number) => void;
  hasMore: boolean;
}) {
    const loader = <CircularProgress isIndeterminate color="green.300" />;

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
            <Link href={"/insight-stream"}>InsightStream</Link>
          </Text>{" "}
          page.
        </Text>
      </Flex>
      <InfiniteScroll
          pageStart={0}
          loadMore={props.loadMoreFeeds}
        hasMore={props.hasMore}
        loader={loader}
      >
        {props.data.map((feed: Feed, index: number) => {
          return <EachFeed feed={feed} index={index} />;
        })}
      </InfiniteScroll>
    </Flex>
  );
}

export default Timeline;
