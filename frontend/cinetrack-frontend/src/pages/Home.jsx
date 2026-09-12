import { useState } from "react";
import { useEffect } from "react";
import getmovies from "../services/movieService";
import Moviecard from "../components/moviecard";
import { NavLink } from "react-router-dom";

function Home() {
  const [movies, setMovies] = useState([]);
  const [loading, setloading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadmovie() {
      try {
        const data = await getmovies();
        setMovies(data);
      } catch {
        setError("something went wrong");
      } finally {
        setloading(false);
      }
    }
    loadmovie();
  }, []);

  if (loading === true) {
    return <div>loading...</div>;
  }
  if (error !== "") {
    return <div>something went wrong</div>;
  }

  const featuredmovies = movies.slice(0, 3);

  return (
    <div className="home">
      <div className="hero-section">
        <div className="small-label"> CINETRACK ARCHIVE</div>

        <div className="big-heading">Every film leaves a trace</div>

        <div className="short-description"> Discover films, track what you've watched,and build your personal movie archive.</div>

        <NavLink to="/discover" className="discover-button">Discover</NavLink>
      </div>
      <div className="featured-movies">
        <div className="featured-header">
          <div className="featured-label">CURATED SELECTION</div>
          <div className="featured-heading">Featured Films</div>
          <div className="featured-description">
            A few standout films from the archive.
          </div>
        </div>
        <div className="movie-container">
          {featuredmovies.map((m) => (
            <Moviecard movie={m} key={m.id}></Moviecard>
          ))}
        </div>
      </div>
    </div>
  );
}

export default Home;
