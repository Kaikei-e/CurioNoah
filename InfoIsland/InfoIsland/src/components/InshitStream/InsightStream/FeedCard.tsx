import { Item } from "../../../lib/models/feedModel";
import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Text,
  Flex,
  Avatar,
  Heading,
  IconButton,
  Box,
  Image,
  Button,
} from "@chakra-ui/react";
import { StarIcon, CopyIcon, CheckIcon } from "@chakra-ui/icons";

export const FeedCard = ({ feed }: { feed: Item }) => {
  return (
    <Card maxW="md">
      <CardHeader>
        <Flex></Flex>
      </CardHeader>
      <CardBody>
        <Text>
          With Chakra UI, I wanted to sync the speed of development with the
          speed of design. I wanted the developer to be just as excited as the
          designer to create a screen.
        </Text>
      </CardBody>
      <CardFooter
        justify="space-between"
        flexWrap="wrap"
        sx={{
          "& > button": {
            minW: "136px",
          },
        }}
      >
        <Button flex="1" variant="ghost" leftIcon={<CheckIcon />}>
          Like
        </Button>
        <Button flex="1" variant="ghost" leftIcon={<StarIcon />}>
          Comment
        </Button>
        <Button flex="1" variant="ghost" leftIcon={<CopyIcon />}>
          Share
        </Button>
      </CardFooter>
    </Card>
  );
};

