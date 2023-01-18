import React, {useEffect} from 'react';
import {Flex, Spinner} from "@chakra-ui/react";
import Timeline from "./timeline/Timeline";
import {Feed} from "../lib/models/feedModel";

// make this function to return boolean
const InsightStreamBase = () => {

    const [data, setData] = React.useState<Feed[]>([]);
    const [isLoading, setIsLoading] = React.useState<Boolean>(true);
    const [error, setError] = React.useState<Error>();

    useEffect(() => {
        const feeds = async () => {
            const response = await fetch(
                'http://localhost:9000/api/v1/fetch-feed/stored-all',
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
        displayData = <Timeline data={data}/>
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
}


export default InsightStreamBase;
