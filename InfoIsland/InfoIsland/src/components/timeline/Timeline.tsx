import {Container, Flex, Text} from "@chakra-ui/react";
import React from "react";
import {Feed, Item} from "../../lib/models/feedModel";

function Timeline(props: { data: Feed[]; }) {

    return (

        <Flex flexDirection={"column"} h={"100%"} w={"100%"}
        >
            {props.data.map((feed: Feed, index: number) => {
                return (
                    <Flex flexDirection={"row"} key={index} m={"1%"}>
                        <Flex flexDirection={"row"} w={"100%"} h={"100%"}>
                            <Container maxW={"100%"} h={"100%"}
                                       color={"#000"} p={"1%"}>
                                <Text fontSize={{
                                    base: "xl",
                                    md: "2xl",
                                    lg: "2xl"
                                }}>
                                    {feed.title}
                                </Text>
                                <Text>
                                    {feed.published}
                                </Text>
                                <Text>
                                    {feed.link}
                                </Text>
                                <Flex flexDirection={"column"}>
                                    {feed.items.map((item: Item, index: number) => {

                                        if (index > 5) {
                                            return;
                                        }
                                        return (
                                            <Flex flexDirection={"column"} key={index}>
                                                <Text>
                                                    {item.link}
                                                </Text>
                                            </Flex>
                                        )
                                    })}
                                </Flex>
                            </Container>
                        </Flex>
                    </Flex>
                );
            })}
        </Flex>
    );
}

export default Timeline;