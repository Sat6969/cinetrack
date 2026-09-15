const API_URL = import.meta.env.VITE_API_URL;

async function getmovies() {
  const response = await fetch(API_URL + "/movies");

  if (!response.ok) {
    throw new Error("could not fetch movies");
  }

  const rawmovies = await response.json();

  const movies = rawmovies.map((movie) => ({
    id: movie.ID,
    title: movie.title,
    year: movie.release_year,
    rating: movie.rating,

    genres: movie.genres.map((genre) => {
      return genre.name;
    }),

    poster: "/images/fallback.jpg",

    status: null,
    userRating: null,
    review: "",
  }));

  return movies;
}

async function Getmoviebyid(id) {
  const response = await fetch(
    API_URL + "/movies/" + id
  );

  if (!response.ok) {
    throw new Error("could not fetch movie");
  }

  const rawmovie = await response.json();

  const movie = {
    id: rawmovie.ID,
    title: rawmovie.title,
    year: rawmovie.release_year,
    rating: rawmovie.rating,

    genres: rawmovie.genres.map((genre) => {
      return genre.name;
    }),

    poster: "/images/fallback.jpg",

    status: null,
    userRating: null,
    review: "",
  };

  return movie;
}

export default getmovies;
export { Getmoviebyid };