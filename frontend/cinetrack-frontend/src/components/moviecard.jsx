

function Moviecard({movie}){

    return(
        <div className="card">
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
            
        </div>
    )

}

export default Moviecard