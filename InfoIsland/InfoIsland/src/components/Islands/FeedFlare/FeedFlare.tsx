import { SubmitHandler, useForm } from "react-hook-form";
import {
  insightStreamAPI,
  Method,
} from "../../../lib/insightStream/InsightStreamAPI";
import React, { useEffect } from "react";
import { ByTitleOrDescription } from "../../../lib/models/searchFeeds/byTitleOrDescription";
import { TitleOrDescription } from "../../../lib/models/searchFeeds/baseFeed";
import { Button, CircularProgress, Flex } from "@chakra-ui/react";

type Props = {};

type Input = {
  title: string;
  description: string;
};

export const FeedFlare = (props: Props) => {
  const {
    register,
    setValue,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Input>();

  const [isLoading, setIsLoading] = React.useState<boolean>(false);
  const [items, setItems] = React.useState<ByTitleOrDescription[]>([]);
  const [searchKeyword, setSearchKeyword] = React.useState<TitleOrDescription>({
    title: "",
    description: "",
  });

  const onSubmit: SubmitHandler<Input> = async (data: Input, event: any) => {
    event.preventDefault();
    setIsLoading(true);
    try {
      const res = await insightStreamAPI(Method.GET, "/search/baseFeeds", {
        title: data.title,
        description: data.description,
      });
      if (Array.isArray(res.feeds)) {
        setItems([]);
        setItems(res.feeds);
      } else {
        console.error("Unexpected response structure", res);
      }
    } catch (err) {
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };
  return (
    <Flex h={"100%"} w={"100%"} flexDirection={"column"}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          defaultValue="Go"
          {...register("title")}
          {...register("description")}
        />
        <input type="submit" />
      </form>

      <Flex h={"100%"} w={"100%"}>
        {isLoading ? (
          <CircularProgress
            isIndeterminate
            color="green.300"
            trackColor="green.100"
            size="48px"
            thickness="10px"
          />
        ) : (
          <Flex flexDirection={"column"}>
            {items.map((item, index) => (
                <div key={index}>
                  <a href={item.feed_url} style={feedsTitleStyle} target={"_blank"}>
                    {item.title}
                  </a>
                  <p style={feedsDescriptionStyle}>{item.description}</p>
                </div>
            ))}
          </Flex>
        )}
      </Flex>
    </Flex>
  );
};


const feedsTitleStyle: any = {
  fontSize: "1.4rem",
  fontWeight: "bold",
  color: "#000",
  fontFamily: "Jost",
  textAlign: "left",
  margin: "0.5rem",
};

const feedsDescriptionStyle: any = {
  fontSize: "1rem",
  fontWeight: "regular",
  color: "#000",
  fontFamily: "Jost",
  textAlign: "left",
  margin: "0.5rem",
};