

const API_URL = import.meta.env.VITE_API_URL;


async function getmovies() {
  let response = await fetch(API_URL + "/movies");

  if (!response.ok) {
    console.log("error");
    return;
  }
  const rawmovies = await response.json();

  const movies = rawmovies.map((movie) => ({
    id: movie.id,
    title: movie.title,
    year: movie.year,
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


async function  Getmoviebyid(id) {
  
  const data=await fetch(API_URL+"/movies/" +id)

   if(!data.ok){
    console.log("error cant fetch id");
   }

  const rawmovie=await data.json()

  const movie={
    id: rawmovie.id,
    title: rawmovie.title,
    year: rawmovie.year,
    rating: rawmovie.rating,
    genres: rawmovie.genres.map((genre)=>{
      return genre.name
  }),
    poster: "/images/fallback.jpg",
    status: null,
    userRating: null,
    review: ""
  }

  return movie;

}


export default getmovies 

export {Getmoviebyid}