import { createBrowserRouter } from "react-router-dom";
import NoPage from "./pages/NoPage";
import PWRPage from "./pages/PWRPage";
import SignUp from "./pages/SignUp";
import LogIn from "./pages/LogIn";
import PWRchange from "./pages/PWRC";
import VEmail from "./pages/VEmail";
import Root from "./components/Root";

const router = createBrowserRouter(
    [
        {
            path:"",
            element:<Root></Root>
        },
        {
            path:"/signup/",
            element:<SignUp/>,
        },
        {
            path:"/recover/",
            element:<PWRPage/>,
        },
        {
            path:"/login/",
            element:<LogIn/>,
        },
        {
            path:"/recovery/:hash/:email/",
            element:<PWRchange/>

        },
        {
            path:"/verify-email/:hash/:email/",
            element:<VEmail/>

        },
        {
            path:"*",
            element:<NoPage/>
        },
        
    ]
);

export default router;