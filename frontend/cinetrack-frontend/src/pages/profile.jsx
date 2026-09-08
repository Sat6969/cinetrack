import Moviecard from "../components/moviecard";
import { movies } from "../data/movies";

function Profile() {
  const watchedCount = movies.filter(
    (movie) => movie.status === "Watched",
  ).length;

  const watchingCount = movies.filter(
    (movie) => movie.status === "Watching",
  ).length;

  const wantToWatchCount = movies.filter(
    (movie) => movie.status === "Want to Watch",
  ).length;

  const watchedMovies = movies.filter((movie) => movie.status === "Watched");

  const watchedGenres = watchedMovies.flatMap((movie) => movie.genres);

  const genreCount = {};

  watchedGenres.forEach((genre) => {
    if (genreCount[genre]) {
      genreCount[genre]++;
    } else {
      genreCount[genre] = 1;
    }
  });

  const genreArray = Object.entries(genreCount);

  const topGenres = genreArray.sort((a, b) => b[1] - a[1]).slice(0, 3);

  const ratedMovies = movies.filter((movie) => movie.userRating != null);

  const topRatedMovies = [...ratedMovies]
    .sort((a, b) => b.userRating - a.userRating)
    .slice(0, 3);

  const totalRating = ratedMovies.reduce((sum, movie) => {
    return sum + movie.userRating;
  }, 0);

  let averageRating;

  if (ratedMovies.length > 0) {
    averageRating = (totalRating / ratedMovies.length).toFixed(1);
  } else {
    averageRating = 0;
  }
  const reviewedMovies = movies.filter((movie) => movie.review);
  return (
    <div className="profile-page">
      <div className="profile-header">
        <div className="name">Name</div>
        <div className="joined-date">Joined Date</div>
      </div>

      <div className="profile-stats">
        <div className="stat-card">
          <p>Watched</p>
          <h2>{watchedCount}</h2>
        </div>

        <div className="stat-card">
          <p>Watching</p>
          <h2>{watchingCount}</h2>
        </div>

        <div className="stat-card">
          <p>Want to Watch</p>
          <h2>{wantToWatchCount}</h2>
        </div>

        <div className="stat-card">
          <p>Average Rating</p>
          <h2>{averageRating}</h2>
        </div>
      </div>

      <div className="favorite-genres">
        <div className="heading">Favorite Genres</div>

        <div className="genre-list">
          {topGenres.map(([genre, count]) => (
            <div className="favorite-genre" key={genre}>
              <span>{genre}</span>
              <span>{count}</span>
            </div>
          ))}
        </div>
      </div>

      <div className="top-rated-section">
        <div className="heading">Top Rated By You</div>

        <div className="movie-container">
          {topRatedMovies.map((movie) => (
            <Moviecard movie={movie} key={movie.id} />
          ))}
        </div>
      </div>

      <div className="recent-activity">
        <div className="heading">Recent Reviews</div>

        <div className="review-list">
          {reviewedMovies.map((movie) => (
            <div className="review-item" key={movie.id}>
              <div className="review-title">{movie.title}</div>
              <div className="review-rating">
                Your Rating: {movie.userRating}/10
              </div>
              <p className="review-text">{movie.review}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Profile;
