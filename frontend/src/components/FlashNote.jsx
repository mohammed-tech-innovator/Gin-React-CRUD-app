import React from "react";

import ImgLogo from "../assets/Images/ImgLogo.png"
import AudioLogo from "../assets/Images/AudioLogo.png"


const FlashNote = ()=> {

    const photo = "https://wonderfulengineering.com/wp-content/uploads/2014/10/image-wallpaper-15.jpg";

    return(
        <div className="card">
            <div className="card-body">
                <div className="row inline">
                    <div className="col-lg-3 col-sm-3 col-3" >
                        <div className="mw-50 mh-50 text-center" >
                            <img src={photo} className="border border-dark-1 w-75 h-75 rounded-2 shadow-4" alt="Profile" />
                        </div>
                    </div>
                    <div className="col-lg-9 col-sm-9 col-9 mu-2">
                        <div className="row">
                            <form className="d-flex input-group">
                                <input
                                type="search"
                                id="form1"
                                className="col-lg-9 form-control cover" 
                                placeholder="Flash Note ⚡"
                                aria-label="Search"
                                />
                                <button type="submit" className="btn btn-outline-primary ml-2">
                                    Add
                                </button>
                            </form>
                        </div>
                        <div className="row mu-2">

                            <div className = "col-lg-4 col-md-4 col-sm-6 col-6 text-center">
                                <img className = "w-25" src = {ImgLogo} alt="Media"/>
                                <span className="text-info m-2">Media</span>
                            </div>

                            <div className = "col-lg-4 col-md-4 col-sm-6 col-6 text-center">
                                <img className = " w-25" src = {AudioLogo} alt="Audio"/>
                                <span className="text-info m-2" >Voice</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )

}

export default FlashNote;