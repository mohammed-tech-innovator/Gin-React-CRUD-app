import React,{useState} from "react";

import VRecorder from "./flash_note/voiceRecorder";
import MUpploader from "./flash_note/mediaUpploader";


const FlashNote = ({fastNoteForm})=> {

    const photo = "https://wonderfulengineering.com/wp-content/uploads/2014/10/image-wallpaper-15.jpg";
    const [voiceForm,setVForm] = useState(() => {return false});


    const UpdateVoiceForm = (val) => {
        setVForm(val)
    }


    const [mediaForm,setMForm] = useState(() => {return false});


    const UpdateMediaForm = (val) => {
        setMForm(val)
    }

    return(
        <>
            <div className="card">
                <div className="card-body pb-0">
                    <div className="row inline">
                        <div className="col-lg-12 col-sm-12 col-12">
                            <div className="row">
                                <form className="d-flex input-group">
                                    <input
                                    type="search"
                                    id="form1"
                                    className="col-lg-9 form-control cover" 
                                    placeholder="Flash Note ⚡"
                                    aria-label="Search"
                                    value={fastNoteForm.noteBody}
                                    />
                                    <button type="submit" className="btn btn-outline-primary ml-2">
                                        Add
                                    </button>
                                </form>
                            </div>
                            <div className="row d-flex align-items-center">

                                <div className = "col-lg-4 col-md-4 col-sm-4 col-4 text-center">
                                    <button className="btn text-info btn-sm h6" onClick={()=>UpdateMediaForm(true)}><span className="h2">📸</span>Media</button>
                                </div>

                                <div className = "col-lg-4 col-md-4 col-sm-4 col-4 text-center ">
                                    <button className="btn text-info btn-sm h6" onClick={()=>UpdateVoiceForm(true)}><span className="h2">🔉</span>Voice</button>
                                </div>

                                <div className="col-lg-4 col-md-4 col-sm-4 col-4 d-flex align-items-center">

                                    <div className="form-check form-switch ">
                                        <input className="form-check-input form-switch-lg" type="checkbox" id="flexSwitchCheckDefault"
                                        checked={fastNoteForm.public}/>
                                        <span className="form-check-label text-info">Public</span>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            {voiceForm && (<VRecorder UpdateVoiceForm={UpdateVoiceForm}/>)}
            {mediaForm && (<MUpploader UpdateMediaForm={UpdateMediaForm}/>)}

        </>
    )

}

export default FlashNote;