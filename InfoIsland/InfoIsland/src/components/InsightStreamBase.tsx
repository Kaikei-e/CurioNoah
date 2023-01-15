import React from 'react';
import {Flex, Box, Text, Container} from "@chakra-ui/react";


const InsightStreamBase = () => {
    return (
        <Flex flexDirection={"column"} h={"100%"} w={"100%"}
              bgColor={"#EAF2F8"} rounded={"xl"}
              overflow={"scroll"} overflowX={"hidden"}>
            {dummyFeeds.map((feed, index) => {
                return (
                    <Flex flexDirection={"row"} key={index} m={"1%"}>
                        <Flex flexDirection={"row"} w={"100%"} h={"100%"}>
                            <Container maxW={"100%"} h={"100%"}>
                                <Text fontSize={{
                                    base: "xl",
                                    md: "2xl",
                                    lg: "2xl"
                                }}>
                                    {feed.name}
                                </Text>
                                <Text>
                                    {feed.description}
                                </Text>
                                <Text>
                                    {feed.url}
                                </Text>
                            </Container>
                        </Flex>
                    </Flex>
                )

            })}
        </Flex>
    );

};

const dummyFeeds = [
    {
        name: "Dummy feed Curiosity",
        description: "Curiosity is a personal feed aggregator that allows you to collect and organize your favorite feeds in one place.",
        url: "https://curiosity.feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed Curiosity",
        description: "Curiosity is a personal feed aggregator that allows you to collect and organize your favorite feeds in one place.",
        url: "https://curiosity.feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    }, {
        name: "Dummy feed Curiosity",
        description: "Curiosity is a personal feed aggregator that allows you to collect and organize your favorite feeds in one place.",
        url: "https://curiosity.feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
    {
        name: "Dummy feed FeedFlare: Feed Collector",
        description: "FeedFlare is a feed collector that allows you to collect and organize your favorite feeds in one place.",
        url: "https://feedflare.com",
    },
]


export default InsightStreamBase;
