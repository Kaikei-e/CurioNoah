import { Flex, Spinner } from "@chakra-ui/react";
import {FeedCard }from "./FeedCard";
import { Item } from "../../../../lib/models/feedModel";
import React from "react";
import useSWR from "swr";

async function fetcher(key: string, init?: RequestInit) {
  return fetch(key, init).then((res) => res.json() as Promise<Item | null>);
}

export const StreamLine = () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const { data: item, error } = useSWR(
    apiURL + "/fetch-feed/stored-all",
    fetcher
  );

  return (
    <Flex flexDirection={"column"} m={"2%"}>
      {typeof item === "undefined" ? (
        <Spinner />
      ) : item ? (
        <FeedCard feed={item} />
      ) : (
        <p>Not found</p>
      )}
      {error ? <p>{error}</p> : null}
    </Flex>
  );
};

