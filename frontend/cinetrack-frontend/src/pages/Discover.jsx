import { useEffect, useState } from "react";
import getmovies from "../services/movieService";
import Moviecard from "../components/moviecard";

function Discover() {
  const [movies, setmovies] = useState([]);
  const [loading, setloading] = useState(true);
  const [error, seterror] = useState("");
  const [selectedgenres, setgrenres] = useState({});
  const [search, setsearch] = useState("");

  useEffect(() => {
    async function loadmovies() {
      try {
        const data = await getmovies();

        const genreobj = {};

        setmovies(data);

        const rawgenres = data.flatMap((movie) => {
          return movie.genres;
        });

        const uniquegenres = [...new Set(rawgenres)];

        uniquegenres.forEach((gen) => {
          genreobj[gen] = false;
        });

        setgrenres(genreobj);
      } catch {
        seterror("something went wrong");
      } finally {
        setloading(false);
      }
    }

    loadmovies();
  }, []);

  if (loading === true) {
    return <div>loading...</div>;
  }

  if (error !== "") {
    return <div>something went wrong</div>;
  }

  const rawgenres = movies.flatMap((movie) => {
    return movie.genres;
  });

  const uniquegenres = [...new Set(rawgenres)];
  const genres = [...uniquegenres];

  const filtermovie1 = movies.filter((movie) => {
    return movie.title.toLowerCase().includes(search.toLowerCase());
  });

  const activegenres = Object.keys(selectedgenres).filter((genre) => {
    return selectedgenres[genre] === true;
  });

  const filtermovie2 = filtermovie1.filter((movie) => {
    if (activegenres.length === 0) {
      return true;
    }

    return activegenres.every((genre) => {
      return movie.genres.includes(genre);
    });
  });

  return (
    <div className="discover">
      <div className="heading">heading</div>

      <input
        type="text"
        placeholder="search here"
        className="search-input"
        onChange={(e) => {
          setsearch(e.target.value);
        }}
      />

      <div className="genre-buttons">
        {genres.map((genre) => (
          <button
            className={
              selectedgenres[genre] === true
                ? "genre-button-area active"
                : "genre-button-area"
            }
            key={genre}
            onClick={() => {
              const updatedgenres = {
                ...selectedgenres,
              };

              updatedgenres[genre] = !updatedgenres[genre];

              setgrenres(updatedgenres);
            }}
          >
            {genre}
          </button>
        ))}
      </div>
      <div className="movie-results-area">
        {filtermovie2.map((movie) => (
          <Moviecard movie={movie} key={movie.id} />
        ))}
      </div>
    </div>
  );
}

export default Discover;
