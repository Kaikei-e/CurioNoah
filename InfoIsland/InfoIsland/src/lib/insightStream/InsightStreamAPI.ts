enum Method {
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
  inputAction: string,
  url: string,
  body?: Object
) => {
  const apiURL = import.meta.env.VITE_INSIGHT_STREAM;
  const origin = import.meta.env.VITE_ORIGIN;

  const wantedAction = action(inputAction as Method);
  if (wantedAction === Method.GET) {
    const queryParam = new URLSearchParams(
      body as { [key: string]: string }
    ).toString();

    const response = await fetch(`${apiURL}${url}?${queryParam}`, {
      method: wantedAction,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Origin: origin,
      },
    });
    return await response.json().catch((err) => {
      console.log(err);
      return false;
    });
  } else if (wantedAction === Method.POST) {
    const response = await fetch(apiURL + url, {
      method: wantedAction,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
        Origin: origin,
      },
      body: JSON.stringify(body),
    });
    return await response.json().catch((err) => {
      console.log(err);
      return false;
    });
  }
};
