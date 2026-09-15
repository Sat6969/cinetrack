import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import Moviecard from "../components/Moviecard";
import { getMyMovies } from "../services/trackingService";
import { getToken } from "../services/authService";

function MyMovies() {
  const [movies, setMovies] = useState([]);
  const [selectedstate, setSelectedstate] = useState("All");

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadMyMovies() {
      const token = getToken();

      if (!token) {
        setError("please login first");
        setLoading(false);
        return;
      }

      try {
        const data = await getMyMovies();

        setMovies(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }

    loadMyMovies();
  }, []);

  if (loading) {
    return <div>loading...</div>;
  }

  if (error === "please login first") {
    return (
      <div className="my-movies">
        <h1>My Movies</h1>

        <p>Please login to see your movies.</p>

        <Link to="/login">
          Login
        </Link>
      </div>
    );
  }

  if (error !== "") {
    return (
      <div>
        {error}
      </div>
    );
  }

  const rawstatuses = movies.map((movie) => {
    return movie.status;
  });

  const validstatus = rawstatuses.filter((status) => {
    return status !== null && status !== "";
  });

  const uniquestatus = [
    "All",
    ...new Set(validstatus),
  ];

  const filteredmovies = movies.filter((movie) => {
    if (selectedstate === "All") {
      return true;
    }

    return movie.status === selectedstate;
  });

  return (
    <div className="my-movies">
      <h1>My Movies</h1>

      <div className="status-filters">
        {uniquestatus.map((status) => (
          <button
            key={status}
            onClick={() => {
              setSelectedstate(status);
            }}
            className={
              selectedstate === status
                ? "active"
                : ""
            }
          >
            {status}
          </button>
        ))}
      </div>

      {movies.length === 0 ? (
        <div>
          You have not tracked any movies yet.
        </div>
      ) : filteredmovies.length === 0 ? (
        <div>
          No movies found for this status.
        </div>
      ) : (
        <div className="movie-grid">
          {filteredmovies.map((movie) => (
            <Moviecard
              key={movie.id}
              movie={movie}
            />
          ))}
        </div>
      )}
    </div>
  );
}

export default MyMovies;