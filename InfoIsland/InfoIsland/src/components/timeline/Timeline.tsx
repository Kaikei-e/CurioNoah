import {Container, Flex, Link, Text} from "@chakra-ui/react";
import {ExternalLinkIcon} from '@chakra-ui/icons'
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
                                <Flex h={"20%"} w={"100%"}
                                      flexDirection={"row"}>
                                    <Link href={feed.link} w={"50%"} h={"100%"}
                                          overflowWrap={"break-word"} isExternal>
                                        <Text fontSize={{
                                            base: "xl",
                                            md: "xl",
                                            lg: "xl"
                                        }}>
                                            {/*&#x1F4DC; {feed.title}*/}
                                            &#x1F58B; {feed.title}
                                            <ExternalLinkIcon mx={"2px"}/>
                                        </Text>
                                    </Link>
                                    {/*<Flex w={"100%"} h={"100%"}>*/}
                                    {/*    <Suspense fallback={<Spinner/>}>*/}
                                    {/*        <ImageRenderer url={feed.image.url}/>*/}
                                    {/*    </Suspense>*/}
                                    {/*</Flex>*/}

                                    <Text w={"30%"} h={"100%"}>
                                        {feed.published}
                                    </Text>
                                    {/*<Link w={"20%"} h={"100%"} href={"/insight-stream"}>More*/}
                                    {/*    feeds ={">"}</Link>*/}
                                </Flex>
                                <Flex flexDirection={"column"} p={"1%"}
                                      border={"1px"} borderRadius={"md"}
                                      bgColor={"#fffbfb"}>
                                    {feed.items.map((item: Item, index: number) => {
                                        if (index > 5) {
                                            return;
                                        }
                                        return (
                                            <Flex flexDirection={"column"} key={index}>
                                                <Link href={item.link} isExternal>
                                                    <Text>
                                                        {/*&#x1F4CC;: {item.title}*/}
                                                        &#x1F587;: {item.title}
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

// const ImageRenderer = (props: {
//     url: string
// }) => {
//     return (
//         // <Image src={base64Img} boxSize='30%' rounded={"xl"}/>
//         <Image src={props.url} boxSize='40%' rounded={"xl"}/>
//     )
// }
export default Timeline;