import Moviecard from "./components/moviecard";
import {movies} from "./data/movies.js";
import Navbar from "./components/navbar.jsx";
function App(){
  return(
    <>
    <Navbar />

    <div className="movie-container">
      {movies.map((movie)=>(
      <Moviecard key={movie.id} movie={movie}/>
    ))}
    </div>
    </>
  )
}

export default App