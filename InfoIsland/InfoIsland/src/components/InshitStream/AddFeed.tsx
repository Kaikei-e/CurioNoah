import {Button, Flex, Text, Input} from "@chakra-ui/react";
import {CheckIcon} from "@chakra-ui/icons";
import RegisterFeed from "../../lib/feedFlare/registerFeed";
import {useState} from "react";
import registeredFeed from "../../lib/models/apiExchange/registeredFeed";

const AddFeed = () => {
    const [feedUrl, setFeedUrl] = useState("");

    return (
        <Flex flexDirection={"column"} w={"100%"} h={"100%"}
              fontFamily={"Jost"}>
            <Text fontSize={{base: "lg", md: "xl", lg: "xl"}}
                  mb={"2%"}
            >
                You can register a new feed by entering the URL of the feed
                in this input field below.
            </Text>
            <form method={"POST"}>
                <Flex flexDirection={"column"}
                      w={"100%"} h={"100%"} p={"2%"}
                >
                    <Input type="text"
                           placeholder="Enter Feed URL"
                           _placeholder={{opacity: 0.65, color: 'black'}}
                           size={"lg"} w={"100%"} h={"100%"}
                           mb={"2%"}
                           value={feedUrl}
                           onChange={(e) => setFeedUrl(e.target.value)}
                    />
                    <Button size={{base: "lg", sm: "sm", md: "md", lg: "lg"}}
                            variant='outline'
                            textColor={'#FFF'}
                            _hover={{color: '#000', bg: '#FFF'}}
                            rightIcon={<CheckIcon
                                boxSize={{base: "xl", md: "2xl", lg: "20px"}}
                            />}
                            bgColor={"#84a3ad"}

                            onClick={async (e) => {
                                e.preventDefault();
                                const response: registeredFeed = await RegisterFeed(feedUrl);

                                if (response.message === "success") {
                                    alert("Feed registered successfully!");
                                } else {
                                    alert("Feed registration failed!");
                                }
                            }}
                    >
                        <Text fontWeight="regular"
                              fontSize={{base: "xl", md: "xl", lg: "2xl"}}
                              m={"2%"}
                              textAlign={"center"}>
                            Resister
                        </Text>
                    </Button>
                </Flex>

            </form>
        </Flex>
    )
}

export default AddFeed;