import { type SubmitHandler, useForm } from "react-hook-form";
import {
  insightStreamAPI,
  Method,
} from "../../../lib/insightStream/InsightStreamAPI";
import React from "react";
import type { ByTitleOrDescription } from "../../../lib/models/searchFeeds/byTitleOrDescription";
import type { TitleOrDescription } from "../../../lib/models/searchFeeds/baseFeed";
import { Flex } from "@chakra-ui/react";
import { CircularProgress } from "@chakra-ui/progress";

type Props = {
  title: string;
  description: string;
};

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
  } = useForm<Input>({
    defaultValues: {
      title: "",
      description: "",
    },
  });

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
      if (res.feeds === null) {
        console.log("No feeds found");
        res.feeds = [];
      }

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
            size="48px"
            borderWidth="10px" // thickness→borderWidth
          />
        ) : (
          <Flex flexDirection={"column"}>
            {items.map((item) => (
              <div key={item.feed_url}>
                <a
                  href={item.feed_url}
                  style={feedsTitleStyle}
                  target={"_blank"}
                  rel="noopener noreferrer" // isExternal→target/rel
                >
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

const feedsTitleStyle: React.CSSProperties = {
  fontSize: "1.4rem",
  fontWeight: "bold",
  color: "#000",
  fontFamily: "Jost",
  textAlign: "left",
  margin: "0.5rem",
};

const feedsDescriptionStyle: React.CSSProperties = {
  fontSize: "1rem",
  fontWeight: "regular",
  color: "#000",
  fontFamily: "Jost",
  textAlign: "left",
  margin: "0.5rem",
};
