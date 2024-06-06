import React from 'react';
import './Footer.css';

function Footer() {
  return (
    <footer className="footer">
      <div className="container">
        <div className="row">
          <div className="col-lg-6">
            <h5>Contact Us</h5>
            <p>Email: support@jobs.gov</p>
            <p>Phone: 123-456-7890</p>
          </div>
          <div className="col-lg-6">
            <h5>Follow Us</h5>
            <div className="social-icons">
              {/* Add your social media icons here */}
              <a href="#"><i className="fab fa-facebook"></i></a>
              <a href="#"><i className="fab fa-twitter"></i></a>
              <a href="#"><i className="fab fa-instagram"></i></a>
            </div>
          </div>
        </div>
      </div>
      <div className="text-center py-2">
        <p>&copy; 2024 Your Website. All Rights Reserved.</p>
      </div>
    </footer>
  );
}

export default Footer;
