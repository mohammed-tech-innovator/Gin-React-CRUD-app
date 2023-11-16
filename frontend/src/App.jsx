import ReactDOM from "react-dom/client";
import { useState } from "react";
import { RouterProvider } from "react-router-dom";
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import router from "./router";

import 'bootstrap/dist/css/bootstrap.css'




function App() {

  return (
    <>
      <RouterProvider router = {router}></RouterProvider>
    </>
  )
}

export default App
