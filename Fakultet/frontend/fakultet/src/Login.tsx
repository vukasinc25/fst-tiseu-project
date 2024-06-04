import './Login.css';
import { useState } from "react";
import customFetch from './intersceptor/interceptor';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        console.log(`Username: ${username}, Password: ${password}`);

        try {
            // Sending credentials to the backend
            const response1 = await fetch('http://localhost:8000/users/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username: username,
                    password: password,
                }),
            });

            const data1 = await response1.json();
            console.log('Response:', data1);

            let accessToken = data1.access_token;
            console.log('Access Token:', accessToken);

            if (accessToken) {
                localStorage.setItem('accessToken', accessToken); // Save token to localStorage
                console.log('Access Token:', localStorage.getItem("accessToken"));
            }

            // Fetching study programs using customFetch
            const response2 = await customFetch('http://localhost:8001/fakultet/studyPrograms', {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            console.log('Response:', response2);

        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Username:
            </label>
            <input type="text" value={username} onChange={e => setUsername(e.target.value)} />
            <label>
                Password:
            </label>
            <input type="password" value={password} onChange={e => setPassword(e.target.value)} />
            <input type="submit" value="Submit" />
        </form>
    );
}

export default Login;