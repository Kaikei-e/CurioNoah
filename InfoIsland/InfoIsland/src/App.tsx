import "./App.css";
import { Box, Flex, VStack, Spacer, Text, Link, Button, Icon } from "@chakra-ui/react";
import { FaArrowRight } from "react-icons/fa";

export const App = () => (
  <Box
    w="100vw"
    h="100vh"
    minW="100%"
    minH="100%"
    backgroundImage="/images/ship-sunset.jpg"
    backgroundSize="cover"
    backgroundRepeat="no-repeat"
    p={10}
  >
    <Flex minWidth="max-content" h="100%" w="100%">
      <VStack align="flex-start" gap={4}>
        <Text
          fontFamily="Jost"
          fontWeight="regular"
          fontSize={{ base: "3xl", md: "4xl", lg: "80px" }}
          color="#FFF"
          w="100%"
          maxW="100%"
        >
          CurioNoah
          <br />
          <Box as="span" fontWeight="light" fontSize={{ base: "xl", md: "2xl", lg: "6xl" }}>
            Explore, Store, Extend
          </Box>
          <br />
          <Box as="span" fontWeight="light" fontSize={{ base: "xl", md: "2xl", lg: "6xl" }}>
            : by curiosity.
          </Box>
        </Text>
      </VStack>
      <Spacer />
      <Flex flexDir="column" justifyContent="flex-end" p="4%">
        <Link href="/login">
          <Button
            size={{ base: "lg", sm: "sm", md: "md", lg: "lg" }}
            variant="outline"
            color="#FFFFFF"
            _hover={{ color: "#000000", bg: "#FFFFFF" }}
          >
            <Icon 
              as={FaArrowRight}
              color="#000000"
              boxSize={{ base: "1.5em", md: "1.5em", lg: "2em" }}
            />
          </Button>
        </Link>
      </Flex>
    </Flex>
  </Box>
);
