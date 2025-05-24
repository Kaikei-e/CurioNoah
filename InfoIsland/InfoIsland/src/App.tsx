import "./App.css";
import { Box, Flex, VStack, Spacer, Text, Link, Button } from "@chakra-ui/react";

export const App = () => (
  <Box
    w={"100vw"}
    h={"100vh"}
    minW={"100%"}
    minH={"100%"}
    backgroundImage={"/images/ship-sunset.jpg"}
    backgroundSize={"cover"}
    backgroundRepeat={"no-repeat"}
    p={10}
  >
    <Flex minWidth="max-content" h={"100%"} w={"100%"}>
      <VStack
        align="flex-start"
        spaceX={4}
        spaceY={4}
      >
        <Text
          fontFamily="Jost"
          fontWeight="regular"
          fontSize={{ base: "3xl", md: "4xl", lg: "80px" }}
          color="#FFF"
          width={{ base: "100%", md: "100%", lg: "100%" }}
          maxWidth="100%"
        >
          <span>CurioNoah</span>
          <br />
          <Box
            as="span"
            fontWeight="light"
            fontSize={{ base: "xl", md: "2xl", lg: "6xl" }}
          >
            Explore, Store, Extend
          </Box>
          <br />
          <Box
            as={"span"}
            fontWeight="light"
            fontSize={{ base: "xl", md: "2xl", lg: "6xl" }}
          >
            : by curiosity.
          </Box>
        </Text>
      </VStack>
      <Spacer />
      <Flex flexDir={"column"} justifyContent={"flex-end"} p="4%">
        <Link href="/login">
          <Button
            size={{ base: "lg", sm: "sm", md: "md", lg: "lg" }}
            variant="outline"
            color={"#FFFFFF"}
            _hover={{ color: "#000000", bg: "#FFFFFF" }}
            rightIcon={
              <ArrowForwardIcon
                boxSize={{ base: "xl", md: "2xl", lg: "20px" }}
              />
            }
          >
            <Text
              fontFamily="Jost"
              fontSize={{ base: "md", md: "2md", lg: "2xl" }}
            >
              Explore
            </Text>
          </Button>
        </Link>
      </Flex>
    </Flex>
  </Box>
);
