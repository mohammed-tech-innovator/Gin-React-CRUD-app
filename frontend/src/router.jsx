import { createBrowserRouter } from "react-router-dom";
import Root from "./pages/Root";
import Home from "./pages/Home";
import Estate from "./pages/Estate";
import NoPage from "./pages/NoPage";
import Owner from "./pages/Owner"
import Owners from "./pages/owners";

const router = createBrowserRouter(
    [
        {
            path:"/",
            element:<Root/>,
            children : [
            {
                path:"/",
                element:<Home/>,
                index: true,
                
            },
            {
                path:"/estate/:id",
                element:<Estate/>,
                
            },
            {
                path:"/owner/:id",
                element: <Owner/>,

            },
            {
                path:"/owners/",
                element: <Owners/>,

            },
        ],
        },
        {
            path:"*",
            element:<NoPage/>
        }
    ]
);

export default router;