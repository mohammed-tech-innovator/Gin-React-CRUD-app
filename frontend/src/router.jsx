import { createBrowserRouter } from "react-router-dom";
import NoPage from "./pages/NoPage";
import PWRPage from "./pages/PWRPage";
import SignUp from "./pages/SignUp";
import LogIn from "./pages/LogIn";

const router = createBrowserRouter(
    [
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
            path:"*",
            element:<NoPage/>
        },
        
    ]
);

export default router;