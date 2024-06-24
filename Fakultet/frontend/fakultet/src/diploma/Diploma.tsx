import "./Diploma.css"
import { useEffect, useState } from "react";
import customFetch from "../intersceptor/interceptor";
const Diploma = () => {
    const [diploma, setDiploma] = useState<any>(null);

    useEffect(() => {
        const fetchDiploma = async () => {
            try {
                const data = await customFetch(`http://localhost:8001/fakultet/user/diplomaByUserId`);
                setDiploma(data);
                console.log("Data: ", data);
            } catch (error) {
                console.error('Failed to fetch diploma:', error);
            }
        };

        fetchDiploma();
    }, []);
    
    return (
        <div className="diploma-details">
            <h2>Diploma Details</h2>
            {diploma && (
                <div>
                    <p><strong>User ID:</strong> {diploma.userId}</p>
                    <p><strong>Issue Date:</strong> {new Date(diploma.issueDate).toLocaleDateString()}</p>
                    <p><strong>Average Grade:</strong> {diploma.averageGrade}</p>
                </div>
            )}
        </div>
    );
}
 
export default Diploma;