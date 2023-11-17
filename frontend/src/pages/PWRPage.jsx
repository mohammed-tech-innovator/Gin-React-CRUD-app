import PWrecovery from "../components/PWrecovery";
import react,{useState, useEffect} from 'react';
import api from "../api/Api";

export default function PWRPage() {
    const [formData,setFormData] = useState({
        Email:"",
    })
    const [errorMessage ,setErrorMessage] = useState({
        message:"",
        isError:false,
    })
    const handleSubmit = async (e) => {
        e.preventDefault()
        let response;
        try{
            response = await api.post("/recover/",formData)
            console.log(response.data)
            setErrorMessage(() => ({
                message: '',
                isError: false,
            }))
        } catch(err){
            console.log(`Error:${err.message},${err.response.data.Error}`)
            setErrorMessage(() => ({
                message: err.response.data.tag,
                isError: true,
            }))
        }
        
    }

    const handleInputChange = async (e) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value,
        }));
    }
    
  return(
    <>
        <PWrecovery handleSubmit = {handleSubmit} formData = {formData} handleInputChange={handleInputChange} errorMessage={errorMessage}/>
    </> 
    )
}