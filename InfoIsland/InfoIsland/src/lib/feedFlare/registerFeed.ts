import { registeredFeed } from "../models/apiExchange/registeredFeed";

export async function RegisterFeed(feedURL: string): Promise<registeredFeed> {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  if (!checkUrlIsValid(feedURL)) {
    throw new Error("Invalid URL");
  }

  const response = await fetch(apiURL + "/register-feed/store/single", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
      Origin: origin,
    },
    body: JSON.stringify({
      url: feedURL,
      user_id: "1",
    }),
  });
  const data = await response.json().catch((err) => {
    console.log(err);
    return false;
  });

  // TODO I have to change the API model tp return a boolean
  return await JSON.parse(JSON.stringify(data));
}

function checkUrlIsValid(url: string): boolean {
  const parsedURL = new URL(url);
  return parsedURL.protocol === "https:";
}
