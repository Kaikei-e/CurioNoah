import {Flex, Text} from "@chakra-ui/react";
import React from "react";

const InsightStreamMain = () => {
    return (
        <Flex h={"100vh"} w={"100vw"}
              minH={"100%"} minW={"100%"}
              overflowX={"hidden"}
              flexDirection={"row"}
        >
            <Flex flexDirection={"column"} w={"20%"} h={"100%"} bgColor={"#88A6B1"}>
                <Flex h={"6%"}
                      justifyContent={"center"} alignItems={"center"}
                      rounded={"xl"} m={"5%"} p={"1%"}
                      textAlign={"center"}
                >
                    <Text fontFamily="Jost"
                          fontWeight="regular"
                          fontSize={{base: "lg", md: "xl", lg: "4xl"}}
                          color="#fff"
                          textAlign={"start"}>
                        InsightStream
                    </Text>
                </Flex>
                <Flex flexDirection={"column"}
                      w={"100%"} h={"94%"}
                      rounded={"xl"} p={"2%"}
                >
                </Flex>
            </Flex>
            <Flex flexDirection={"column"}
                  w={"80%"} h={"100%"}
                  bgColor={"#88A6B1"} p={"2%"}>
                <Flex flexDirection={"column"}
                      w={"70%"} h={"100%"}
                      bgColor={"#EAF2F8"}
                      rounded={"xl"} p={"2%"}
                >
                </Flex>
            </Flex>
        </Flex>
    );

}

export default InsightStreamMain
