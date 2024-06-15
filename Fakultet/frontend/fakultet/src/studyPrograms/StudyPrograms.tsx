import { Link } from "react-router-dom";
import "./StudyPrograms.css"
import { useEffect, useState } from "react";
import customFetch from "../intersceptor/interceptor";
const StudyPrograms = () => {
    const [studyPrograms, setStudyPrograms] = useState([]);

    useEffect(() => {
        fetchCompetitions();
    }, []);
    
    const fetchCompetitions = async () => {
        try {
            const data = await customFetch('http://localhost:8001/fakultet/studyPrograms');
            setStudyPrograms(data);
            console.log("Data: ",data)
        } catch (error) {
            console.error('Failed to fetch competitions:', error);
        }
    };

    return (  
    <div className="App1">
      <h1>StudyPrograms</h1>
      <div className="study-programs-grid">
        {studyPrograms.map((studyProgram: any) => (
           <div className="study-program-card">
           <h3 className="study-program-name">
             <Link className="study-program-link" to={`/studyProgram/${studyProgram._id}`}>{studyProgram.name}</Link>
           </h3>
           <p><strong>Study Level:</strong> {studyProgram.studyLevel}</p>
           <p><strong>Duration:</strong> {studyProgram.duration}</p>
           <p><strong>Internship:</strong> {studyProgram.internship ? 'Yes' : 'No'}</p>
           <p><strong>Accreditation:</strong> {studyProgram.accreditation ? 'Yes' : 'No'}</p>
           <p><strong>Contact Person ID:</strong> {studyProgram.contactPersonID}</p>
           <p><strong>Development Plan:</strong> {studyProgram.developmentPlan}</p>
         </div>
        ))}
      </div>
    </div>
    );
}
 
export default StudyPrograms;