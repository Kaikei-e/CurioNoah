export const collectAllNewFeeds = async () => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;
  const headers = {
    "Content-Type": "application/json",
    Accept: "application/json",
    Origin: origin,
  };
  
  let options: RequestInit = {
    method: "POST",
    headers: headers,
  };
  
  const url = `${apiURL}/feeds/fetchAllfeeds`;
  
  const response = await fetch(url, options);
  
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  
  return response.json();
}