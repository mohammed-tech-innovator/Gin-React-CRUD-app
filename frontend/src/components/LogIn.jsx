import React from 'react';

import Logo from '../assets/Images/Logo.jpg';
import Login from '../assets/Images/Login.webp'

const LogInForm = ({submitCallback, formData, handleInputChange, checkBox, errorMessage}) => {
  const cascadingLeftStyle = {
    marginLeft: '-50px',
  };

  const cardStyle = {
    background: 'hsla(0, 0%, 100%, 0.55)',
    backdropFilter: 'blur(30px)',
  };

  const style = `
  .cascading-left {
    margin-left: -50px;
  }

  @media (max-width: 991.98px) {
    .cascading-left {
      margin-left: 0;
    }


  @media (max-width: 1000px) {
      .sideimg {
        display: none;
      }
  }`

  return (
    <section className="text-center text-lg-start">
      {/* Cascading Right Style */}
      <style>{style}</style>

      {/* Jumbotron */}
      <div className="container py-5 my-2 ">
        <div className="row g-0 align-items-center">
            <div className="sideimg col-lg-4 mb-5 mb-lg-0">
                <img src={Login} className="w-100 rounded-4 shadow-4" alt="" />
            </div>
            <div className="col-lg-6 mb-5 mb-lg-0">
                
                <div className="card cascading-left" style={cardStyle}>
                <div className="card-body p-5 shadow-5 text-center">
                    <div className="row d-flex justify-content-center">
                        <div className="col-lg-3 mb-3 ">
                            <img src={Logo} className="w-100 rounded-4 shadow-4" alt="" />
                        </div>
                    </div>
                    <h2 className="fw-bold mb-2">Log In</h2>


                    <form onSubmit={submitCallback}>
                    {/* 2 column grid layout with text inputs for the first and last names */}


                    {/* Email input */}
                    <div className="form-outline mb-4">
                        <input type="email" id="form3Example3" placeholder='Email address' className="form-control" 
                        value={formData.Email} name='Email' onClick={handleInputChange} onChange={handleInputChange}/>
                    </div>

                    {/* Password input */}
                    <div className="form-outline mb-4">
                        <input type="password" id="form3Example4" placeholder="Password" className="form-control" 
                        value={formData.Password} name= "Password" onClick={handleInputChange} onChange={handleInputChange}/>
                        
                    </div>

                    

                    {/* Checkbox */}
                    <div className="form-check d-flex justify-content-center mb-4">
                        <input className="form-check-input me-2" type="checkbox" id="form2Example33" 
                        checked={formData.RememberDevice} name= "RememberDevice" 
                        onChange={checkBox}/>
                        <label className="form-check-label" htmlFor="form2Example33">
                        Remember this device
                        </label>
                    </div>
                    {errorMessage.isError && (<div className="alert alert-warning" role="alert">{errorMessage.message}</div>)}

                    {/* Submit button */}
                    <button type="submit" className="btn btn-primary btn-block mb-4">
                        Log In
                    </button>


                    {/* Register buttons */}
                    <div className="text-center">
                        <p>or Log In with:</p>
                        <button type="button" className="btn btn-link btn-floating mx-1">
                        <i className="fab fa-facebook-f"></i>
                        </button>

                        <button type="button" className="btn btn-link btn-floating mx-1">
                        <i className="fab fa-google"></i>
                        </button>

                        <button type="button" className="btn btn-link btn-floating mx-1">
                        <i className="fab fa-twitter"></i>
                        </button>

                        <button type="button" className="btn btn-link btn-floating mx-1">
                        <i className="fab fa-github"></i>
                        </button>
                        
                    </div>
                    </form>
                    <p>You don't have an account yet ? <a href="/signup/">Create a new account</a></p>
                </div>
                </div>
                <div className='text-center py-2'>
                  <p>Forgot your password ?<a href='google.com'>Click here</a></p>
                </div>
            </div>


            </div>
        </div>
        {/* Jumbotron */}
        </section>
  );
}

export default LogInForm;
