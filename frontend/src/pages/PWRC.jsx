import React ,{useState} from "react";
import api from "../api/Api";
import PWRC from "../components/PWRC";
import { useParams } from "react-router-dom";

export default function PWRchange() {

    const {hash,email} = useParams();

    console.log(typeof hash)
    console.log(typeof email)

    const [formData,setFormData] = useState(
        {
            NPassword1:"",
            NPassword2:"",
        }
    );

    const [errorMessage,setErrorMessage] = useState(
        {
            message:"",
            isError:false,
        }
    );

    const handleSubmit = async (e) => {
        e.preventDefault()
        let response;
        try {

            response = await api.post("/recovery/" + hash +"/"+email,formData)
    
            setErrorMessage(() => ({
                message: '',
                isError: false,
            }))
            

        } catch (err) {

            console.log(`Error:${err.message},${err.response.data.Error}`)
            setErrorMessage(() => ({
                message: err.response.data.tag,
                isError: true,
            }))

        }
    };

    const handleInputChange = async (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
          ...prevData,
          [name]: value,
          
        }));

    };

    return(<PWRC handleSubmit = {handleSubmit} formData = {formData} handleInputChange={handleInputChange} 
    errorMessage = {errorMessage} />);
}