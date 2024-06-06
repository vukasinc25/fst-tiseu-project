import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import LoginPage from "../Pages/Login/LoginPage";
import HomePage from "../Pages/Home/HomePage";
import JobsPage from "../Pages/Jobs/JobsPage";
import JobDetailPage from "../Pages/Jobs/JobDetailPage";
import EmployerPage from "../Pages/EmployerPage/EmployerPage";
import CreateJobListing from "../Pages/Jobs/CreateJobListing";

export const router = createBrowserRouter ([
    {
        path: "/",
        element: <App/>,
        children: [
            {path: "", element: <HomePage/>},
            {path: "login", element: <LoginPage/>},
            {path: "employer_page", element: <EmployerPage/>},
            {path: "job/list", element: <JobsPage/>},
            {path: "job/info/:jobId", element: <JobDetailPage/>},
            {path: "job/create", element: <CreateJobListing/>}
        ],
    }
])
