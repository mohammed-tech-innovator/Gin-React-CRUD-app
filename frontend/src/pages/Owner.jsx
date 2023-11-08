import React,{useState, useEffect} from "react";
import axios from "axios";
import { useParams } from "react-router-dom";
import Profile from "../components/Profile";



export default function Owner(props){

    const urlparam = useParams();
    const id = urlparam.id;
    
    const [owner ,setData] = useState([])
    useEffect(() => {
        axios.get(`http://localhost:8000/owner/${id}`)
        .then(res => setData(res.data))
        .catch(err => console.log(err));
    },[]);




    return(
        <div>

            <Profile onr={owner}/>
            
            
        </div>
    )

    
}