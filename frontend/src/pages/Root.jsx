import React from "react";

import { Outlet } from "react-router-dom";
import Footer from "../components/Footer";
import Navbar from "../components/navbar";

const Root = () => {

    return (
        <>
            <Navbar com_name = {"Buyut"} logo ={"https://th.bing.com/th/id/OIP.l590m8rOs0GW_XV2hHgRYAHaHa?pid=ImgDet&rs=1"}/>
                <Outlet/>
            <Footer/>
        </>
        
    )

}

export default Root;