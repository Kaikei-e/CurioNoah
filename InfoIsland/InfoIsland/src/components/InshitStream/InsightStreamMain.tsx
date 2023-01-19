import {
    Flex, Text, Table,
    Thead, Tbody, Tfoot,
    Tr, Th, Td,
    TableCaption,
    TableContainer,
} from "@chakra-ui/react";
import React from "react";

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
                        <TableContainer>
                            <Table variant='simple'>
                                <TableCaption>Imperial to metric conversion factors</TableCaption>
                                <Thead>
                                    <Tr>
                                        <Th>To convert</Th>
                                        <Th>into</Th>
                                        <Th isNumeric>multiply by</Th>
                                    </Tr>
                                </Thead>
                                <Tbody>
                                    <Tr>
                                        <Td>inches</Td>
                                        <Td>millimetres (mm)</Td>
                                        <Td isNumeric>25.4</Td>
                                    </Tr>
                                    <Tr>
                                        <Td>feet</Td>
                                        <Td>centimetres (cm)</Td>
                                        <Td isNumeric>30.48</Td>
                                    </Tr>
                                    <Tr>
                                        <Td>yards</Td>
                                        <Td>metres (m)</Td>
                                        <Td isNumeric>0.91444</Td>
                                    </Tr>
                                </Tbody>
                                <Tfoot>
                                    <Tr>
                                        <Th>To convert</Th>
                                        <Th>into</Th>
                                        <Th isNumeric>multiply by</Th>
                                    </Tr>
                                </Tfoot>
                            </Table>
                        </TableContainer>
                    </Flex>

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
