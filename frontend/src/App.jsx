import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import Navbar from './components/navbar'
import Home from './pages/Home'
import 'bootstrap/dist/css/bootstrap.css'
import testhome from './pages/testhome'


function App() {
  const [count, setCount] = useState(0)

  return (
    <Home />
    
  )
}

export default App
