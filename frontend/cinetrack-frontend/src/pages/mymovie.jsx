import { useState } from "react";
import { movies } from "../data/movies";
import Moviecard from "../components/moviecard";

function MyMovies() {
  const allStatus = movies.map((movie) => movie.status);
  const uniqueStatus = [...new Set(allStatus)];
  const statuses = ["All", ...uniqueStatus];

  const [selectedStatus, setSelectedStatus] = useState("All");
  const [movieSearch, setMovieSearch] = useState("");

  const statusFilteredMovies = movies.filter((movie) => {
    if (selectedStatus === "All") {
      return true;
    }

    return movie.status === selectedStatus;
  });

  const filteredMovies = statusFilteredMovies.filter((movie) =>
    movie.title.toLowerCase().includes(movieSearch.toLowerCase())
  );

  return (
    <div className="my-movie-page">
      <div className="heading">
        My Movies
      </div>

      <input
        type="text"
        placeholder="Search your movies..."
        value={movieSearch}
        onChange={(e) => setMovieSearch(e.target.value)}
        className="my-movie-search"
      />

      <div className="filter-buttons">
        {statuses.map((status) => (
          <button
            key={status}
            className={
              selectedStatus === status
                ? "status-button active"
                : "status-button"
            }
            onClick={() => setSelectedStatus(status)}
          >
            {status}
          </button>
        ))}
      </div>

      <div className="movie-result">
        {filteredMovies.map((movie) => (
          <Moviecard
            key={movie.id}
            movie={movie}
          />
        ))}
      </div>
    </div>
  );
}

export default MyMovies;