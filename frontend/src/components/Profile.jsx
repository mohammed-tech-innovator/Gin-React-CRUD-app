import React from 'react';

const Profile = (props) => {

  const owner = props.onr;
  let Maxprice = 3000000
  let priceRatio = 100*(owner.avgprice/Maxprice)
  let Maxest = 50
  let estratio = 100*(owner.numofest/Maxest)
  let Maxrestime = 300
  let restimeratio = 100*(owner.avgrestim/Maxrestime)
  
  


  return (
    <section style={{ backgroundColor: '#eee' }}>
      <div className="container py-5">
        <div className="row">
          <div className="col">
            <nav aria-label="breadcrumb" className="bg-light rounded-3 p-3 mb-4">
              <ol className="breadcrumb mb-0">
                <li className="breadcrumb-item"><a href="/">Home</a></li>
                <li className="breadcrumb-item"><a href="/Owners/">Owners</a></li>
                <li className="breadcrumb-item active" aria-current="page">User Profile</li>
              </ol>
            </nav>
          </div>
        </div>

        <div className="row">
          <div className="col-lg-4">
            <div className="card mb-4">
              <div className="card-body text-center">
                <img src={owner.photo}
                  alt="avatar" className="rounded-circle img-fluid" style={{ width: '150px' }} />
                <h5 className="my-3">{owner.name}</h5>
                <p className="text-muted mb-1">{owner.bio}</p>
                <p className="text-muted mb-4">{owner.address}</p>
                <div className="d-flex justify-content-center mb-2">
                  <button type="button" className="btn btn-primary">Show All Offers</button>
                  <button type="button" className="btn btn-outline-primary ms-1">Message</button>
                </div>
              </div>
            </div>
            <div className="card mb-4 mb-lg-0">
              <div className="card-body p-0">
                <ul className="list-group list-group-flush rounded-3">
                  <li className="list-group-item d-flex justify-content-between align-items-center p-3">
                    <i className="fas fa-globe fa-lg text-warning"></i>
                    <a href = {owner.website} className="mb-0">{owner.website}</a>
                  </li>
            
                  <li className="list-group-item d-flex justify-content-between align-items-center p-3">
                    <i className="fab fa-twitter fa-lg" style={{ color: '#55acee' }}></i>
                    <a href = {owner.x} className="mb-0">{owner.x}</a>
                  </li>
                  <li className="list-group-item d-flex justify-content-between align-items-center p-3">
                    <i className="fab fa-instagram fa-lg" style={{ color: '#ac2bac' }}></i>
                    <a href = {owner.instagram} className="mb-0">{owner.instagram}</a>
                  </li>
                  <li className="list-group-item d-flex justify-content-between align-items-center p-3">
                    <i className="fab fa-facebook-f fa-lg" style={{ color: '#3b5998' }}></i>
                    <a href = {owner.facebook} className="mb-0">{owner.facebook}</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <div className="col-lg-8">
            <div className="card mb-4">
              <div className="card-body">
                <div className="row">
                  <div className="col-sm-3">
                    <p className="mb-0">Full Name</p>
                  </div>
                  <div className="col-sm-9">
                    <p className="text-muted mb-0">{owner.name}</p>
                  </div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-3">
                    <p className="mb-0">Email</p>
                  </div>
                  <div className="col-sm-9">
                    <p className="text-muted mb-0">{owner.email}</p>
                  </div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-3">
                    <p className="mb-0">Phone</p>
                  </div>
                  <div className="col-sm-9">
                    <p className="text-muted mb-0">{owner.phone}</p>
                  </div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-3">
                    <p className="mb-0">Mobile</p>
                  </div>
                  <div className="col-sm-9">
                    <p className="text-muted mb-0">{owner.mobile}</p>
                  </div>
                </div>
                <hr />
                <div className="row">
                  <div className="col-sm-3">
                    <p className="mb-0">Address</p>
                  </div>
                  <div className="col-sm-9">
                    <p className="text-muted mb-0">{owner.address}</p>
                  </div>
                </div>
              </div>
            </div>
            <div className="row">
              <div className="col-lg-12 col-md-12">
                <div className="card mb-4 mb-md-0">
                  <div className="card-body">
                    <p className="mb-4"><span className="text-primary font-italic me-1">Statistics</span>Estates</p>
                    <p className="mb-1" style={{ fontSize: '.77rem' }}>Avarage Price: {owner.avgprice} vs Max Avarage: {Maxprice}</p>
                    <div className="progress rounded" style={{ height: '5px' }}>
                      <div className="progress-bar" role="progressbar" style={{ width: `${priceRatio}%` }} aria-valuenow="80"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <p className="mt-4 mb-1" style={{ fontSize: '.77rem' }}>Total Estates : {owner.numofest} vs Max Num of Estates : {Maxest}</p>
                    <div className="progress rounded" style={{ height: '5px' }}>
                      <div className="progress-bar" role="progressbar" style={{ width: `${estratio}%` }} aria-valuenow="72"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <p className="mt-4 mb-1" style={{ fontSize: '.77rem' }}>Avg Responce time : {owner.avgrestime}, vs Max Responce time: {Maxrestime}</p>
                    <div className="progress rounded" style={{ height: '5px' }}>
                      <div className="progress-bar" role="progressbar" style={{ width: `${restimeratio}%` }} aria-valuenow="89"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <p className="mt-4 mb-1" style={{ fontSize: '.77rem' }}>Rating :{owner.rating} / 5.00</p>
                    <div className="progress rounded" style={{ height: '5px' }}>
                      <div className="progress-bar" role="progressbar" style={{ width: `${20*(owner.rating)}%` }} aria-valuenow="55"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                </div>
              </div>
              
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Profile;
