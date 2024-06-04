import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import './JobDetailPage.css';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

interface Job {
    _id: string;
    jobTitle: string;
    jobDescription: string;
    requirements: string;
}

const JobDetailPage = () => {
    const { jobId } = useParams();
    const [job, setJob] = useState<Job | null>(null);

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
                console.log(jobsData)
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
                throw new Error('Failed to submit job application');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to submit job application');
        }
    };

    return (
        <div className="container">
            <h1 className="heading">Job Details</h1>
            {job ? (
                <div className="job-details">
                    <i className="job-title"><strong>Job Name:</strong> {job.jobTitle}</i>
                    <p className="job-description"><strong>Description:</strong> {job.jobDescription}</p>
                    <p className="requirements"><strong>Requirements:</strong> {job.requirements}</p>
                    <button className="btn btn-primary" onClick={() => createJobApplication()}>Apply</button>
                </div>
            ) : (
                <p>Job Does Not Exist</p>
            )}
        </div>
    );
};

export default JobDetailPage;
