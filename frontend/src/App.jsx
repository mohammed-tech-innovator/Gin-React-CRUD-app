import ReactDOM from "react-dom/client";
import { useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import Navbar from './components/navbar'
import Home from './pages/Home'
import Estate from './pages/Estate'
import 'bootstrap/dist/css/bootstrap.css'
import testhome from './pages/testhome'
import NoPage from "./pages/NoPage";


function App() {
  const [count, setCount] = useState(0)

  return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />}/>
          <Route path="/estate/:id" element={<Estate />} />
          <Route path="*" element={<NoPage />} />
        </Routes>
      </BrowserRouter>
  )
}

export default App
