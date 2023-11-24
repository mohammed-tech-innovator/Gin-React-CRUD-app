import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import api from "../api/Api";


export default function VEmail() {
    const { hash, email } = useParams();
    const cardStyle = {
        background: 'hsla(0, 0%, 100%, 0.55)',
        backdropFilter: 'blur(30px)',
    };

    const [message, setMessage] = useState({
        message: '',
        state: false,
    });

    useEffect(() => {
        verifyEmail();
    }, [hash, email]);

    const verifyEmail = () => {
        api.get(`/verify-email/${hash}/${email}`)
            .then((response) => {
                setMessage(() => ({
                    message: response.data.tag,
                    state: true,
                }));
            })
            .catch((err) => {
                setMessage(() => ({
                    message:  err.response.data.tag,
                    state: false,
                }));
            });
    };

    console.log(message.message)
    
    return(
        <>
            <section className="d-flex align-items-center ">
                <div className="container py-4 ">
                    <div className="row g-0 align-items-center justify-content-center">
                        <div className="col-lg-6 mb-5 mb-lg-0">
                            <div className="card cascading" style={cardStyle}>
                                <div className="card-body p-5 shadow-5 text-center">
                                {
                                    message.state
                                        ? <div className="alert alert-success" role="alert">{message.message}</div>
                                        : <div className="alert alert-warning" role="alert">{message.message}</div>
                                }




                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </>
    )
}