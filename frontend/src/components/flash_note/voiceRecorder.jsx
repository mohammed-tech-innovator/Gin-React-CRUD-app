import React from "react";



const VRecorder = ({voiceForm,UpdateVoiceForm}) => {
    return(
        <div className="position-fixed top-50 start-50 translate-middle card  col-lg-4 col-6" style={{ zIndex: "1" }}>
            <div className="container p-2">
                <div className="text-center my-2">
                    <span className="h5 text-info">Record Your VoiceNote</span>
                </div>
                <div className="mx-4">
                    <div className="my-1">
                        <span className="text-info">0.00 sec /120.00 sec</span>
                    </div>
                    <div className="progress">
                        <div className="progress-bar progress-bar-striped progress-bar-animated" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" style={{width: "75%"}}></div>
                    </div>
                    <div className="text-center my-2">
                        <button className="btn">Start ◀️</button>
                        <button className="btn">Pause ⏯️</button>
                        <button className="btn">Stop ⏹️</button>
                    </div>
                </div>
                <div className="row text-center">
                    <div className="col-6">
                        <button className="btn btn-outline-danger btn-sm"
                        onClick={()=>UpdateVoiceForm(false)} >❌ Cancle</button>
                    </div>

                    <div className="col-6">
                        <button className="btn btn-outline-success btn-sm">➕ Attach</button>
                    </div>
                </div>
            </div>
        </div>
    )
}


export default VRecorder;