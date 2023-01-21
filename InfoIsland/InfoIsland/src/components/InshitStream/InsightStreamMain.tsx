import {
    Flex, Text, Table,
    Thead, Tbody,
    Tr, Th, Td,
    TableCaption,
    TableContainer,
} from "@chakra-ui/react";
import React from "react";
import AddFeed from "./AddFeed";

const InsightStreamMain = () => {
    return (
        <Flex h={"100vh"} w={"100vw"}
              minH={"100%"} minW={"100%"}
              overflowX={"hidden"}
              flexDirection={"row"}
              fontFamily="Jost"
        >
            <Flex flexDirection={"column"} w={"20%"} h={"100%"} bgColor={"#88A6B1"}>
                <Flex h={"6%"}
                      justifyContent={"center"} alignItems={"center"}
                      rounded={"xl"} p={"4%"}
                      textAlign={"center"}
                >
                    <Text
                        style={{
                            borderBottom: "2px solid #000"
                        }}
                        fontWeight="regular"
                        fontSize={{base: "lg", md: "xl", lg: "4xl"}}
                        color="#fff"
                        textAlign={"start"}>
                        InsightStream
                    </Text>
                </Flex>
                <Flex flexDirection={"column"}
                      w={"100%"} h={"94%"}
                      alignItems={"center"}
                      p={"4%"}
                >
                    <Flex flexDirection={"column"}
                          w={"80%"} h={"60%"}
                          bgColor={"#EAF2F8"}
                          rounded={"xl"}
                          border={"1px solid #000"}
                    >
                        <Text fontWeight="regular"
                              fontSize={{base: "xl", md: "xl", lg: "2xl"}}
                              color="#000" m={"2%"}
                              textAlign={"center"}>
                            Status
                        </Text>
                        <TableContainer w={"100%"}>
                            <Table variant='simple'>
                                <TableCaption>You can find the stats of reading</TableCaption>
                                <Thead>
                                    <Tr>
                                        <Th>Topic</Th>
                                        <Th isNumeric>Statistic</Th>
                                    </Tr>
                                </Thead>
                                <Tbody>
                                    <Tr>
                                        <Td>Following</Td>
                                        <Td isNumeric>20</Td>
                                    </Tr>
                                    <Tr>
                                        <Td>Number of feeds <br/>read by now</Td>
                                        <Td isNumeric>100</Td>
                                    </Tr>
                                </Tbody>
                            </Table>
                        </TableContainer>
                    </Flex>
                </Flex>
            </Flex>
            <Flex flexDirection={"row"}
                  w={"80%"} h={"100%"}
                  bgColor={"#88A6B1"} p={"2%"}>
                <Flex flexDirection={"column"}
                      w={"70%"} h={"100%"}
                      bgColor={"#EAF2F8"}
                      rounded={"xl"} p={"2%"}
                >
                </Flex>
                <Flex w={"25%"} h={"40%"}
                      rounded={"xl"} p={"2%"} ml={"5%"}
                      border={"1px solid #000"} shadow={"md"}
                >
                    <AddFeed/>
                </Flex>
            </Flex>

        </Flex>
    );
}

export default InsightStreamMain
