// import * as jwt from 'jsonwebtoken';
import { useHistory } from 'react-router-dom';
import './Login.css';
import { useState } from "react";

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const history = useHistory();

    function getRoleFromJWT(token: string): string[] {
        const parts = token.split('.');
      
        if (parts.length !== 3) {
          throw new Error('Invalid JWT token format');
        }
      
        const payload = JSON.parse(window.atob(parts[1]));
      
        if (!payload.roles) {
          throw new Error('Role not found in JWT token');
        }
      
        return payload.roles;
      }
    
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

            const roles = getRoleFromJWT(accessToken);
            console.log('Roles:', roles);

            if (accessToken) {
                sessionStorage.setItem('userRoles', JSON.stringify(roles)); // Save user roles to session storage
                console.log('User roles:', JSON.parse(sessionStorage.getItem('userRoles') || '[]'));
                localStorage.setItem('accessToken', accessToken); // Save token to localStorage
                console.log('Access Token:', localStorage.getItem("accessToken"));
            }


            history.push("/competitions");


        } catch (error) {
            console.error('Error:', error);
        }
    };

    return (
        <div>
            <h1>Login</h1>
            <form className='form1' onSubmit={handleSubmit}>
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
        </div>
    );
}

export default Login;