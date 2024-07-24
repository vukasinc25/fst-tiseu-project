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
          var sorted = data.sort((a:any,b:any) => a.score < b.score);
          setExamResults(sorted);
          console.log("Data: ",data)
        } catch (error) {
          console.error('Failed to fetch competitions:', error);
        }
      };
      return (
        <div className="exam-results-container">
          <h2>Exam Results</h2>
          <p>Students after the red line did not make the cut</p>
          {examResults && (
            <table className="exam-results-table">
              <thead>
                <tr>
                  <th></th>
                  <th>User Name</th>
                  <th>Score</th>
                  <th>Score Entry Date</th>
                </tr>
              </thead>
              <tbody>
                {examResults.map((result: any, index: number) => (
                  <tr key={result._id} className={index === 5 ? "red-line" : ""}>
                    <td>{index+1+"."}</td>
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