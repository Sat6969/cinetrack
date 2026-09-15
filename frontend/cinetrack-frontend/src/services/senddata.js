import { getToken } from "./authService";

const API_URL = import.meta.env.VITE_API_URL;

async function senddata(userreview, id) {
  const token = getToken();

  const response = await fetch(
    API_URL + "/movies/" + id + "/tracking",
    {
      method: "PUT",

      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + token,
      },

      body: JSON.stringify(userreview),
    }
  );

  const savedData = await response.json();

  if (!response.ok) {
    throw new Error(
      savedData.message || "could not save movie data"
    );
  }

  return savedData;
}

export { senddata };