import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

import Moviecard from "../components/Moviecard";
import {
  getCurrentUser,
  getMyMovies,
} from "../services/trackingService";

import { getToken } from "../services/authService";

function Profile() {
  const [user, setUser] = useState(null);
  const [movies, setMovies] = useState([]);

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadProfile() {
      const token = getToken();

      if (!token) {
        setError("please login first");
        setLoading(false);
        return;
      }

      try {
        const userData = await getCurrentUser();
        const movieData = await getMyMovies();

        setUser(userData);
        setMovies(movieData);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }

    loadProfile();
  }, []);

  if (loading) {
    return <div>loading...</div>;
  }

  if (error === "please login first") {
    return (
      <div className="profile-page">
        <h1>Profile</h1>

        <p>Please login to view your profile.</p>

        <Link to="/login">
          Login
        </Link>
      </div>
    );
  }

  if (error !== "") {
    return <div>{error}</div>;
  }

  if (!user) {
    return null;
  }

  // -----------------------------
  // STATUS COUNTS
  // -----------------------------

  const watchedMovies = movies.filter((movie) => {
    return movie.status === "Watched";
  });

  const watchingMovies = movies.filter((movie) => {
    return movie.status === "Watching";
  });

  const wantToWatchMovies = movies.filter((movie) => {
    return movie.status === "Want to Watch";
  });

  const droppedMovies = movies.filter((movie) => {
    return movie.status === "Dropped";
  });

  // -----------------------------
  // FAVORITE GENRES
  // -----------------------------

  const watchedGenres = watchedMovies.flatMap((movie) => {
    return movie.genres;
  });

  const genreCount = {};

  watchedGenres.forEach((genre) => {
    if (genreCount[genre]) {
      genreCount[genre] = genreCount[genre] + 1;
    } else {
      genreCount[genre] = 1;
    }
  });

  const favoriteGenres = Object.entries(genreCount)
    .sort((a, b) => {
      return b[1] - a[1];
    })
    .slice(0, 3);

  // -----------------------------
  // USER RATINGS
  // -----------------------------

  const ratedMovies = movies.filter((movie) => {
    return (
      movie.userRating >= 1 &&
      movie.userRating <= 10
    );
  });

  let averageRating = 0;

  if (ratedMovies.length > 0) {
    const totalRating = ratedMovies.reduce(
      (total, movie) => {
        return total + movie.userRating;
      },
      0
    );

    averageRating =
      totalRating / ratedMovies.length;
  }

  // -----------------------------
  // TOP RATED MOVIES
  // -----------------------------

  const topRatedMovies = [...ratedMovies]
    .sort((a, b) => {
      return b.userRating - a.userRating;
    })
    .slice(0, 3);

  // -----------------------------
  // REVIEWS
  // -----------------------------

  const reviewedMovies = movies.filter((movie) => {
    return (
      movie.review &&
      movie.review.trim() !== ""
    );
  });

  // -----------------------------
  // JOIN DATE
  // -----------------------------

  const joinedDate = new Date(
    user.joined_at
  ).toLocaleDateString();

  return (
    <div className="profile-page">

      {/* USER INFO */}

      <div className="profile-header">
        <h1>{user.name}</h1>

        <p>{user.email}</p>

        <p>
          Joined: {joinedDate}
        </p>
      </div>

      {/* STATS */}

      <div className="profile-stats">
        <div className="stat-box">
          <h3>Watched</h3>
          <p>{watchedMovies.length}</p>
        </div>

        <div className="stat-box">
          <h3>Watching</h3>
          <p>{watchingMovies.length}</p>
        </div>

        <div className="stat-box">
          <h3>Want to Watch</h3>
          <p>{wantToWatchMovies.length}</p>
        </div>

        <div className="stat-box">
          <h3>Dropped</h3>
          <p>{droppedMovies.length}</p>
        </div>

        <div className="stat-box">
          <h3>Average Rating</h3>

          <p>
            {ratedMovies.length > 0
              ? averageRating.toFixed(1)
              : "0"}
          </p>
        </div>
      </div>

      {/* FAVORITE GENRES */}

      <div className="favorite-genres">
        <h2>Favorite Genres</h2>

        {favoriteGenres.length === 0 ? (
          <p>
            Watch some movies to discover your
            favorite genres.
          </p>
        ) : (
          <div className="genre-list">
            {favoriteGenres.map(
              ([genre, count]) => (
                <div key={genre}>
                  {genre} ({count})
                </div>
              )
            )}
          </div>
        )}
      </div>

      {/* TOP RATED */}

      <div className="top-rated-section">
        <h2>Your Top Rated Movies</h2>

        {topRatedMovies.length === 0 ? (
          <p>
            You have not rated any movies yet.
          </p>
        ) : (
          <div className="movie-grid">
            {topRatedMovies.map((movie) => (
              <div key={movie.id}>
                <Moviecard movie={movie} />

                <p>
                  Your rating:{" "}
                  {movie.userRating}/10
                </p>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* REVIEWS */}

      <div className="reviews-section">
        <h2>Your Reviews</h2>

        {reviewedMovies.length === 0 ? (
          <p>
            You have not written any reviews yet.
          </p>
        ) : (
          reviewedMovies.map((movie) => (
            <div
              key={movie.id}
              className="profile-review"
            >
              <h3>{movie.title}</h3>

              <p>
                Rating: {movie.userRating}/10
              </p>

              <p>{movie.review}</p>
            </div>
          ))
        )}
      </div>

    </div>
  );
}

export default Profile;