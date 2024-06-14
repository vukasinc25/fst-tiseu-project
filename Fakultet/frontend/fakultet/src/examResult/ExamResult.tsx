import { useState } from "react";
import customFetch from "../intersceptor/interceptor";
import "./ExamResult.css"
import { useLocation } from "react-router-dom";
interface ExamResultState {
    competitionId: string;
  }
const ExamResult = () => {
    const location = useLocation();
    const state = location.state as ExamResultState;
    const competitionId = state?.competitionId || '';
    const [formData, setFormData] = useState({
        userName: '',
        competitionId: competitionId,
        score: '',
      });
    
    const clearForm = () => {
    setFormData({
        userName: '',
        competitionId: competitionId,
        score: '',
    });
    };
    
    const handleChange = (e: { target: { name: any; value: any; }; }) => {
    const { name, value } = e.target;
    setFormData({
        ...formData,
        [name]: value,
    });
    };
      
    const handleSubmit = async (e: { preventDefault: () => void; }) => {
        e.preventDefault();
        try {
          const response = await customFetch('http://localhost:8001/fakultet/user/examResults', {
            method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
          });
          console.log('Exam result created:', response.data);
          clearForm();
        } catch (error) {
          console.error('There was an error creating the exam result!', error);
        }
      };
    return ( 
        <div>
        <h2>Create Exam Result</h2>
        <form onSubmit={handleSubmit}>
          <div>
            <label>
              Student User Name:
              <input type="text" name="userName" value={formData.userName} onChange={handleChange} required />
            </label>
          </div>
          <div>
            <label>
              Score:
              <input type="text" name="score" value={formData.score} onChange={handleChange} required />
            </label>
          </div>
          <div>
            <button type="submit">Submit</button>
          </div>
        </form>
      </div>
    );
}
 
export default ExamResult;