import React from "react";
import Logo from '../assets/Images/Logo.jpg'

export default function PWRC({handleSubmit,formData,handleInputChange,errorMessage,sucMessage}) {
    const cascadingStyle = {
      };
    
    const cardStyle = {
        background: 'hsla(0, 0%, 100%, 0.55)',
        backdropFilter: 'blur(30px)',
      };
    return (
        <section className="d-flex align-items-center">
                <div className="container py-4 ">
                    <div className="row g-0 align-items-center justify-content-center">
                        <div className="col-lg-6 mb-5 mb-lg-0">
                            <div className="card cascading" style={cardStyle}>
                                <div className="card-body p-5 shadow-5 text-center">
                                    <div className="row d-flex justify-content-center">
                                        <div className="col-lg-3 mb-3">
                                            <img src={Logo} className="w-100 rounded-4 shadow-4" alt="" />
                                        </div>
                                    </div>
                                    <h2 className="fw-bold mb-5">Account Recovery</h2>

                                    <p>
                                        Now you can reset your password 😎.
                                    </p>

                                    {!sucMessage.isSuc&&(<form onSubmit={handleSubmit}>
                                        {/* password input */}
                                        <div className="form-outline mb-4">
                                            <input type="password" id="form3Example4" placeholder="New Password" className="form-control" 
                                            value={formData.NPassword1} onChange={handleInputChange} name='NPassword1'/>
                                            
                                        </div>
                                        {/* password input */}
                                        <div className="form-outline mb-4">
                                            <input type="password" id="form3Example3" placeholder="Repeat New Password" className="form-control" 
                                            value={formData.NPassword2} onChange={handleInputChange} name='NPassword2'/>
                                            
                                        </div>

                                        
                                    
                                        {errorMessage.isError && (<div class="alert alert-warning" role="alert">{errorMessage.message}</div>)}

                                        {/* Submit button */}
                                        <button type="submit" className="btn btn-outline-primary btn-block mb-4">
                                            Reset My Password
                                        </button>
                                    </form>)}
                                    {sucMessage.isSuc && (<div class="alert alert-success" role="alert">{sucMessage.message}</div>)}
                                </div>
                            </div>
                            <div className="text-center">
                                <a href="google.com">Help</a>
                            </div>
                        </div>
                    </div>
                </div>
            </section>

    )}