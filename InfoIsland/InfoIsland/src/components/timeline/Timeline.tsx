import {
  Button,
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

function Timeline(props: {
  data: Feed[];
  hasMore: boolean;
  isLoading: boolean;
  fetchMoreFeeds: () => void;
}) {
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
      <Flex flexDirection={"column"} h={"100%"} w={"100%"}>
        {props.data.map((feed: Feed, index: number) => {
          return <EachFeed feed={feed} index={index} />;
        })}
      </Flex>
      {props.hasMore ? (
        <Button onClick={props.fetchMoreFeeds} isLoading={props.isLoading}>
          Go to next page
        </Button>
      ) : (
        <Text>No more feeds</Text>
      )}
    </Flex>
  );
}

export default Timeline;
