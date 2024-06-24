import React, { useState, useEffect } from "react";
import './Competitions.css';
import customFetch from "../intersceptor/interceptor";
import { Link, useHistory } from "react-router-dom";

const Competitions = () => {
  const [competitions, setCompetitions] = useState([]);
  const history = useHistory();
  // const [formData, setFormData] = useState({
  //   programName: '',
  //   admissionRequirements: '',
  //   examDate: '',
  //   applicationMethod: '',
  //   tuitionFees: ''
  // });

  useEffect(() => {
    fetchCompetitions();
  }, []);

  const fetchCompetitions = async () => {
    try {
      const data = await customFetch('http://localhost:8001/fakultet/competitions');
      setCompetitions(data);
      console.log("Data: ",data)
    } catch (error) {
      console.error('Failed to fetch competitions:', error);
    }
  };

  // const handleSubmit = async (event: { preventDefault: () => void; }) => {
  //   event.preventDefault();
  //   try {
  //     // Send formData to backend
  //     await customFetch('http://localhost:8001/fakultet/createCompetition', {
  //       method: 'POST',
  //       headers: {
  //         'Content-Type': 'application/json',
  //       },
  //       body: JSON.stringify(formData),
  //     });
  //     // Fetch competitions again after creating competition
  //     fetchCompetitions();
  //     // Reset form data
  //     setFormData({
  //       programName: '',
  //       admissionRequirements: '',
  //       examDate: '',
  //       applicationMethod: '',
  //       tuitionFees: ''
  //     });
  //   } catch (error) {
  //     console.error('Failed to create competition:', error);
  //   }
  // };

  // const handleChange = (event: { target: { name: any; value: any; }; }) => {
  //   const { name, value } = event.target;
  //   setFormData({
  //     ...formData,
  //     [name]: value
  //   });
  // };
  function handleSubmit(e: { preventDefault: () => void; }) {
    e.preventDefault();    
    history.push("/competition");
  }

  return (
    <div className="App1">
      <h2>Competitions List</h2>
      <div className="button-container">
        <button className="createCompetition" onClick={handleSubmit}>Create Competition</button>
      </div>
      <div className="competitions-grid">
        {competitions.map((competition: any) => (
          <div key={competition._id} className="competition-card">
            <h3 className="competition-name">
              <Link className="competition-link" to={`/competition/${competition._id}`}>{competition.programName}</Link>
            </h3>
            <p><strong>Admission Requirements:</strong> {competition.admissionRequirements}</p>
            <p><strong>Exam Date:</strong> {competition.examDate}</p>
            <p><strong>Application Method:</strong> {competition.applicationMethod}</p>
            <p><strong>Tuition Fees:</strong> {competition.tuitionFees}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Competitions;
