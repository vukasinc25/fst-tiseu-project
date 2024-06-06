import './Competition.css';
import { useState } from "react";
import customFetch from "../intersceptor/interceptor";

const Competition = () => {
    
    const [formData, setFormData] = useState({
        programName: "",
        admissionRequirements: "",
        examDate: "",
        examFormat: "",
        examMaterials: "",
        applicationDeadlines: "",
        applicationDocuments: "",
        applicationMethod: "",
        applicationContact: "",
        tuitionFees: "",
        contactInformation: ""
    });

    const clearForm = () => {
        setFormData({
            programName: "",
            admissionRequirements: "",
            examDate: "",
            examFormat: "",
            examMaterials: "",
            applicationDeadlines: "",
            applicationDocuments: "",
            applicationMethod: "",
            applicationContact: "",
            tuitionFees: "",
            contactInformation: ""
        });
    };
    
    const handleSubmit = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        try {
            await customFetch('http://localhost:8001/fakultet/createCompetition', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(formData),
            });
            setFormData({
                programName: formData.programName,
                admissionRequirements: formData.admissionRequirements,
                examDate: formData.examDate,
                examFormat: formData.examFormat,
                examMaterials: formData.examMaterials,
                applicationDeadlines: formData.applicationDeadlines,
                applicationDocuments: formData.applicationDocuments,
                applicationMethod: formData.applicationMethod,
                applicationContact: formData.applicationContact,
                tuitionFees: formData.tuitionFees,
                contactInformation: formData.contactInformation
            });

            clearForm();
        } catch (error) {
            console.error('Failed to create competition:', error);
        }
        clearForm();
    };
    
    
    const handleChange = (event: { target: { name: any; value: any; }; }) => {
        const { name, value } = event.target;
        setFormData({
            ...formData,
            [name]: value
        });
    };
    
    const isFormEmpty = () => {
        const isNotEmpty = Object.values(formData).every(value => value !== "");
        console.log("Form is not empty:", isNotEmpty);
        return !isNotEmpty;
    };
    
    
    return (
        <div className="App2">
            <h2>Create Competition</h2>
            <form className='form2' onSubmit={handleSubmit}>
                <label>
                    Program Name:
                    <input
                        type="text"
                        name="programName"
                        value={formData.programName}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Admission Requirements:
                    <input
                        type="text"
                        name="admissionRequirements"
                        value={formData.admissionRequirements}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Exam Date:
                    <input
                        type="date"
                        name="examDate"
                        value={formData.examDate}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Exam Format:
                    <input
                        type="text"
                        name="examFormat"
                        value={formData.examFormat}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Exam Materials:
                    <input
                        type="text"
                        name="examMaterials"
                        value={formData.examMaterials}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Application Deadlines:
                    <input
                        type="date"
                        name="applicationDeadlines"
                        value={formData.applicationDeadlines}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Application Documents:
                    <input
                        type="text"
                        name="applicationDocuments"
                        value={formData.applicationDocuments}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Application Method:
                    <input
                        type="text"
                        name="applicationMethod"
                        value={formData.applicationMethod}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Application Contact:
                    <input
                        type="text"
                        name="applicationContact"
                        value={formData.applicationContact}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Tuition Fees:
                    <input
                        type="text"
                        name="tuitionFees"
                        value={formData.tuitionFees}
                        onChange={handleChange}
                    />
                </label>
                <label>
                    Contact Information:
                    <input
                        type="text"
                        name="contactInformation"
                        value={formData.contactInformation}
                        onChange={handleChange}
                    />
                </label>
                <button type="submit" disabled={isFormEmpty()}>Create Competition</button>
            </form>
        </div>
    );
}

export default Competition;
