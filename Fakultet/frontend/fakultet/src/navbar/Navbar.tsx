import { Link } from "react-router-dom";
import "./Navbar.css"
const Navbar = () => {
    return (
        <nav className="navbar">
          <ul className="nav-list">
            <div className="nav-group">
              <li className="nav-item"><Link to="/competitions">Home</Link></li>
            </div>
            <div className="nav-group">
              <li className="nav-item"><Link to="/">Login</Link></li>
              <li className="nav-item"><Link to="/studyPrograms">Study Programs</Link></li>
              <li className="nav-item"><Link to="/departments">Departments</Link></li>
            </div>
          </ul>
        </nav>
      );
}
 
export default Navbar;