import React from "react";

export default function Carousal() {

    const carousal_img = {
        width:"100%",
        height: "30%",

    }
    const carousal_style = {
        maxHeight: '400px'
    }
    return(
        <div id="carouselExampleControls" className ="carousel slide" data-ride="carousel">
            <div className ="carousel-inner" style={carousal_style} >
                <div className ="carousel-item active">
                    <img className ="d-block w-100" style = {carousal_img} src="https://www.toolsgroup.com/wp-content/uploads/2019/05/wayfair8FAB1617-5056-A830-8C54D49938B6682C.jpg" alt="First slide"/>
                </div>
                <div className ="carousel-item">
                    <img className ="d-block w-100" style = {carousal_img} src="https://www.toolsgroup.com/wp-content/uploads/2019/05/wayfair8FAB1617-5056-A830-8C54D49938B6682C.jpg" alt="Second slide"/>
                </div>
                <div className ="carousel-item">
                    <img className ="d-block w-100" style = {carousal_img} src="https://www.toolsgroup.com/wp-content/uploads/2019/05/wayfair8FAB1617-5056-A830-8C54D49938B6682C.jpg" alt="Third slide"/>
                </div>
            </div>
            <a className ="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
                <span className ="carousel-control-prev-icon" aria-hidden="true"></span>
                <span className ="sr-only">Previous</span>
            </a>
            <a className ="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
                <span className ="carousel-control-next-icon" aria-hidden="true"></span>
                <span className ="sr-only">Next</span>
            </a>
        </div>
    )
}