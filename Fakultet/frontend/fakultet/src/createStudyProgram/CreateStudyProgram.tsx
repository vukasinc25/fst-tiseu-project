import { useEffect, useState } from "react";
import customFetch from "../intersceptor/interceptor";
import "./CreateStudyProgram.css"
import useRoles from "../role-base/userValidation";
const CreateStudyProgram = () => {
    const [formData, setFormData] = useState({
        name: "",
        studyLevel: "",
        duration: "",
        objectives: "",
        programStructure: "",
        internship: false,
        graduationRequirements: "",
        accreditation: false,
        contactPersonID: "",
        developmentPlan: "",
        departmentID: ""
    });

    const [departments, setDepartments] = useState([]);
    const [fetchError, setFetchError] = useState(false);
    const {hasRole} = useRoles();

    useEffect(() => {
        fetchDepartments();
    }, []);

    const fetchDepartments = async () => {
        try {
            const data = await customFetch('http://localhost:8001/fakultet/departments');
            setDepartments(data);
            console.log("Departments: ",data)
            setFetchError(false)
        } catch (error) {
            console.error('Failed to fetch departments:', error);
            setFetchError(true)
        }
    };

    const clearForm = () => {
        setFormData({
            name: "",
            studyLevel: "",
            duration: "",
            objectives: "",
            programStructure: "",
            internship: false,
            graduationRequirements: "",
            accreditation: false,
            contactPersonID: "",
            developmentPlan: "",
            departmentID: ""
        });
    };
    
    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            await customFetch('http://localhost:8001/fakultet/studyProgram', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });

            clearForm();
        } catch (error) {
            console.error('Failed to create study program:', error);
        }
    };
    
    const handleChange = (event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = event.target;

        if (event.target instanceof HTMLInputElement && event.target.type === 'checkbox') {
            setFormData({
                ...formData,
                [name]: event.target.checked
            });
        } else {
            setFormData({
                ...formData,
                [name]: value
            });
        }
    };
    
    const isFormEmpty = () => {
        const isNotEmpty = Object.values(formData).every(value => value !== "");
        return !isNotEmpty;
    };
    
    return (
        <div className="App2">
            <h2>Create Study Program</h2>
            <form className='form2' onSubmit={handleSubmit}>
                <label>
                    Name:
                    <input
                        type="text"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Study Level:
                    <input
                        type="text"
                        name="studyLevel"
                        value={formData.studyLevel}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Duration:
                    <input
                        type="text"
                        name="duration"
                        value={formData.duration}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Objectives:
                    <input
                        type="text"
                        name="objectives"
                        value={formData.objectives}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Program Structure:
                    <input
                        type="text"
                        name="programStructure"
                        value={formData.programStructure}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Internship:
                    <input
                        type="checkbox"
                        name="internship"
                        checked={formData.internship}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Graduation Requirements:
                    <input
                        type="text"
                        name="graduationRequirements"
                        value={formData.graduationRequirements}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Accreditation:
                    <input
                        type="checkbox"
                        name="accreditation"
                        checked={formData.accreditation}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Contact Person ID:
                    <input
                        type="text"
                        name="contactPersonID"
                        value={formData.contactPersonID}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Development Plan:
                    <input
                        type="text"
                        name="developmentPlan"
                        value={formData.developmentPlan}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Department ID:
                    <select
                        name="departmentID"
                        value={formData.departmentID}
                        onChange={handleChange}
                    >
                        <option value=""></option>
                        {departments?.map((department: any) => (
                            <option key={department._id} value={department._id}>
                                {department.name}
                            </option>
                        ))}
                    </select>
                </label>
                <br/>
                <p>{fetchError}</p>
                {/* {!fetchError && <p>Failed to load departments. Please try again later.</p>} */}
                <p>{fetchError ? "Failed to load departments. Please try again later." : ""}</p>
                <br/>
                {hasRole("ADMIN") && <button type="submit" disabled={isFormEmpty()}>Create Study Program</button>}
            </form>
        </div>
    );
}
 
export default CreateStudyProgram;