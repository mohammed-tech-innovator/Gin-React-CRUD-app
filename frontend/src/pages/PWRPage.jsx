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
    const [sucMessage ,setSucMessage] = useState({
        message:"",
        isSuc:false,
    })
    const handleSubmit = async (e) => {
        e.preventDefault()
        let response;
        try{
            response = await api.post("/recover/",formData)

            setErrorMessage(() => ({
                message: '',
                isError: false,
            }))
            setSucMessage( ()=> ({
                message:response.data.tag,
                isSuc: true,
            }))
        } catch(err){
            console.log(`Error:${err.message},${err.response.data.Error}`)
            setErrorMessage(() => ({
                message: err.response.data.tag,
                isError: true,
            }));
            setSucMessage( ()=> ({
                message:'',
                isSuc: false,
            }));
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
        <PWrecovery handleSubmit = {handleSubmit} formData = {formData} handleInputChange={handleInputChange} errorMessage={errorMessage}
        sucMessage={sucMessage}/>
    </> 
    )
}