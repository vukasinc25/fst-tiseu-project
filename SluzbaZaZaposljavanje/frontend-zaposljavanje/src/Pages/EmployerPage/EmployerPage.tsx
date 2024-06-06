import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { toast } from 'react-toastify';
import { JobApplication } from '../../Interfaces/JobApplication';


const EmployerPage = () => {

    const [jobApplications, setJobApplications] = useState<JobApplication[]>([]);
    const userId = localStorage.getItem("userId")

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8012/job/employerPage/' + userId);
            if (response.ok) {
                const jobApplications: JobApplication[] = await response.json();
                setJobApplications(jobApplications);
            } else {
                throw new Error('Failed to fetch job applications');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    const handleAcceptClick = () => {}

    const declineJobApplication = async (id: string) => {
        try {
            const response = await fetch(`http://localhost:8012/job/jobApplications/${id}`, {
                method: 'DELETE',
            });
            if (response.ok) {
                setJobApplications(prevState => prevState.filter(jobApplication => jobApplication.id !== id));
                toast.success('Job application declined successfully');
            } else {
                throw new Error('Failed to decline job application');
            }
        } catch (error) {
            console.error('Error:', error);
            toast.error('Failed to decline job application');
        }
    };

    return (
        <div className="d-flex flex-wrap justify-content-start">
          {jobApplications.map((jobApplication: JobApplication, index: number) => (
            <div key={index} className="p-2" style={{ flex: '0 0 calc(50% - 2rem)' }}>
              <div className="card h-100">
                <div className="card-body">
                  <h5 className="card-title">Job Listing ID: {jobApplication.jobListingId}</h5>
                  <h6 className="card-subtitle mb-2 text-muted">Employee ID: {jobApplication.employeeId}</h6>
                  <p className="card-text">Diploma: {jobApplication.diploma}</p>
                  <button className="btn btn-success me-2" onClick={handleAcceptClick}>Accept</button>
                  <button className="btn btn-danger" onClick={() => declineJobApplication(jobApplication.id)}>Decline</button>
                </div>
              </div>
            </div>
          ))}
        </div>
      );
      
      
}

export default EmployerPage;