import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { toast } from 'react-toastify';

interface JobApplications {

}

const EmployerPage = () => {

    const [jobApplications, setJobApplications] = useState<JobApplications[]>([]);

    const fetchData = async () => {
        try {
            const response = await fetch('http://localhost:8012/employerPage');
            if (response.ok) {
                const jobApplications: JobApplications[] = await response.json();
                setJobApplications(jobApplications);
            } else {
                throw new Error('Failed to fetch job applications');
            }
        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <>
            
        </>
    );
}

export default EmployerPage;