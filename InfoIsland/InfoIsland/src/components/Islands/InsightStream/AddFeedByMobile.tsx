import { Button, Flex, Input, Text } from "@chakra-ui/react";
import { CircularProgress } from "@chakra-ui/progress";
import { registeredFeed } from "../../../lib/models/apiExchange/registeredFeed";
import { RegisterFeed } from "../../../lib/feedFlare/registerFeed";
import { useState } from "react";

export const AddFeedByMobile = () => {
  const [feedUrl, setFeedUrl] = useState("");
  const [onLoading, setOnLoading] = useState(false);

  return (
    <Flex
      flexDirection={"column"}
      minH={"100vh"}
      minW={"100vw"}
      width={"100%"}
      height={"100%"}
      justifyContent={"center"}
      alignItems={"center"}
      textAlign={"center"}
    >
      {onLoading && (
        <CircularProgress isIndeterminate size="100px" thickness="4px" />
      )}

      {!onLoading && (
        <form
          method={"POST"}
          onSubmit={async (e) => {
            e.preventDefault();

            setOnLoading(true);

            const response: registeredFeed = await RegisterFeed(feedUrl);

            setOnLoading(false);
            if (response.message === "success") {
              alert("Feed registered successfully!");
            } else {
              alert("Feed registration failed!");
            }
          }}
        >
          <Input
            type="text"
            placeholder="Enter Feed URL"
            _placeholder={{ opacity: 0.65, color: "black" }}
            size={"lg"}
            w={"100%"}
            h={"100%"}
            mb={"2%"}
            value={feedUrl}
            onChange={(e) => setFeedUrl(e.target.value)}
          />
          <Button
            size={{ sm: "md", md: "md", lg: "md" }}
            variant="outline"
            textColor={"#FFF"}
            _hover={{ color: "#000", bg: "#FFF" }}
            rightIcon={<CheckIcon boxSize={{ md: "2em", lg: "2em" }} />}
            bgColor={"#84a3ad"}
            type="submit"
          >
            <Text
              fontWeight="regular"
              fontSize={{ base: "xl", md: "xl", lg: "2xl" }}
              m={"2%"}
              textAlign={"center"}
            >
              Register
            </Text>
          </Button>
        </form>
      )}
    </Flex>
  );
};
