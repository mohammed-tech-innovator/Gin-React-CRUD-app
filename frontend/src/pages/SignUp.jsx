import React,{useState}from "react";
import api from '../api/Api'
import SignUpForm from "../components/SignUp";






export default function SignUp(){
    
    const handleSubmit = async (e) => {
        e.preventDefault()
        let response;
        try {

            response = await api.post("/signup/",formData)
            console.log(response.data)
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
            Fname:'',
            Lname:'',
            Email : '',
            Password: '',
            RPassword: '',
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
    return(
        <SignUpForm errorMessage = {errorMessage} handleSubmit = {handleSubmit} formData = {formData} handleInputChange = {handleInputChange}/>
    )
}