import React from "react";
import {
  SummarizedFeed,
  SummarizedFeeds,
} from "../../../lib/models/summarizedFeeds";
<<<<<<< HEAD
import { Flex, Link, Text } from "@chakra-ui/react";
import { CircularProgress } from "@chakra-ui/progress";
=======
import { Flex, Link as ChakraLink, Text } from "@chakra-ui/react";
>>>>>>> dca0000 (refactor: update component styles and structure, add new UI components, and improve API function signatures)
import InfiniteScroll from "react-infinite-scroller";

export const SummarizeTodayFeeds: React.FC = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [items, setItems] = React.useState<SummarizedFeed[]>([]);
  const [hasMoreItems, setHasMoreItems] = React.useState(true);

  let data: SummarizedFeeds | null = null;
  const fetchItems = async (page: number): Promise<SummarizedFeeds | null> => {
    try {
      const response = await fetch(
        `${apiURL}/infinite-scroll/summarize/today?page=${page}`,
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
        data = await response.json();
        return data;
      } else {
        return null;
      }
    } catch (e) {
      console.error(e);
      return null;
    }
  };

  const loadMoreItems = async (page: number) => {
    const result = await fetchItems(page);
    if (result) {
      if (result.summarizedFeeds && result.summarizedFeeds.length > 0) {
        setItems((prevItems) => [...prevItems, ...result.summarizedFeeds]);
        setHasMoreItems(!result.hasExceeded);
      } else {
        setHasMoreItems(false);
      }
    } else {
      setHasMoreItems(false);
    }
  };

  const loader = <CircularProgress isIndeterminate color={"green.200"} />;

  return (
    <div
      style={{
        width: "100%",
        overflowY: "scroll",
        borderRadius: "2xl",
      }}
    >
      <InfiniteScroll
        pageStart={0}
        loadMore={loadMoreItems}
        hasMore={hasMoreItems}
        loader={loader}
        useWindow={false}
      >
        {items.map((item, index) => {
          return (
            <Flex
              key={item.feed_url + index}
              p={2}
              my={2}
              bgColor={"blackAlpha.50"}
              boxShadow="md"
              borderWidth="1px"
              borderRadius="lg"
              overflow="hidden"
              _hover={{
                boxShadow: "xl",
              }}
            >
              <Flex
                flexDirection={"row"}
                alignItems={"center"}
                justifyContent={"space-between"}
                width={"100%"}
              >
                <ChakraLink
                  href={item.feed_url}
                  fontSize="sm"
                  fontWeight="bold"
                >
                  {item.title}
                </ChakraLink>
                <Text>{item.description}</Text>
              </Flex>
            </Flex>
          );
        })}
      </InfiniteScroll>
    </div>
  );
};
