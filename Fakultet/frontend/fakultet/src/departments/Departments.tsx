import { useEffect, useState } from "react";
import "./Departments.css"
import customFetch from "../intersceptor/interceptor";
import { useHistory } from "react-router-dom";
import useRoles from "../role-base/userValidation";
const Departments = () => {
    const [departments, setDepartments] = useState([]);
    const history = useHistory();
    const { hasRole } = useRoles();
    
    useEffect(() => {
        fetchCompetitions();
      }, []);
    
    const fetchCompetitions = async () => {
      try {
        const data = await customFetch('http://localhost:8001/fakultet/departments');
        setDepartments(data);
        console.log("Data: ",data)
      } catch (error) {
        alert("Departments mush be created!!!!!")
        console.log('Failed to fetch departments:', error);
      }
    };

    const redirectToCreateDepartment = () => {
      history.push('/department');
    };

    return (  
    <div className="App1">
      <h1>Departments</h1>
      {hasRole("ADMIN") && <div>
        <button className="create-department-button" onClick={redirectToCreateDepartment}>
            Create Department
        </button>
      </div>}
      <div className="departments-grid">
        {departments?.map((department: any) => (
          <div key={department._id} className="department-card">
            <h2 className="department-name">{department.name}</h2>
            {department.staff.map((worker:any) => (
                <li key={worker._id} className="staff-list-item">
                     <strong>Employee: {worker.username}</strong>
                </li>
            ))}
          </div>
        ))}
      </div>
    </div>
    );
}
 
export default Departments;