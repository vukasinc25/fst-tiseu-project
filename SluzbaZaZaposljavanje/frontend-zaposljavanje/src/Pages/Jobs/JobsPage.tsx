import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import { JobListing } from "../../Interfaces/JobListing";
import './JobsPage.css'; // Import the CSS file

function JobsPage() {
    const [jobs, setJobs] = useState<JobListing[]>([]);

    useEffect(() => {
        fetchData();
    }, []);

    let navigate = useNavigate();

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8012/job/getJobListings');
            if (response.ok) {
                const jobsData: JobListing[] = await response.json();
                setJobs(jobsData);
            } else {
                throw new Error('Failed to fetch jobs');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };
  
    return (
        <div className="jobs-page-container">
            <div className="row justify-content-center flex-grow-1">
                {jobs.map((job: JobListing, index: number) => (
                    <div key={index} className="col-sm-4 mb-4">
                        <div className="card h-100">
                            <div className="card-body">
                                <h5 className="card-title fw-bold">{job.jobTitle}</h5>
                                <p className="card-text">{job.jobDescription}</p>
                                <button className="btn btn-primary mt-3" onClick={() => navigate("/job/info/" + job._id)}>More Information</button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default JobsPage;
