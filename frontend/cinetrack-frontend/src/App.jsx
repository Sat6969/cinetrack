
import Navbar from "./components/navbar.jsx";
import Home from  "./pages/Home";
import Discover from "./pages/Discover.jsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
function App(){

  return(
    <BrowserRouter>

      <Navbar />
      <Routes>
        <Route path="/" element={<Home />}/>
        <Route  path="/discover" element={<Discover />}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App