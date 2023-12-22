import { Flex } from "@chakra-ui/react";
import { Timeline } from "./Timeline";

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
      <Timeline />
    </Flex>
  );
};
