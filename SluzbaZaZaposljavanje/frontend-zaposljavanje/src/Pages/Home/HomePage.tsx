import React from 'react';
import { Button } from 'react-bootstrap';
import { useNavigate } from 'react-router';

function HomePage() {
  let navigate = useNavigate();

  const handleClick = () => {
    return navigate("/job_list")
  };

  return (
    <div className="d-flex justify-content-center align-items-center vh-100">
      <Button className="btn-lg" onClick={handleClick}>Jobs Page</Button>
    </div>
  );
}

export default HomePage;
