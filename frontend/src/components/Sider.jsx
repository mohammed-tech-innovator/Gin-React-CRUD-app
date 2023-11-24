import React from "react";


const Sider = () => {

    const photo = "https://wonderfulengineering.com/wp-content/uploads/2014/10/image-wallpaper-15.jpg";

    const cardText = {
        fontSize:"0.8rem",
    }
    return(
        <>
            <div className="card rounded-3">
                <div className="col">
                    <div className="card rounded-3">
                        <div className = "overflow-hidden" style={{maxHeight:"70px"}}>
                            <img src = {photo} className="col-lg-12 rounded-3" alt="cover"/>
                        </div>

                        <div className = "position-absolute top-100 start-50 translate-middle" style={{zIndex:"1"}}>
                            <img src = {photo} className="border border-secondary-1 col-lg-12 rounded-circle" alt = "Profile"/>
                        </div>
                    </div>
                    <div className="card pt-5">

                        <h6 className="text-center  text-secondary text-sm-left" style={cardText}>Mohammed Yahya Yousif</h6>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>

                    </div>

                    
                </div>
            
            </div>

            <div className="card mt-2">

                        <h6 className="text-center  text-secondary text-sm-left" style={cardText}>Mohammed Yahya Yousif</h6>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>
                        <hr className="mt-2 mb-3"/>
                        <span className="text-secondary text-xs-left mx-1" style={cardText}>Mohammed Yahya Yousif</span>

            </div>
        </>
        
    )
}


export default Sider;