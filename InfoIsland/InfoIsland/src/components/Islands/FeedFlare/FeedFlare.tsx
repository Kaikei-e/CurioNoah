import { SubmitHandler, useForm } from "react-hook-form";
import { insightStreamAPI } from "../../../lib/insightStream/InsightStreamAPI";
import React from "react";
import { ByTitleOrDescription } from "../../../lib/models/searchFeeds/byTitleOrDescription";
import { TitleOrDescription } from "../../../lib/models/searchFeeds/baseFeed";

type Props = {};

type Input = {
  searchKeyword: string;
};

export const FeedFlare = (props: Props) => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Input>();

  const [items, setItems] = React.useState<ByTitleOrDescription[]>([]);
  const [searchKeyword, setSearchKeyword] = React.useState<TitleOrDescription>({
    title: "",
    description: "",
  });

  const onSubmit: SubmitHandler<Input> = (data: Input) => {
    const res = insightStreamAPI("GET", "/search/baseFeeds", {
      title: data.searchKeyword,
      description: data.searchKeyword,
    })
      .then((res) => {
        setItems(res.feeds);
      })
      .catch((err) => {
        console.log(err);
      });

    console.log(res);
  };
  return (
    <div>
      <h3>FeedFlare</h3>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input defaultValue="Go" {...register("searchKeyword")} />
        <input type="submit" />
      </form>

      {items.map((item: ByTitleOrDescription, index: number) => (
        <div key={index}>
          <h1>{item.title}</h1>
          <p>{item.description}</p>
          <p>{item.feed_url}</p>
        </div>
      ))}
    </div>
  );
};
