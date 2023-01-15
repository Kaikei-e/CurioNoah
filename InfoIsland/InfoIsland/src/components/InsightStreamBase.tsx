import React, {useEffect} from 'react';
import {Flex, Text, Container, Spinner} from "@chakra-ui/react";

// make this function to return boolean
const InsightStreamBase = () => {

    const [data, setData] = React.useState([]);
    const [isLoading, setIsLoading] = React.useState(true);
    const [error, setError] = React.useState(null);

    useEffect(() => {
        const feeds = async () => {
            const response = await fetch(
                'http://localhost:9000/api/v1/fetch-feeds',
                {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json',
                        // 'Access-Control-Allow-Origin': '*',
                        'Origin': 'http://localhost:5173'
                    },
                },
            );
            const feeds = await response.json();
            setData(feeds);
            setIsLoading(false);
        };

        feeds().catch((error) => {
            setError(error);
            setIsLoading(false);
        });
    }, []);

    let displayData;

    if (!isLoading) {
        displayData = <TimeLine data={data}/>
    } else if (error) {
        displayData = <div>Something went wrong ...</div>
    } else {
        displayData = <FetchingFeeds/>
    }

    return (
        <Flex flexDirection={"column"} w={"100%"} h={"100%"}
              bgColor={"#EAF2F8"} rounded={"xl"}
              overflow={"scroll"} overflowX={"hidden"}>
            {displayData}
        </Flex>
    );
};

function TimeLine(props) {
    return (
        <Flex flexDirection={"column"} h={"100%"} w={"100%"}
        >
            {props.data.map((feed, index: number) => {
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
                                <Flex>
                                    {feed.items.map((item, index: number) => {

                                        if (index > 5) {
                                            return;
                                        }
                                        return (
                                            <Flex flexDirection={"column"} key={index}>
                                                <Text>
                                                    {item.title}
                                                </Text>
                                            </Flex>
                                        )
                                    })}
                                </Flex>
                            </Container>
                        </Flex>
                    </Flex>
                )
            })}
        </Flex>
    );
}

function FetchingFeeds() {
    return (
        <Flex flexDirection={"column"}
              w={"100%"} h={"100%"}>
            <Spinner
                thickness='4px'
                speed='0.65s'
                emptyColor='gray.200'
                color='blue.500'
                size='xl'
            />
        </Flex>
    );
};


export default InsightStreamBase;
