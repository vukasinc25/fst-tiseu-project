import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import Navigationbar from "../Components/Navbar";
import LoginPage from "../Pages/LoginPage";
import HomePage from "../Pages/HomePage";
import JobsPage from "../Pages/JobsPage";
import JobDetailPage from "../Pages/JobDetailPage";

export const router = createBrowserRouter ([
    {
        path: "/",
        element: <App/>,
        children: [
            {path: "", element: <HomePage/>},
            {path: "login", element: <LoginPage/>},
            {path: "job_list", element: <JobsPage/>},
            {path: "job_info/:jobId", element: <JobDetailPage/>}
        ],
    }
])
