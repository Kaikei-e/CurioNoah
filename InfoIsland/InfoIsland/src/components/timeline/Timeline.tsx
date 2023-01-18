import {Container, Flex, Image, Link, Spinner, Text} from "@chakra-ui/react";
import {ExternalLinkIcon} from '@chakra-ui/icons'
import React, {Suspense} from "react";
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
                                <Flex h={"20%"} w={"80%"}
                                      flexDirection={"row"}>
                                    <Link href={feed.link} w={"70%"} h={"100%"}
                                          overflowWrap={"break-word"} isExternal>
                                        <Text fontSize={{
                                            base: "xl",
                                            md: "2xl",
                                            lg: "2xl"
                                        }}>
                                            {feed.title}
                                            <ExternalLinkIcon mx={"2px"}/>
                                        </Text>
                                    </Link>
                                    <Flex w={"100%"} h={"100%"}>
                                        <Suspense fallback={<Spinner/>}>
                                            <ImageRenderer url={feed.image.url}/>
                                        </Suspense>
                                    </Flex>
                                </Flex>
                                <Text>
                                    {feed.published}
                                </Text>
                                <Flex flexDirection={"column"}>
                                    {feed.items.map((item: Item, index: number) => {
                                        if (index > 5) {
                                            return;
                                        }
                                        return (
                                            <Flex flexDirection={"column"} key={index}>
                                                <Link href={item.link} isExternal>
                                                    <Text>
                                                        {item.title}
                                                        <ExternalLinkIcon mx={"2px"}/>
                                                        / {item.description!.slice(0, 40)} ...
                                                    </Text>
                                                </Link>
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

const ImageRenderer = (props: {
    url: string
}) => {
    return (
        // <Image src={base64Img} boxSize='30%' rounded={"xl"}/>
        <Image src={props.url} boxSize='40%' rounded={"xl"}/>
    )
}
export default Timeline;