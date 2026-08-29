function Navbar() {
  return (
    <nav className="navbar">
      <div className="navbar-inner">
        <div className="logo">
          Cine<span>Track</span>
        </div>

        <div className="nav-links">
          <button className="nav-link active">Home</button>
          <button className="nav-link">Discover</button>
          <button className="nav-link">My Movies</button>
          <button className="nav-link">Profile</button>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;