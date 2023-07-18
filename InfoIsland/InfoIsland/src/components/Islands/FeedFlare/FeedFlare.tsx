import { SubmitHandler, useForm } from "react-hook-form";
import {insightStreamAPI} from "../../../lib/insightStream/InsightStreamAPI";
import React from "react";
import {EachFeed} from "../../../lib/models/eachFeed";

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

  const [items, setItems] = React.useState<EachFeed[]>([]);

  const onSubmit: SubmitHandler<Input> = (data: Input) => {
    const res = insightStreamAPI("GET", "/search/feeds", data)
        .then((res) => {
            console.log(res);
            setItems(res.feeds)
        })
        .catch((err) => {
            console.log(err);
        });

    console.log(res);
  }
  return (
    <div>
      <h3>FeedFlare</h3>
      <form onSubmit={handleSubmit(onSubmit)}>
        <input defaultValue="Go" {...register("searchKeyword")} />
        <input type="submit" />
      </form>

      {items.map((item: EachFeed, index: number) => (
          <div key={index}>
            <h1>{item.title}</h1>
            <p>{item.description}</p>
            <p>{item.feed_url}</p>
            <p>{item.dt_updated}</p>
          </div>
      ))}
    </div>
  );
};
