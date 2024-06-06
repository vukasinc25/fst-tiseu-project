import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import LoginPage from "../Pages/Login/LoginPage";
import HomePage from "../Pages/Home/HomePage";
import JobsPage from "../Pages/Jobs/JobsPage";
import JobDetailPage from "../Pages/JobDetails/JobDetailPage";
import EmployerPage from "../Pages/EmployerPage/EmployerPage";

export const router = createBrowserRouter ([
    {
        path: "/",
        element: <App/>,
        children: [
            {path: "", element: <HomePage/>},
            {path: "login", element: <LoginPage/>},
            {path: "job_list", element: <JobsPage/>},
            {path: "job_info/:jobId", element: <JobDetailPage/>},
            {path: "employer_page", element: <EmployerPage/>}
        ],
    }
])
