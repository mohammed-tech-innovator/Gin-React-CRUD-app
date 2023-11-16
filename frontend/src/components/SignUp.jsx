import React from 'react';
import SignUp from '../assets/Images/SignUp.jpg';
import Logo from '../assets/Images/Logo.jpg'

const SignUpForm = () => {
  const cascadingRightStyle = {
    marginRight: '-50px',
  };

  const cardStyle = {
    background: 'hsla(0, 0%, 100%, 0.55)',
    backdropFilter: 'blur(30px)',
  };

  const style = `
  .cascading-right {
    margin-right: -50px;
  }

  @media (max-width: 991.98px) {
    .cascading-right {
      margin-right: 0;
    }

  @media (max-width: 900px) {
      .sideimg {
        display: none;
      }
  }
`;

  return (
    <section className="text-center text-lg-start">
      {/* Cascading Right Style */}
      <style>{style}</style>

      {/* Jumbotron */}
      <div className="container py-4">
        <div className="row g-0 align-items-center">
          <div className="col-lg-6 mb-5 mb-lg-0">
            <div className="card cascading-right" style={cardStyle}>
              <div className="card-body p-5 shadow-5 text-center">
              <div className="row d-flex justify-content-center">
                        <div className="col-lg-3 mb-3 ">
                            <img src={Logo} className="w-100 rounded-4 shadow-4" alt="" />
                        </div>
                </div>
                <h2 className="fw-bold mb-5">Join us today</h2>
                
                <form>
                  {/* 2 column grid layout with text inputs for the first and last names */}
                  <div className="row">
                    <div className="col-md-6 mb-4">
                      <div className="form-outline">
                        <input type="text" id="form3Example1" className="form-control" placeholder="First name"/>
                      </div>
                    </div>
                    <div className="col-md-6 mb-4">
                      <div className="form-outline">
                        <input type="text" id="form3Example2" className="form-control" placeholder="Last name" />
                      </div>
                    </div>
                  </div>

                  {/* Email input */}
                  <div className="form-outline mb-4">
                    <input type="email" id="form3Example3" placeholder='Email address' className="form-control" />
                  </div>

                  {/* Password input */}
                  <div className="form-outline mb-4">
                    <input type="password" id="form3Example4" placeholder="Password" className="form-control" />
                    
                  </div>

                  <div className="form-outline mb-4">
                    <input type="password" id="form3Example4" placeholder="Repeat Password" className="form-control" />
                    
                  </div>

                  {/* Checkbox */}
                  <div className="form-check d-flex justify-content-center mb-4">
                    <input className="form-check-input me-2" type="checkbox" value="" id="form2Example33" defaultChecked />
                    <label className="form-check-label" htmlFor="form2Example33">
                      Remember this device
                    </label>
                  </div>

                  {/* Submit button */}
                  <button type="submit" className="btn btn-primary btn-block mb-4">
                    Sign up
                  </button>

                  {/* Register buttons */}
                  <div className="text-center">
                    <p>or sign up with:</p>
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
                <p>already have an account ? <a href="/login">then LogIn</a></p>
              </div>
            </div>
          </div>

          <div className="sideimg col-lg-6 mb-5 mb-lg-0">
            <img src={SignUp} className="w-100 rounded-4 shadow-4" alt="" />
          </div>
        </div>
      </div>
      {/* Jumbotron */}
    </section>
  );
}

export default SignUpForm;
