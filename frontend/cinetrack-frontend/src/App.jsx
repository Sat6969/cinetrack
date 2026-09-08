
import Navbar from "./components/navbar.jsx";
import Home from  "./pages/Home";
import Discover from "./pages/Discover.jsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import MyMovies from "./pages/mymovie.jsx";
import Profile from "./pages/profile.jsx";
import ScrollToTop from "./components/scrolltotop.jsx";

function App(){

  return(
    <BrowserRouter>
      <ScrollToTop></ScrollToTop>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />}/>
        <Route  path="/discover" element={<Discover />}/>
        <Route path="/my-movies" element={<MyMovies />} />
        <Route path="/profile" element={<Profile></Profile>}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App