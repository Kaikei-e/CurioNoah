import React from "react";
import InfiniteScroll from "react-infinite-scroller";
import {
  Box,
  Button,
  CircularProgress,
  Flex,
  Link,
  Text,
} from "@chakra-ui/react";
import { EachFeed, Feeds } from "../../../../lib/models/eachFeed";

type Props = {};

export const InfiniteFeeds: React.FC<Props> = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const [items, setItems] = React.useState<EachFeed[]>([]);
  const [hasMoreItems, setHasMoreItems] = React.useState(true);

  const fetchFeeds = async (page: number): Promise<Feeds | null> => {
    try {
      const response = await fetch(
        `${apiURL}/infinite-scroll/stored-all?page=${page}`,
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
        return await response.json();
      } else {
        throw new Error("Failed to fetch data");
      }
    } catch (error) {
      console.error(error);
      return null;
    }
  };

  const loadMore = async (page: number) => {
    const result = await fetchFeeds(page);

    if (result) {
      if (result.feeds && result.feeds.length > 0) {
        setItems((prevItems) => [...prevItems, ...result.feeds]);
        setHasMoreItems(!result.hadExceeded);
      } else {
        setHasMoreItems(false);
      }
    } else {
      setHasMoreItems(false);
    }
  };

  const loader = <CircularProgress isIndeterminate color="green.300" />;

  return (
    <InfiniteScroll
      pageStart={0}
      loadMore={loadMore}
      hasMore={hasMoreItems}
      loader={loader}
      useWindow={false}
    >
      <Flex flexDirection={"column"}>
        {items.map((item, index) => (
          <Box
            key={item.title + index}
            p={5}
            my={2}
            bgColor={"teal.50"}
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
            >
              <Link
                href={item.feed_url}
                isExternal
                fontSize="lg"
                fontWeight="bold"
              >
                {item.title}
              </Link>
              <Button
                  variant='outline'
                  border={"1px"}
                  color={"blackAlpha.800"} backgroundColor={"teal.50"}>
                Like !
              </Button>
            </Flex>
            <Text mt={2}>{item.description}</Text>
          </Box>
        ))}
      </Flex>
    </InfiniteScroll>
  );
};
