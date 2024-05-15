import React, { useState } from 'react';
import "./LoginPage.css"
import { Button } from 'react-bootstrap';
import JobsPage from './JobsPage';

function HomePage() {
  const [showJobsPage, setShowJobsPage] = useState(false);

  const handleClick = () => {
    setShowJobsPage(true);
  };

  return (
    <>
      {showJobsPage ? <JobsPage /> : (
        <Button >Jobs Page</Button>
      )}
    </>
  );
}

export default HomePage;
