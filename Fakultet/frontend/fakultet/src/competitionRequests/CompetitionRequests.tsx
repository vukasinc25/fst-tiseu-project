import { useEffect, useState } from 'react';
import './CompetitionRequests.css'
import { useParams } from 'react-router-dom';
import customFetch from '../intersceptor/interceptor';
import { RouteParams } from '../intefaces/routeParams';
const CompetitionRequests = () => {
    const { id } = useParams<RouteParams>();
    const [requests, setRequests] = useState([]);

    useEffect(() => {
        fetchCompetitionRequests();
    }, []);

    const fetchCompetitionRequests = async () => {
    try {
        const data = await customFetch(`http://localhost:8001/fakultet/user/getRegistrationsToCompetition/${id}`);
        setRequests(data);
        console.log("Data: ",data)
    } catch (error) {
        console.error('Failed to fetch competition request:', error);
    }
    };

    return ( 
        <div className="requests-container">
            <h1>Competition Requests</h1>
            <ul className="requests-list">
                {requests?.map((request: any) => (
                <li key={request._id} className="request-item">
                    <p>UserId: {request.userID}</p>
                </li>
                ))}
            </ul>
        </div>
    );
}
 
export default CompetitionRequests;