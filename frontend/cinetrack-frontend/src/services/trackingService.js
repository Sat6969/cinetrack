import { getToken } from "./authService";

const API_URL = import.meta.env.VITE_API_URL;

async function getTracking(id) {
  const token = getToken();

  const response = await fetch(
    API_URL + "/movies/" + id + "/tracking",
    {
      headers: {
        Authorization: "Bearer " + token,
      },
    }
  );

  if (!response.ok) {
    throw new Error("could not get tracking");
  }

  const data = await response.json();

  return data;
}

async function getMyMovies() {
  const token = getToken();

  const response = await fetch(
    API_URL + "/my-movies",
    {
      headers: {
        Authorization: "Bearer " + token,
      },
    }
  );

  if (!response.ok) {
    throw new Error("could not get my movies");
  }

  const rawmovies = await response.json();

  const movies = rawmovies.map((movie) => ({
    id: movie.id,
    title: movie.title,
    year: movie.release_year,
    rating: movie.rating,

    genres: movie.genres.map((genre) => {
      return genre.name;
    }),

    poster: "/images/fallback.jpg",

    status: movie.status,
    userRating: movie.userRating,
    review: movie.review,
  }));

  return movies;
}

async function getCurrentUser() {
  const token = getToken();

  const response = await fetch(
    API_URL + "/me",
    {
      headers: {
        Authorization: "Bearer " + token,
      },
    }
  );

  if (!response.ok) {
    throw new Error("could not get user");
  }

  return await response.json();
}

export {
  getTracking,
  getMyMovies,
  getCurrentUser,
};