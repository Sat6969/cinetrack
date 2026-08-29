import { NavLink } from "react-router-dom";

function Navbar(){
    return(
        <nav>
        <div className="logo">CineTrack</div>

        <div className="nav-links">
            <NavLink to={"/"} className="nav-link">Home</NavLink>
            <NavLink to={"/discover"} className="nav-link">Discover</NavLink>
            <NavLink to={"/my-movies"} className="nav-link">My movies</NavLink>
            <NavLink to={"/profile"} className="nav-link">Profile</NavLink>
        </div>
        </nav>
    )
}
export default Navbar;