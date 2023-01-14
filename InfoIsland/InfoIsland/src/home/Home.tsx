import {Box, Center, Container, Flex, Text} from "@chakra-ui/react";

const Home = () => {
    return (
        <Flex bgColor={'#EAF2F8'}
              w={'100vw'}
              h={'100vh'}
              minH={'100%'}
              minW={'100%'}
              flexDir={'row'}
              p={'2%'}>
            <Flex flexDirection={"column"} w={"50%"} h={"100%"}>
                <Text fontFamily="Jost"
                      fontWeight="regular"
                      fontSize={{base: "xl", md: "2xl", lg: "5xl"}}
                      color="#000"
                      m={"2%"}>
                    Home
                </Text>
                <Flex w={"100%"} h={"40%"}
                      m={"2%"} justifyContent={"center"}
                >
                    <Flex w={"90%"} h={"100%"} bgColor={"#CDFFD7"} rounded={"xl"}
                          justifyContent={"center"} alignItems={"center"}>
                        <Flex justifyContent={"start"} w={"100%"} h={"100%"} p={"1%"}>
                            <Box>
                                <Text
                                    fontWeight="regular"
                                    fontSize={{base: "xl", md: "2xl", lg: "2xl"}}
                                    color="#000">
                                    Curiosity
                                </Text>
                            </Box>
                        </Flex>
                    </Flex>
                </Flex>
                <Flex w={"100%"} h={"40%"}
                      m={"2%"} justifyContent={"center"}
                >
                    <Flex justifyContent={"start"} w={"90%"} h={"100%"} flexDirection={"row"}>
                        <Flex bgColor={"#FFE4C3"}
                              w={"48%"} h={"100%"}
                              p={"1%"} mr={"2%"} rounded={"xl"}
                        >
                            <Flex justifyContent={"start"} alignItems={"start"}>
                                <Text
                                    fontWeight="regular"
                                    fontSize={{base: "xl", md: "2xl", lg: "2xl"}}
                                    color="#000">
                                    FeedFlare: Feed timeline
                                </Text>
                            </Flex>
                        </Flex>
                        <Flex flexDirection={"column"}
                              bgColor={"#FFE4C3"}
                              w={"48%"} h={"100%"}
                              p={"1%"} ml={"2%"} rounded={"xl"}
                              justifyContent={"center"} alignItems={"center"}>
                            <Box bgColor={"#D5D5D5"}
                                 w={"100%"} h={"48%"} mb={"2%"}>
                                <Text>
                                    SlateFlex: Notes
                                </Text>
                            </Box>
                            <Box bgColor={"#D5D5D5"}
                                 w={"100%"} h={"48%"} mt={"2%"}>
                                <Text>
                                    Account Settings
                                </Text>
                            </Box>
                        </Flex>
                    </Flex>
                </Flex>
            </Flex>

        </Flex>

    )
};

export default Home;