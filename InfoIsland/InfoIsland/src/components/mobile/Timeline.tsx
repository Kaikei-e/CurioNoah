import { Flex, Text } from "@chakra-ui/react";
import { InfiniteFeeds } from "./InfiniteScroll/InfiniteFeeds";

export const Timeline = () => {
  return (
      <Flex
        w={"100%"}
        h={"100%"}
        bgColor={"#EAF2F8"}
        rounded={"xl"}
        p={"0.5rem"}
        border={"1px solid #000"}
        flexDirection={"column"}
      >
        <div style={{ width: "100%", height: "8%" }}>
          <Text textAlign={"center"}>Timeline</Text>
        </div>
        <Flex width={"100%"} height={"85%"}>
          <InfiniteFeeds />
      </Flex>
    </Flex>
  );
};