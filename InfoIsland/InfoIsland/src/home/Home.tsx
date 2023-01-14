import {Flex, Text} from "@chakra-ui/react";

const Home = () => {
    return (
        <Flex bgColor={'#EAF2F8'}
              w={'100vw'}
              h={'100vh'}
              minH={'100%'}
              minW={'100%'}
              p={'2%'}>
            <Text fontFamily="Jost"
                  fontWeight="regular"
                  fontSize={{base: "xl", md: "2xl", lg: "5xl"}}
                  color="#000">Home</Text>
        </Flex>

    )
};

export default Home;