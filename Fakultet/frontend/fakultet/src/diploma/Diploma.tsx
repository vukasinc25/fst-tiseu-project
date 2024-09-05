import "./Diploma.css"
import { useEffect, useState } from "react";
import customFetch from "../intersceptor/interceptor";
import { isParameter } from "typescript";
import { useAuth0 } from "@auth0/auth0-react";
import {jsPDF} from 'jspdf'
const Diploma = () => {
    const [diploma, setDiploma] = useState<any>(null);
    const [diplomaRequests, setDiplomaRequests] = useState<any[]>([]);
    const [isAnyApproved, setIsAnyApproved] = useState<boolean>(false);
    const [inPending, setInPending] = useState<boolean>(false);
    const [isRequestSent, setIsRequestSent] = useState<boolean>(false);
    const { loginWithRedirect, logout, user, isLoading } = useAuth0();

    useEffect(() => {
        const fetchDiploma = async () => {
            const userId = user?.sub?.split('|')[1];
            try {
                const data = await customFetch(`http://localhost:8001/fakultet/user/diplomaByUserId/${userId}`);
                setDiploma(data);
                console.log("Data: ", data);
            } catch (error) {
                console.error('Failed to fetch diploma:', error);
            }
        };

        const fetchUserDiplomaRequests = async () => {
            const userId = user?.sub?.split('|')[1];
            try {
                const data = await customFetch(`http://localhost:8001/fakultet/getDiplomaRequestsForUserId/${userId}`);
                setDiplomaRequests(data);
                setIsAnyApproved(data.some((request: any) => request.IsApproved === true));
                setInPending(data.some((request: any) => request.InPending === true));
                console.log("isAnyApproved:", isAnyApproved);
                console.log("inPending:", inPending);
                console.log("Data: ", data);
            } catch (error) {
                console.error('Failed to fetch diploma requests:', error);
            }
        };

        fetchUserDiplomaRequests();
        fetchDiploma();
    }, []);

    useEffect(() => {
        console.log("isAnyApproved:", isAnyApproved);
        console.log("inPending:", inPending);
    }, [isAnyApproved, inPending]);

    const sendDiplomaRequest = async () => {
        const userId = user?.sub?.split('|')[1];
        const userName = user?.name;
        console.log("User: ", user)
        try {
            const response = await customFetch(`http://localhost:8001/fakultet/diplomaRequest/${userId}/${userName}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body:"",
            });
            setDiplomaRequests(response)
            setIsRequestSent(true)
            console.log('Diploma request sent successfully:', response);
            alert("Request sent successfully")
        } catch (error) {
            console.error('Failed to send diploma request:', error);
        }
    };

    const handleGeneratePDF = () => {
        if (diploma) {
            const doc = new jsPDF();
            doc.setFontSize(16);
            doc.text("Diploma Details", 20, 20);
            
            doc.setFontSize(12);
            doc.text(`User ID: ${diploma.userId}`, 20, 40);
            doc.text(`User Name: ${diploma.userName}`, 20, 50);
            doc.text(`Issue Date: ${new Date(diploma.issueDate).toLocaleDateString()}`, 20, 60);
            doc.text(`Average Grade: ${diploma.averageGrade}`, 20, 70);
            
            // Save the PDF
            doc.save(`Diploma_${diploma.userName}.pdf`);
        } else {
            alert("Diploma data is not available.");
        }
    };

    return (
        <div>
            {(inPending || isAnyApproved || isRequestSent) ? null : (
            <div className="diploma-request-container">
                <p>To get youre diploma you need first to send request to the ADMIN</p>
                <button onClick={sendDiplomaRequest}>Send diploma request</button>
            </div>)}
            {diploma && (<div className="diploma-details">
                <h2>Diploma Details</h2>
                <div>
                    <p><strong>User ID:</strong> {diploma.userId}</p>
                    <p><strong>User Name:</strong> {diploma.userName}</p>
                    <p><strong>Issue Date:</strong> {new Date(diploma.issueDate).toLocaleDateString()}</p>
                    <p><strong>Average Grade:</strong> {diploma.averageGrade}</p>
                </div>
                <button onClick={handleGeneratePDF}>Generate PDF</button>
            </div>)}
        </div>
    );
}
 
export default Diploma;