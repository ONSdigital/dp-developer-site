export const fetchData = async (url) =>
  fetch(url)
      .then((response) => response.json())
      .catch((err) => {
        throw new Error(err);
      });
