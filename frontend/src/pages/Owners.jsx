import React, { useEffect, useState } from "react";
import axios from "axios";


const spinner_style = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",

}

const body_card_style = {
    margin: "5px",
    padding: "50px",
    backgroundColor: "#dddddd",
    borderRadius:"10px",
    minHeight:"1000px",
}

const card_style = {
    margin:"5px",
    borderRadius:"10px",   
}

export default function Owners() {
  const [owners, setOwners] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    axios
      .get("http://localhost:8000/owners/")
      .then((res) => {
        setOwners(res.data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err);
        setLoading(false);
      });
  }, []);



  if (loading) {
    return (
        <>
            <div className="spinner-border container" role="status" style ={spinner_style}>
                <span className="visually-hidden text-center">Loading...</span>
            </div>
        </>
    );
  }

  if (error) {
    return (<div className="container">
    <div className="alert alert-danger" role="alert">error ! : {error.tag}</div>
    </div>);
  }


    const CardsList = owners.map((item, index) => {
        return (
        <div key={index} className="col-lg-5 col-md-12 col-sm-12 card" style={card_style}>
            <div className="row no-gutters">
            <div className="col-sm-4 my-2">
                <img className="card-img rounded-circle" src={item.photo} alt={`${item.name}'s Card`} />
            </div>
            <div className="col-sm-7">
                <div className="card-body">
                <h5 className="card-title">{item.name}</h5>
                <p className="card-text">{item.bio}</p>
                <p className="card-text">{item.address}</p>
                <a href={`/owner/${item._id}`} className="btn btn-primary">
                    View Profile
                </a>
                </div>
            </div>
            </div>
        </div>
        );
    });
    
    return (
        <div style={body_card_style}>
            <div className="row">
                <div className="col">
                    <nav aria-label="breadcrumb" className="bg-light rounded-3 p-3 mb-4">
                    <ol className="breadcrumb mb-0">
                        <li className="breadcrumb-item"><a href="/">Home</a></li>
                        <li className="breadcrumb-item active" aria-current="page">Owners</li>
                    </ol>
                    </nav>
                </div>
            </div>
            <div className="row px-3" >
                    {owners.length > 0 ? CardsList : <p>No data available.</p>}
                    {owners.length > 0 ? CardsList : <p>No data available.</p>}
                    {owners.length > 0 ? CardsList : <p>No data available.</p>}
            </div>
        </div>
    );
  


 }
