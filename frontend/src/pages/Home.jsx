import React, { useState, useEffect } from 'react';
import ReactPaginate from 'react-paginate';
import Navbar from "../components/Navbar";
import Carousal from '../components/Carousal';
import Card from "../components/Card";
import Sidebar from "../components/Sidebar";
import Footer from '../components/Footer';
import LogInForm from '../components/LogInForm';
import axios from 'axios';

export default function Home() {

    const [estates,setData] = useState([])
    useEffect(() => {
        axios.get('http://localhost:8000/estate')
        .then(res => setData(res.data))
        .catch(err => console.log(err));
    },[]);


    return (

        <div>

            <Navbar com_name = {"Buyut"} logo ={"https://th.bing.com/th/id/OIP.l590m8rOs0GW_XV2hHgRYAHaHa?pid=ImgDet&rs=1"}/>

            <Card items = {estates} />

        </div>

        

    )
}

