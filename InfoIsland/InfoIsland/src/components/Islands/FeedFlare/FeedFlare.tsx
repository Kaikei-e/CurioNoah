import { SubmitHandler, useForm } from "react-hook-form";
import {
  insightStreamAPI,
  Method,
} from "../../../lib/insightStream/InsightStreamAPI";
import React, { useEffect } from "react";
import { ByTitleOrDescription } from "../../../lib/models/searchFeeds/byTitleOrDescription";
import { TitleOrDescription } from "../../../lib/models/searchFeeds/baseFeed";
import { Button, CircularProgress } from "@chakra-ui/react";

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

  useEffect(() => {
    setIsLoading(true);
  });

  const onSubmit: SubmitHandler<Input> = (data: Input, event: any) => {
    event.preventDefault();
    setIsLoading(true);
    const res = insightStreamAPI(Method.GET, "/search/baseFeeds", {
      title: data.title,
      description: data.description,
    })
      .then((res) => {
        setItems(res.feeds);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };
  return (
    <div>
      <h3>FeedFlare</h3>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input
          defaultValue="Go"
          {...register("title")}
          {...register("description")}
        />
        <input type="submit" />
      </form>

      {isLoading ? (
        <CircularProgress
          isIndeterminate
          color="green.300"
          trackColor="green.100"
          size="48px"
          thickness="10px"
        />
      ) : (
        items.map((item, index) => (
          <div key={index}>
            <h3>
              {item.title}
              <a href={item.feed_url}>{item.feed_url}</a>
            </h3>
            <p>{item.description}</p>
          </div>
        ))
      )}
    </div>
  );
};
