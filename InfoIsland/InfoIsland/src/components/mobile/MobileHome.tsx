import { Flex, Text } from "@chakra-ui/react";
import { Timeline } from "./Timeline";
import { Route } from "react-router-dom";
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
      <Flex flexDirection={"column"}>
        <a href={"mobile/register"}>
          <Text textAlign={"center"} p={"0.5rem"}>
            To Mobile
          </Text>
        </a>
        <Timeline />
      </Flex>
    </Flex>
  );
};
