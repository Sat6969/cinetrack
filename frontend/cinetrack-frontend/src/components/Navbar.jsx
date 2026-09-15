import { useEffect, useState } from "react";
import {
  Link,
  useLocation,
  useNavigate,
} from "react-router-dom";

import {
  getToken,
  logoutUser,
} from "../services/authService";

function Navbar() {
  const [loggedIn, setLoggedIn] = useState(false);

  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const token = getToken();

    if (token) {
      setLoggedIn(true);
    } else {
      setLoggedIn(false);
    }
  }, [location]);

  function handleLogout() {
    logoutUser();

    setLoggedIn(false);

    navigate("/");
  }

  return (
    <nav className="navbar">
      <div className="navbar-logo">
        <Link to="/">
          CineTrack
        </Link>
      </div>

      <div className="navbar-links">
        <Link to="/">
          Home
        </Link>

        <Link to="/discover">
          Discover
        </Link>

        {loggedIn && (
          <>
            <Link to="/my-movies">
              My Movies
            </Link>

            <Link to="/profile">
              Profile
            </Link>
          </>
        )}

        {!loggedIn && (
          <>
            <Link to="/login">
              Login
            </Link>

            <Link to="/register">
              Register
            </Link>
          </>
        )}

        {loggedIn && (
          <button onClick={handleLogout}>
            Logout
          </button>
        )}
      </div>
    </nav>
  );
}

export default Navbar;