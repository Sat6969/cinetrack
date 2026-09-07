import { NavLink } from "react-router-dom";

function Navbar(){
    return(
        <nav>
        <div className="logo">CineTrack</div>

        <div className="nav-links">
            <NavLink to={"/"} className={({isActive})=>{
              return  isActive ? "nav-link active" : "nav-link"
            }}>Home</NavLink>
            <NavLink to={"/discover"} className={({isActive})=>{
              return  isActive ? "nav-link active" : "nav-link"
            }}>Discover</NavLink>
            <NavLink to={"/my-movies"} className={({isActive})=>{
              return  isActive ? "nav-link active" : "nav-link"
            }}>My movies</NavLink>
            <NavLink to={"/profile"} className={({isActive})=>{
              return  isActive ? "nav-link active" : "nav-link"
            }}>Profile</NavLink>
        </div>
        </nav>
    )
}
export default Navbar;