export const collectAllNewFeeds = async () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  let options: RequestInit = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Origin: origin,
    },
  };

  const url = `${apiURL}/feeds/fetchAllfeeds`;

  const response = await fetch(url, options);

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
};
