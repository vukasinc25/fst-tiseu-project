import React from 'react';
import "./LoginPage.css"
import { Button } from 'react-bootstrap';
import JobsPage from './JobsPage';
import { useNavigate } from 'react-router';
import { Link } from 'react-router-dom';

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
