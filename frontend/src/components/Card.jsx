import React from "react";


export default function Card(props) {
    return (
        <div>
            <div className="card" style={{width: "18rem", margin: "10px"}}>
                <img src="..." className="card-img-top" alt="..."/>
                <div className="card-body">
                    <h5 className="card-title">{props.title}</h5>
                    <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    <a href="#" className="btn btn-primary">Go somewhere</a>
                </div>
            </div>
        </div>
    )
}