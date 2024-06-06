import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import './JobDetailPage.css';
import { toast } from 'react-toastify';

interface Job {
    _id: string;
    employerId: string;
    jobTitle: string;
    jobDescription: string;
    requirements: string;
}

const JobDetailPage = () => {
    const { jobId } = useParams();
    const [job, setJob] = useState<Job | null>(null);
    const [isVisible, setIsVisible] = useState(true);

    useEffect(() => {
        fetchData();
    }, []);
    
    const jwtToken = localStorage.getItem("jwtToken")
    const employeeId = jwtToken;
    const jobListingId = jobId;

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8012/getJobInfo/' + jobId);
            if (response.ok) {
                const jobsData: Job = await response.json();
                setJob(jobsData);
            } else {
                throw new Error('Failed to fetch job info');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };
    
    const createJobApplication = async () => {
        try {
            const response = await fetch('http://localhost:8012/applyForJob', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ jobListingId, employeeId }),
            });
            // console.log("jobId:" + jobId) 
            // console.log("jwtToken:" + jwtToken)
            if (response.ok) {
                console.log("toast where")
                toast.success('Job application submitted successfully!', {position: "top-right"});
            } else {
                toast.error("Failed to aplpy for the job", {position: "top-right"})
                throw new Error('Failed to submit job application');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to submit job application');
        }
    };

    const handleButtonClick = () => {
        createJobApplication();
        setIsVisible(false);
    }

    return (
        <div className="container">
            <h1 className="heading">Job Details</h1>
            {job ? (
                <div className="job-details">
                    <i className="job-title"><strong>Job Name:</strong> {job.jobTitle}</i>
                    <p className="job-description"><strong>Description:</strong> {job.jobDescription}</p>
                    <p className="requirements"><strong>Requirements:</strong> {job.requirements}</p>
                    {isVisible && (
                        <button className="btn btn-primary" onClick={handleButtonClick}>Apply</button>
                    )}
                </div>
            ) : (
                <p>Job Does Not Exist</p>
            )}
        </div>
    );
};

export default JobDetailPage;
