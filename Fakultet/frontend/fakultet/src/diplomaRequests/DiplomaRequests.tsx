import { useEffect, useState } from "react";
import "./DiplomaRequests.css"
import customFetch from "../intersceptor/interceptor";
const DiplomaRequests = () => {
    const [diplomaRequests, setDiplomaRequests] = useState([]);

    useEffect(() => {
      fetchDiplomaRequests();
    }, []);
  
    const fetchDiplomaRequests = async () => {
      try {
        const response = await customFetch('http://localhost:8001/fakultet/diplomaRequestsInPendingState');
        setDiplomaRequests(response);
      } catch (error) {
        console.error('Error fetching diploma requests:', error);
      }
    };

    const handleApprove = async (requestId: any) => {
        try {
            const response = await customFetch(`http://localhost:8001/fakultet/decideDiplomaReques/${requestId}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body:JSON.stringify({isApproved:true}),
            });
            console.log('Diploma request sent successfully:', response);
            alert("Request Approved")
            fetchDiplomaRequests();
        } catch (error) {
            console.error('Failed to send diploma request:', error);
        }
      };
    
      const handleDecline = async (requestId: any) => {
        try {
            const response = await customFetch(`http://localhost:8001/fakultet/decideDiplomaReques/${requestId}`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body:JSON.stringify({isApproved:false}),
            });
            console.log('Diploma request sent successfully:', response);
            alert("Request Approved")
            fetchDiplomaRequests();
        } catch (error) {
            console.error('Failed to send diploma request:', error);
        }
      };
  
    return (
      <div className="diploma-requests">
        <h2>Diploma Requests</h2>
        <table>
          <thead>
            <tr>
              <th>Student Name</th>
              <th>Request Date</th>
              <th></th>
              <th></th>
            </tr>
          </thead>
          <tbody>
            {diplomaRequests?.map((request: any) => (
              <tr key={request._id}>
                <td>{request.userId}</td>
                <td>{request.issueDate}</td>
                <td><button className="approve-btn" onClick={() => handleApprove(request._id)}>Approve</button></td>
                <td><button className="decline-btn" onClick={() => handleDecline(request._id)}>Decline</button></td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };
 
export default DiplomaRequests;