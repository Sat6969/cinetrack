import { Link } from "react-router-dom"

function Moviecard({movie}){

    return(
        <div className="card">
            <Link to={`/movies/${movie.id}`} className="movie-card-link">
            <img src={movie.poster} alt="" className="movie-image" />
             <div className="moviename">{movie.title.toUpperCase()}</div>
            <div className="movie-details">
                <div className="movie-year">{movie.year}</div>
                 <div className="genres">
                    {movie.genres.join(" . ").toUpperCase()}
                </div>
                <div className="rating-container">
                     <div className="rating">{movie.rating}</div>
                </div>
            </div>
            <button className="add-to-watchlist">
                Add to Watchlist
            </button>
            </Link>
        </div>
    )

}

export default Moviecard