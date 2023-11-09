import React from "react";

export default function Navbar(params) {
    const logo_style = {
        hight:"30px",
        width:"30px",
        borderRadius:"round 10px",
    }
    return(
        <div>
            <nav className="navbar navbar-light bg-light">
                <div className="container-fluid">
                    
                    <a className="navbar-brand" href="/"><img style ={logo_style} src = {params.logo}/><span> </span>{params.com_name}</a>
                    <form className="d-flex" method="POST" action={window.homeRoute}>
                        <input className="form-control me-2" type="search" placeholder="Search" aria-label="Search"/>
                        <button className="btn btn-outline-success" type="submit">Search</button>
                    </form>
                </div>
            </nav>
        </div>
    )
}