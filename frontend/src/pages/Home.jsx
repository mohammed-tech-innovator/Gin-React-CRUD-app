import React, { useState, useEffect } from 'react';
import ReactPaginate from 'react-paginate';

import Carousal from '../components/Carousal';
import Card from "../components/Card";
import Sidebar from "../components/Sidebar";
import Footer from '../components/Footer';
import LogInForm from '../components/LogInForm';
import axios from 'axios';

export default function Home() {

    const [estates, setEstates] = useState([]);
    const [estate_owner, setOwners] = useState([]);
   


    useEffect(() => {
        axios.get('http://localhost:8000/estate')
        .then((res) => setEstates(res.data))
        .catch((err) => console.log(err));
    }, []);

  useEffect(() => {
    const fetchOwners = async () => {
        const ownersData = [];
        for (let estate of estates) {
            try {
                const response = await axios.get(`http://localhost:8000/owner/${estate.owner_id}`);
                ownersData.push({
                "estate":estate,
                "owner":response.data,
            })

            } catch (error) {
                console.log(error);
            }
        }
        setOwners(ownersData);
    };
    fetchOwners();}, [estates]);

  console.log(estate_owner);


    return (
        <div>

            <Carousal/>

            <Card items = {estate_owner} />

        </div>
    )
}

