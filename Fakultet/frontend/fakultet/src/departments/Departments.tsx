import { useEffect, useState } from "react";
import "./Departments.css"
import customFetch from "../intersceptor/interceptor";
const Departments = () => {
    const [departments, setDepartments] = useState([]);

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

    return (  
    <div className="App1">
      <h1>Departments</h1>
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