import { Alert, AlertIcon, Button, Flex, Link, Text } from "@chakra-ui/react";
import React from "react";
import { Feed } from "../../lib/models/feedModel";
import EachFeed from "./eachFeed";

function Timeline(props: {
  data: Feed[];
  isLoading: boolean;
  fetchMoreFeeds: () => void;
  error: Error | undefined;
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
      <Flex flexDirection={"column"} h={"90%"} w={"100%"} overflowY={"scroll"}>
        {props.data.map((feed: Feed, index: number) => {
          return <EachFeed feed={feed} index={index} />;
        })}
      </Flex>
      {props.error && (
        <Alert status="error">
          <AlertIcon />
          {props.error.message}
        </Alert>
      )}
      <Flex
        h={"5%"}
        w={"100%"}
        justifyContent={"center"}
        alignItems={"center"}
      >
        <Button onClick={props.fetchMoreFeeds} isLoading={props.isLoading}>
          Load more
        </Button>
      </Flex>
    </Flex>
  );
}

export default Timeline;
