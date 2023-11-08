import React from "react";

export default function Card(props) {
    const iconStyle = {
        margin: "10px",
        textAlign : "center",
    }
    const CardsList = props.items.map((item, index) => {
        return (
            <div key={index} className = " col-xl-3 col-lg-4 col-md-6 col-sm-8 col-xs-12 mb-4">
                <div className="card" style={{ width: "18rem", margin: "10px" }}>
                    <img src={item.estate.photos} className="card-img-top" alt="..." />
                    <div className="card-body">
                        <h5 className="card-title">{item.estate.title}</h5>
                        <p className="card-text">property price :  <strong>{item.estate.price}</strong> <small>{item.estate.currency}</small></p>
                        <small className="card-text">owned by :  <a href={`/owner/${item.owner._id}`}><strong>{item.owner.name}</strong></a></small>
                        <div className="text-center">
                            <div className="row justify-content-center align-items-center" style={{ display: 'flex' }}>
                                <i className="col-1 fas fa-bed" style={iconStyle}>{item.estate.features.bedrooms}</i>
                                <i className="col-1 fas fa-toilet" style={iconStyle}>{item.estate.features.bathrooms}</i>
                                {item.estate.features.pool && (<i className="col-1 fas fa-swimmer" style={iconStyle}>1</i>)}
                            </div>
                        </div>


                        <a href={"/estate/" + item.estate._id} className="btn btn-primary">Go somewhere</a>
                    </div>
                </div>
            </div>
        );
    });

    const CardsDivStyle = {
        backgroundColor: "lightgray",
        padding: "10px",
        margin: "10px",
    };

    return (
        <div className="card text-center bg-light p-3" style={CardsDivStyle}>
            <div className="row">
                {CardsList}
            </div>
        </div>
    );
}
