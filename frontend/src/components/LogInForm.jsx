import React from "react";

class LogInForm extends React.Component {
    render() {
        return (
            <div>
                <div style={{ width: "26rem" }}>
                    <ul className="nav nav-pills nav-justified mb-3" id="ex1" role="tablist">
                        <li className="nav-item" role="presentation">
                            <a className="nav-link active" id="tab-login" data-mdb-toggle="pill" href="#pills-login" role="tab" aria-controls="pills-login" aria-selected="true">Login</a>
                        </li>
                        <li className="nav-item" role="presentation">
                            <a className="nav-link" id="tab-register" data-mdb-toggle="pill" href="#pills-register" role="tab" aria-controls="pills-register" aria-selected="false" tabIndex="-1">Register</a>
                        </li>
                    </ul>

                    <div className="tab-content">
                        <div className="tab-pane fade active show" id="pills-login" role="tabpanel" aria-labelledby="tab-login">
                            <form>
                                <div className="text-center mb-3">
                                    <p>Sign in with:</p>
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

                                <p className="text-center">or:</p>

                                <div className="form-outline mb-4">
                                    <input type="email" id="loginName" className="form-control" />
                                    <label className="form-label" htmlFor="loginName" style={{ marginLeft: "0px" }}>Email or username</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "114.4px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="password" id="loginPassword" className="form-control" />
                                    <label className="form-label" htmlFor="loginPassword" style={{ marginLeft: "0px" }}>Password</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "64.8px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="row mb-4">
                                    <div className="col-md-6 d-flex justify-content-center">
                                        <div className="form-check mb-3 mb-md-0">
                                            <input className="form-check-input" type="checkbox" value="" id="loginCheck" checked="" />
                                            <label className="form-check-label" htmlFor="loginCheck"> Remember me </label>
                                        </div>
                                    </div>

                                    <div className="col-md-6 d-flex justify-content-center">
                                        <a href="#!">Forgot password?</a>
                                    </div>
                                </div>

                                <button type="submit" className="btn btn-primary btn-block mb-4">Sign in</button>

                                <div className="text-center">
                                    <p>Not a member? <a href="#!">Register</a></p>
                                </div>
                            </form>
                        </div>
                        <div className="tab-pane fade" id="pills-register" role="tabpanel" aria-labelledby="tab-register">
                            <form>
                                <div className="text-center mb-3">
                                    <p>Sign up with:</p>
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

                                <p className="text-center">or:</p>

                                <div className="form-outline mb-4">
                                    <input type="text" id="registerName" className="form-control" />
                                    <label className="form-label" htmlFor="registerName" style={{ marginLeft: "0px" }}>Name</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "42.4px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="text" id="registerUsername" className="form-control" />
                                    <label className="form-label" htmlFor="registerUsername" style={{ marginLeft: "0px" }}>Username</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "66.4px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="email" id = "registerEmail" className="form-control" />
                                    <label className="form-label" htmlFor="registerEmail" style={{ marginLeft: "0px" }}>Email</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "40px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="password" id="registerPassword" className="form-control" />
                                    <label className="form-label" htmlFor="registerPassword" style={{ marginLeft: "0px" }}>Password</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "64.8px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-outline mb-4">
                                    <input type="password" id="registerRepeatPassword" className="form-control" />
                                    <label className="form-label" htmlFor="registerRepeatPassword" style={{ marginLeft: "0px" }}>Repeat password</label>
                                    <div className="form-notch"><div className="form-notch-leading" style={{ width: "9px" }}></div><div className="form-notch-middle" style={{ width: "106.4px" }}></div><div className="form-notch-trailing"></div></div>
                                </div>

                                <div className="form-check d-flex justify-content-center mb-4">
                                    <input className="form-check-input me-2" type="checkbox" value="" id="registerCheck" checked="" aria-describedby="registerCheckHelpText" />
                                    <label className="form-check-label" htmlFor="registerCheck">
                                        I have read and agree to the terms
                                    </label>
                                </div>

                                <button type="submit" className="btn btn-primary btn-block mb-3">Sign in</button>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

export default LogInForm;
