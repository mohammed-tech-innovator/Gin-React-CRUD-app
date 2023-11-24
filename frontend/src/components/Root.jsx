import React from "react";
import { Outlet } from "react-router-dom";


import Navbar from "./Navbar";
import FlashNote from "./FlashNote";
import Sider from "./Sider";






export default () => {
    const photo = "https://wonderfulengineering.com/wp-content/uploads/2014/10/image-wallpaper-15.jpg";
    return (
        <>
            <Navbar/>
            <div className="container">
                <div className="row mt-2">
                    <div className="col-xl-2 col-lg-2 d-none d-lg-block">

                        <Sider/>

                    </div>
                    <div className="col-xl-8 col-lg-8 col-md-12 col-sm-12 col-12">
                        <FlashNote/>
                    </div>
                    <div className="col-xl-2 col-lg-2 d-none d-lg-block">
                        <div className="card">
                            
                        </div>
                    </div>
                </div>
            </div>
            </>

    )
}