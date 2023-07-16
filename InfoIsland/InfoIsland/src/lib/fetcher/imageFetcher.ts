export async function ImageFetcher(imgURL: string): Promise<Blob> {
  const imgBody = await fetch(imgURL, {
    method: "GET",
    headers: {
      "Content-Type": "image/jpeg",
    },
  });

  return await imgBody.blob();
}
