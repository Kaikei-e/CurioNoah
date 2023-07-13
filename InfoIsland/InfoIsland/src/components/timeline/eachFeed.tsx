import { Container, Flex, Link, Text } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import { FollowingSiteFeeds, Item } from "../../lib/models/feedModel";
import React from "react";

type EachFeedProps = {
  feed: FollowingSiteFeeds;
  index: number;
};

export const EachFeed = (props: EachFeedProps) => {
  return (
    <Flex flexDirection={"row"} key={props.index} m={"1%"}>
      <Flex flexDirection={"row"} w={"100%"} h={"100%"}>
        <Container maxW={"100%"} h={"100%"} color={"#000"} p={"1%"}>
          <Flex h={"20%"} w={"100%"} flexDirection={"row"}>
            <Link
              href={props.feed.link}
              w={"50%"}
              h={"100%"}
              overflowWrap={"break-word"}
              isExternal
            >
              <Text
                fontSize={{
                  base: "xl",
                  md: "xl",
                  lg: "xl",
                }}
              >
                &#x1F58B; {props.feed.title}
                <ExternalLinkIcon mx={"2px"} />
              </Text>
            </Link>
            <Text w={"30%"} h={"100%"}>
              {props.feed.published}
            </Text>
          </Flex>
          <Flex
            flexDirection={"column"}
            p={"1%"}
            border={"1px"}
            borderRadius={"md"}
            bgColor={"#fffbfb"}
            mb={"1%"}
          >
            {props.feed.items.map((item: Item, index: number) => {
              if (index > 3) {
                return;
              }
              return (
                <Flex flexDirection={"column"} key={index} mb={"1%"}>
                  <Link href={item.link} isExternal>
                    <Text>
                      &#x1F587;: {item.title}
                      <ExternalLinkIcon mx={"2px"} />/{" "}
                      {item.description
                        ? item.description.slice(0, 40)
                        : "No Description"}{" "}
                      ...
                    </Text>
                  </Link>
                </Flex>
              );
            })}
          </Flex>
        </Container>
      </Flex>
    </Flex>
  );
};

