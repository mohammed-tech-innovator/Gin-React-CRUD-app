import React,{useState,useEffect} from "react";
import { Outlet } from "react-router-dom";


import Navbar from "./Navbar";
import FlashNote from "./FlashNote";
import Sider from "./Sider";








export default () => {
    const photo = "https://wonderfulengineering.com/wp-content/uploads/2014/10/image-wallpaper-15.jpg";

    const [fastNoteForm,setFastNF] = useState({
        noteBody: "",
        public: true,
        media:"",
        voice:""
    })

    return (
        <>
            <Navbar/>
            <div className="container">
                <div className="row mt-2">
                    <div className="col-xl-3 col-lg-3 d-none d-lg-block">

                        <Sider/>

                    </div>
                    <div className="col-xl-9 col-lg-9 col-md-12 col-sm-12 col-12">
                        <FlashNote fastNoteForm = {fastNoteForm} />
                    </div>
                </div>
            </div>
            </>

    )
}