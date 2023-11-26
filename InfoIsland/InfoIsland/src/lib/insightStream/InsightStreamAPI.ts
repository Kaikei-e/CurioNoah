export enum Method {
  GET = "GET",
  POST = "POST",
  PUT = "PUT",
  DELETE = "DELETE",
}

const action = (inputMethod: Method) => {
  switch (inputMethod) {
    case Method.GET:
      return "GET";
    case Method.POST:
      return "POST";
    case Method.PUT:
      return "PUT";
    case Method.DELETE:
      return "DELETE";
    default:
      throw new Error("Unsupported method: " + inputMethod);
  }
};

export const insightStreamAPI = async (
  method: Method,
  url: string,
  body?: { [key: string]: string }
) => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;
  const headers = {
    "Content-Type": "application/json",
    Accept: "application/json",
    Origin: origin,
  };

  let options: RequestInit = {
    method: method,
    headers: headers,
  };

  if (method === Method.GET && body) {
    const queryParam = new URLSearchParams(body).toString();
    url = `${apiURL}${url}?${queryParam}`;
  } else if (body) {
    options.body = JSON.stringify(body);
    url = apiURL + url;
  }
  
  console.log("url", url);

  const response = await fetch(url, options);

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
};
