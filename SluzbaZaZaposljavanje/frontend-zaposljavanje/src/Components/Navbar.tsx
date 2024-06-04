import * as React from 'react';
import "bootstrap/dist/css/bootstrap.min.css";
import { Navbar } from "react-bootstrap";
import { Link, useNavigate } from 'react-router-dom';

export default function Navigationbar() {
  let navigate = useNavigate();

  const logout = () => {
    localStorage.removeItem('jwtToken');
  };

  const isLoggedIn = localStorage.getItem("jwtToken")

    return ( 
        <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
        <a id='navbar' className="navbar-brand" href="#">Sluzba Za Zaposljavanje</a>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
      
        <div className="collapse navbar-collapse" id="navbarSupportedContent">
          <ul className="navbar-nav mr-auto">
            <li className="nav-item active">
              <Link className="nav-link" to="/">Home</Link>
            </li>
            <li className="nav-item active">
              <Link className="nav-link" to="/job_list">Jobs</Link>
            </li>
            <div>
              {isLoggedIn ? (<li className="nav-item">
                <Link className="nav-link" to="/login" onClick={() => logout()}>Logout</Link></li>) 
              : 
              (<li className="nav-item">
              <Link className="nav-link" to="/login">Login</Link>
              </li>)}
            </div>
            {/* <li className="nav-item dropdown">
              <a className="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                Dropdown
              </a>
              <div className="dropdown-menu" aria-labelledby="navbarDropdown">
                <a className="dropdown-item" href="#">Action</a>
                <a className="dropdown-item" href="#">Another action</a>
                <div className="dropdown-divider"></div>
                <a className="dropdown-item" href="#">Something else here</a>
              </div>
            </li> */}
            {/* <li className="nav-item">
              <a className="nav-link disabled" href="#">Disabled</a>
            </li> */}
          </ul>
        </div>
      </nav>
     );
}