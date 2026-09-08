import Moviecard from "../components/moviecard";
import { movies } from "../data/movies.js";
import { NavLink } from "react-router-dom";

function Home() {
  const featuredMovies = movies.slice(0, 3);
  return (
    <>
      <div className="hero-section">
        <div className="small-label">CINETRACK ARCHIVE</div>
        <div className="big-heading">Every film leaves a trace</div>
        <div className="short-description">
          Discover films, track what you've watched, and build your personal movie archive.
        </div>
        <NavLink to="/discover" className="discover-button">Discover</NavLink>
      </div>

      <div className="featured-movies">
        <div className="featured-header">
          <p className="featured-label">CURATED SELECTION</p>
          <h2 className="featured-heading">Featured Films</h2>
          <p className="featured-description">A few standout films from the archive.</p>
        </div>

        <div className="movie-container">
          {featuredMovies.map((movie) => (
            <Moviecard key={movie.id} movie={movie} />
          ))}
        </div>
      </div>
    </>
  );
}

export default Home;