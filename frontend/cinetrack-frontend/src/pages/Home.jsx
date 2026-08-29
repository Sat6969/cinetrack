import Moviecard from "../components/moviecard";
import { movies } from "../data/movies.js";

function Home() {
  return (
    <main className="home-page">

      <section className="home-hero">
        <p className="page-label">CineTrack Archive</p>

        <h1 className="home-title">
          Every film leaves a trace.
        </h1>

        <p className="home-description">
          Discover films, build your watchlist and keep track of
          everything you've watched.
        </p>
      </section>

      <section className="home-movies">
        <div className="section-heading">
          <div>
            <p className="page-label">Selected for you</p>
            <h2>Tonight's Selection</h2>
          </div>

          <span>VOL. 01</span>
        </div>

        <div className="movie-container">
          {movies.slice(0, 3).map((movie) => (
            <Moviecard
              key={movie.id}
              movie={movie}
            />
          ))}
        </div>
      </section>

    </main>
  );
}

export default Home;