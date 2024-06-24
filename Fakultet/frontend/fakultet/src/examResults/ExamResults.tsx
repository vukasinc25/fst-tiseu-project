import "./ExamResults.css"
import { useEffect, useState } from "react";
import customFetch from "../intersceptor/interceptor";
import { useParams } from "react-router-dom";
import {RouteParams} from "../intefaces/routeParams"

const ExamResults: React.FC = () => {
    const { id } = useParams<RouteParams>();
    const [examResults, setExamResults] = useState([]);

    useEffect(() => {
        fetchCompetitions();
      }, []);
    
      const fetchCompetitions = async () => {
        try {
          console.log("Id: ", id)
          const data = await customFetch(`http://localhost:8001/fakultet/user/getResultsByCompetitionId/${id}`);
          setExamResults(data);
          console.log("Data: ",data)
        } catch (error) {
          console.error('Failed to fetch competitions:', error);
        }
      };
    return ( 
      <div className="exam-results-container">
        <h2>Exam Results</h2>
        {examResults && (
        <table className="exam-results-table">
        <thead>
          <tr>
            <th>User Name</th>
            <th>Score</th>
            <th>Score Entry Date</th>
          </tr>
        </thead>
        <tbody>
          {examResults.map((result: any) => (
            <tr key={result._id}>
              <td>{result.userName}</td>
              <td>{result.score}</td>
              <td>{new Date(result.scoreEntryDate).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
        )}
    </div>
    );
}
 
export default ExamResults;