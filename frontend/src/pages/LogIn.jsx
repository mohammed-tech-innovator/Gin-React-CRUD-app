import React, { useState } from "react";
import api from '../api/Api'
import LogInForm from "../components/LogIn";


export default function LogIn() {
    const handleSubmit = async (e) => {
        e.preventDefault()
        let response;
        try {
            console.log(formData)
            response = await api.post("/login/",formData)
            
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
    }
    const [errorMessage,setErrorMessage] = useState(
        {
            message : "",
            isError : false,
        }
    )
    const [formData, setFormData] = useState(
        {
            Email : '',
            Password: '',
            RememberDevice: true,
        }
    )
    const handleInputChange = async (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
          ...prevData,
          [name]: value,
          
        }));
      };
    const checkBox = async (e)=>{
        const { name, checked } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: checked,
        }));
    }
    return (
        <LogInForm submitCallback = {handleSubmit} formData = {formData} handleInputChange={handleInputChange} checkBox={checkBox}
        errorMessage = {errorMessage} />
    )
}