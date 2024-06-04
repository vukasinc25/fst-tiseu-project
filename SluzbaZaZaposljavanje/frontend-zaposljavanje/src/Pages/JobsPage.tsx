import { useEffect, useState } from "react";
import { useNavigate } from "react-router";

interface Job {
    _id: string;
    jobTitle: string;
    jobDescription: string;
    requirements: string;
}

function JobsPage() {
    const [jobs, setJobs] = useState<Job[]>([]);

    useEffect(() => {
        fetchData();
    }, []);

    let navigate = useNavigate();

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8012/getJobListings');
            if (response.ok) {
                const jobsData: Job[] = await response.json();
                setJobs(jobsData);
                console.log(jobsData)
            } else {
                throw new Error('Failed to fetch jobs');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };
  
    return (
        <>
            <div className="row">
                {jobs.map((job: Job, index: number) => (
                    <div key={index} className="col-sm-4">
                        <div className="card">
                            <div className="card-body">
                                <h5 className="card-title">{job.jobTitle}</h5>
                                <h2 className="card-text">{job.jobDescription}</h2>
                                <button className="btn btn-primary" onClick={() => navigate("/job_info/" + job._id)}>More Information</button>
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </>
    );
}

export default JobsPage;