import React, { useState } from 'react';
import ReactPaginate from 'react-paginate';
import Navbar from "../components/Navbar";
import Card from "../components/Card";
import Sidebar from "../components/Sidebar";

export default function Home() {
    const cardList = Array.from({ length: 100 }, (_, index) => (
        <Card key={index} />
      ));
    return (
        <div className="app-container">
            {/* Navbar component */}
            <Navbar />

            <div className="content-container">
                {/* Sidebar component */}
                <Sidebar />

                <div className="main-content">
                <div className="card-container">
                    {/* Cards */}
                    {cardList}
                </div>
                </div>
            </div>
        </div>
    )
}

