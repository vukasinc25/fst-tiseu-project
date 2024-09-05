import "./Competition.css"
import { useEffect, useState } from "react";
import { useHistory, useParams } from "react-router-dom";
import customFetch from "../intersceptor/interceptor";
import { RouteParams } from "../intefaces/routeParams";
import useRoles from "../role-base/userValidation";
import { useAuth0 } from "@auth0/auth0-react";

const Competition: React.FC = () => {
  const { id } = useParams<RouteParams>();
  const [competition, setCompetition] = useState<any>(null);
  const history = useHistory();
  const { hasRole } = useRoles();
  const { loginWithRedirect, logout, user, isLoading } = useAuth0();

  useEffect(() => {
    fetchCompetitions();
  }, []);

  const fetchCompetitions = async () => {
    try {
      const data = await customFetch(`http://localhost:8001/fakultet/competition/${id}`);
      setCompetition(data);
      console.log("Data: ",data)
    } catch (error) {
      console.error('Failed to fetch competitions:', error);
    }
  };

  function handleSubmit(e: { preventDefault: () => void; }) {
    e.preventDefault();    

    history.push(`/examResults/${competition._id}`);
  }

  function handleExamResul(e: { preventDefault: () => void; }) {
    e.preventDefault();    

    history.push({
      pathname: '/examResult',
      state: { competitionId: id },
    });
  }
    
  async function handleRegister(e: { preventDefault: () => void; }) {
    e.preventDefault();
    const userId = user?.sub?.split('|')[1];
    const userName = user?.name;

    try {
        // Await the result of the customFetch function to get the response
        const _ = await customFetch(`http://localhost:8001/fakultet/user/registerToCompetition/${id}/${userId}/${userName}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(null),  // Use null as the body
        });

        alert("Registered successfully");
        history.push('/competitions');
    } catch (error: any) {
        console.error('Failed to register to competition:', error);
        alert("An error occurred while registering for the competitio: " + error.message);
    }
}


  function handleCompetitonRequests(e: { preventDefault: () => void; }) {
    history.push(`/competitionRequests/${id}`)
  }

  const isDeadlinePassed = (deadline: string) => {
    const today = new Date();
    const deadlineDate = new Date(deadline);
    return deadlineDate < today;
  };

  return (
    <div className="competition-container">
      <h1>Competition Details</h1>
      {competition && (
        <div className="competition-details">
        <div className="detail">
          <strong>Competition ID:</strong>
          <span>{id}</span>
        </div>
        <div className="detail">
          <strong>Program Name:</strong>
          <span>{competition.programName}</span>
        </div>
        <div className="detail">
          <strong>Admission Requirements:</strong>
          <span>{competition.admissionRequirements}</span>
        </div>
        <div className="detail">
          <strong>Exam Date:</strong>
          <span>{competition.examDate}</span>
        </div>
        <div className="detail">
          <strong>Exam Format:</strong>
          <span>{competition.examFormat}</span>
        </div>
        <div className="detail">
          <strong>Exam Materials:</strong>
          <span>{competition.examMaterials}</span>
        </div>
        <div className="detail">
          <strong>Application Deadlines:</strong>
          <span>{competition.applicationDeadlines}</span>
        </div>
        <div className="detail">
          <strong>Application Documents:</strong>
          <span>{competition.applicationDocuments}</span>
        </div>
        <div className="detail">
          <strong>Application Method:</strong>
          <span>{competition.applicationMethod}</span>
        </div>
        <div className="detail">
          <strong>Application Contact:</strong>
          <span>{competition.applicationContact}</span>
        </div>
        <div className="detail">
          <strong>Tuition Fees:</strong>
          <span>{competition.tuitionFees}</span>
        </div>
        <div className="detail">
          <strong>Contact Information:</strong>
          <span>{competition.contactInformation}</span>
        </div>
        {isDeadlinePassed(competition.applicationDeadlines) && <p className="message">Applications are no longer possible because the deadline has passed</p>}
        <div className="button">
          {hasRole("STUDENT") && !isDeadlinePassed(competition.applicationDeadlines) && <button className="results" onClick={handleRegister}>Register</button>}
          {hasRole("PROFESSOR") && <button className="results" onClick={handleExamResul}>Add Results</button>}
          {hasRole("ADMIN") && <button className="results" onClick={handleCompetitonRequests}>Competition Requests</button>}
          <button onClick={handleSubmit}>Results</button>
        </div>
      </div>
      )}
    </div>
  );
  };
 
export default Competition;