const RandomWords = async (length: number): Promise<string[]> => {
  const url = "https://random-word-api.vercel.app/api?words=" + length;

  const request = await fetch(url);

  const words = (await request.json()) as Array<string>;

  return words;
};

export { RandomWords };
