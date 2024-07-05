import { useState } from 'react';
import './CreateDepartments.css'
import customFetch from '../intersceptor/interceptor';

const CreateDepartments = () => {
    const [name, setName] = useState('');

    const clearForm = () => {
        setName("");
    };

    const handleSubmit = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        console.log('Department name:', name);
        try {
            await customFetch('http://localhost:8001/fakultet/department', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({"name":name}),
            });

            clearForm();
            alert("Department created")
        } catch (error) {
            console.error('Failed to create competition:', error);
        }
    };

    return (
        <div className="create-department-container">
            <h2>Create Department</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="name">Department Name:</label>
                    <input
                        type="text"
                        id="name"
                        name="name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Create</button>
            </form>
        </div>
    );
}
 
export default CreateDepartments;