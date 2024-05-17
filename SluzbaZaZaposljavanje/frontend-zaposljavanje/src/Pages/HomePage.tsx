import React from 'react';
import "./LoginPage.css"
import { Button } from 'react-bootstrap';
import JobsPage from './JobsPage';

function HomePage() {
  return (
    <div className="d-flex justify-content-center align-items-center vh-100">
      <Button className="btn-lg"> <a href='/job_list'></a>Jobs Page</Button>
    </div>
  );
}

export default HomePage;
