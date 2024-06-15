import { useEffect, useState } from "react";
import "./StudyProgram.css"
import customFetch from "../intersceptor/interceptor";
import { useParams } from "react-router-dom";
import { RouteParams } from "../intefaces/routeParams";
const StudyProgram = () => {
    const { id } = useParams<RouteParams>();
    const [studyProgram, setStudyProgram] = useState<any>(null);

    useEffect(() => {
        fetchCompetitions();
    }, []);
    
    const fetchCompetitions = async () => {
        try {
          const data = await customFetch(`http://localhost:8001/fakultet/studyProgram/${id}`);
          setStudyProgram(data);
          console.log("Data: ",data)
        } catch (error) {
          console.error('Failed to fetch competitions:', error);
        }
    };

    return (  
        <div className="study-program-container">
      <h1>Study Program Details</h1>
      {studyProgram && (
        <div className="study-program-details">
          <div className="detail">
            <strong>Program ID:</strong>
            <span>{studyProgram._id}</span>
          </div>
          <div className="detail">
            <strong>Name:</strong>
            <span>{studyProgram.name}</span>
          </div>
          <div className="detail">
            <strong>Study Level:</strong>
            <span>{studyProgram.studyLevel}</span>
          </div>
          <div className="detail">
            <strong>Duration:</strong>
            <span>{studyProgram.duration}</span>
          </div>
          <div className="detail">
            <strong>Objectives:</strong>
            <span>{studyProgram.objectives}</span>
          </div>
          <div className="detail">
            <strong>Program Structure:</strong>
            <span>{studyProgram.programStructure}</span>
          </div>
          <div className="detail">
            <strong>Internship:</strong>
            <span>{studyProgram.internship ? "Yes" : "No"}</span>
          </div>
          <div className="detail">
            <strong>Graduation Requirements:</strong>
            <span>{studyProgram.graduationRequirements}</span>
          </div>
          <div className="detail">
            <strong>Accreditation:</strong>
            <span>{studyProgram.accreditation ? "Yes" : "No"}</span>
          </div>
          <div className="detail">
            <strong>Contact Person ID:</strong>
            <span>{studyProgram.contactPersonID}</span>
          </div>
          <div className="detail">
            <strong>Development Plan:</strong>
            <span>{studyProgram.developmentPlan}</span>
          </div>
          <div className="detail">
            <strong>Department ID:</strong>
            <span>{studyProgram.departmentID}</span>
          </div>
        </div>
      )}
    </div>
  );
}
 
export default StudyProgram;