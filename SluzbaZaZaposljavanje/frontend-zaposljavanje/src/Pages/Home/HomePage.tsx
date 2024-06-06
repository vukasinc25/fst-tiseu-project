import React from 'react';
import { Button } from 'react-bootstrap';
import { useNavigate } from 'react-router';
import backgroundImage from "../../Components/job_interview1.jpg";

function HomePage() {
  let navigate = useNavigate();

  const handleClick = () => {
    return navigate("/job/list")
  };

  return (
    <div 
      className="d-flex flex-column justify-content-center align-items-center"
      style={{ 
        backgroundImage: `linear-gradient(rgba(0,0,0,0.7), rgba(0,0,0,0.7)), url(${backgroundImage})`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        minHeight: '100vh'
      }}
    >
      <h1 className="text-white mb-4">Find Your Next Job</h1>
      <Button 
        className="btn-lg btn-primary mb-2" 
        onClick={handleClick}
        style={{ 
          fontSize: '1.5rem', 
          padding: '15px 30px'
        }}
      >
        Go to Job Listings
      </Button>
    </div>
  );
}

export default HomePage;
