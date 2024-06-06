import React, { useState } from "react";
import { toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { JobListing } from "../../Interfaces/JobListing";
 
const CreateJobListing: React.FC = () => {
  const [jobListing, setJobListing] = useState<JobListing>({
    employerId: "",
    companyName: "",
    cityName: "",
    jobTitle: "",
    jobDescription: "",
    requirements: ""
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setJobListing({ ...jobListing, [name]: value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    jobListing.employerId = localStorage.getItem("userId") ?? "";
    console.log(JSON.stringify(jobListing));
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8012/job/createJobListing', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(jobListing)
      });

      if (response.ok) {
        toast.success("Job listing created successfully!");
        setJobListing({
          employerId: "",
          companyName: "",
          cityName: "",
          jobTitle: "",
          jobDescription: "",
          requirements: ""
        });
      } else {
        throw new Error("Failed to create job listing");
      }
    } catch (error) {
      console.error("Error:", error);
      toast.error("Failed to create job listing");
    }
  };

  return (
    <div className="d-flex justify-content-center align-items-center min-vh-100 bg-light">
      <div className="card p-4 shadow-sm" style={{ width: "100%", maxWidth: "600px" }}>
        <h2 className="card-title text-center mb-4">Create Job Listing</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label htmlFor="companyName" className="form-label">Company Name</label>
            <input
              type="text"
              className="form-control"
              id="companyName"
              name="companyName"
              value={jobListing.companyName}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-3">
            <label htmlFor="cityName" className="form-label">City Name</label>
            <input
              type="text"
              className="form-control"
              id="cityName"
              name="cityName"
              value={jobListing.cityName}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-3">
            <label htmlFor="jobTitle" className="form-label">Job Title</label>
            <input
              type="text"
              className="form-control"
              id="jobTitle"
              name="jobTitle"
              value={jobListing.jobTitle}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-3">
            <label htmlFor="jobDescription" className="form-label">Job Description</label>
            <textarea
              className="form-control"
              id="jobDescription"
              name="jobDescription"
              value={jobListing.jobDescription}
              onChange={handleChange}
              required
            />
          </div>
          <div className="mb-3">
            <label htmlFor="requirements" className="form-label">Requirements</label>
            <textarea
              className="form-control"
              id="requirements"
              name="requirements"
              value={jobListing.requirements}
              onChange={handleChange}
              required
            />
          </div>
          <button type="submit" className="btn btn-primary w-100">Create Job Listing</button>
        </form>
      </div>
    </div>
  );
}


export default CreateJobListing;
