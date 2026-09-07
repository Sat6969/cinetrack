import Moviecard from "../components/moviecard";
import { movies } from "../data/movies";
import { useState } from "react";

function Discover() {
  const allGenres = movies.flatMap((movie) => movie.genres);
  const uniqueGenres = [...new Set(allGenres)];
  const genres = ["All", ...uniqueGenres];
  const [selectedGenre, setSelectedGenre] = useState("All");
  const [search, setSearch] = useState("");
  const filteredMovies = movies.filter((movie) => {
    const matchSearch = movie.title
      .toLowerCase()
      .includes(search.toLowerCase());

    const matchesGenre =
      selectedGenre === "All" || movie.genres.includes(selectedGenre);

    return matchSearch && matchesGenre;
  });
  return (
    <div className="discover">
      <p className="small-label">DISCOVER</p>
      <h1 className="heading">Find your next film.</h1>
      <input
        type="text"
        className="search-input"
        placeholder="Search movies..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
      <div className="genre-button-area">
        {genres.map((genre) => (
          <button
            className={
              selectedGenre === genre ? "genre-button active" : "genre-button"
            }
            key={genre}
            onClick={() => setSelectedGenre(genre)}
          >
            {genre}
          </button>
        ))}
      </div>
      <div className="movie-results-area">
        {filteredMovies.map((movie) => (
          <Moviecard movie={movie} key={movie.id}></Moviecard>
        ))}
      </div>
    </div>
  );
}

export default Discover;
