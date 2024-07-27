import { Flex, Text } from "@chakra-ui/react";
import { Timeline } from "./Timeline";
import React from "react";

export const MobileHome = () => {
  return (
    <Flex
      w={"100vw"}
      h={"100vh"}
      minH={"100%"}
      minW={"100%"}
      flexDirection={"column"}
      fontFamily={"Jost"}
      backgroundColor={"blue.50"}
    >
      <a href={"mobile/register"}>
        <Text textAlign={"center"} p={"0.5rem"}>
          To Registering page
        </Text>
      </a>
      <Timeline />
    </Flex>
  );
};
