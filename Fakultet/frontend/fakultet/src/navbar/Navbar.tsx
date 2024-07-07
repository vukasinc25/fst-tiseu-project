// import { Link } from "react-router-dom";
// import "./Navbar.css"
// import useRoles from "../role-base/userValidation";
// const Navbar = () => {
//   const { hasRole } = useRoles();
//   console.log("Is user a STUDENT?", hasRole("STUDENT"));
//     return (
//         <nav className="navbar">
//           <ul className="nav-list">
//             <div className="nav-group">
//               <li className="nav-item"><Link to="/competitions">Home</Link></li>
//             </div>
//             <div className="nav-group">
//               <li className="nav-item"><Link to="/">Login</Link></li>
//               <li className="nav-item"><Link to="/studyPrograms">Study Programs</Link></li>
//               <li className="nav-item"><Link to="/departments">Departments</Link></li>
//              {hasRole("STUDENT") && <li className="nav-item"><Link to="/diploma">Diploma</Link></li>}
//               <li className="nav-item"><Link to="/diplomaRequests">Diploma Requests</Link></li>
//             </div>
//           </ul>
//         </nav>
//       );
// }
 
// export default Navbar;
import { Link } from "react-router-dom";
import "./Navbar.css";
import { useAuth0 } from "@auth0/auth0-react";
import useRoles from "../role-base/userValidation";

const Navbar = () => {
  const { loginWithRedirect, logout, user, isLoading } = useAuth0();
  const { hasRole } = useRoles();
  if (!isLoading) {
    console.log("User: ",user);
  }

  return (
    <nav className="navbar">
      <ul className="nav-list">
        <div className="nav-group">
          <li className="nav-item">
            <Link to="/competitions">Home</Link>
          </li>
        </div>
        <div className="nav-group">
          {/* <li className="nav-item">
            <Link to="/">Login</Link>
          </li> */}
          {/* {this.state.value == 'news'? <Text>data</Text>: null } */}
          {!isLoading && !user ? (
            <li className="nav-item">
              <button className="nav-item" onClick={() => loginWithRedirect()}>
                Log In2
              </button>
            </li>
          ) : (
            <li className="nav-item">
              <button
                className="nav-item"
                onClick={() =>
                  logout({ logoutParams: { returnTo: window.location.origin } })
                }
              >
                Log out
              </button>
            </li>
          )}

          <li className="nav-item">
            <Link to="/studyPrograms">Study Programs</Link>
          </li>
          <li className="nav-item">
            <Link to="/departments">Departments</Link>
          </li>
          {user?.user_metadata?.roles.includes("STUDENT") && <li className="nav-item">
            <Link to="/diploma">Diploma</Link>
          </li>}
          {user?.user_metadata?.roles.includes("ADMIN") && <li className="nav-item"><Link to="/diplomaRequests">Diploma Requests</Link></li>}
        </div>
      </ul>
    </nav>
  );
};

export default Navbar;
