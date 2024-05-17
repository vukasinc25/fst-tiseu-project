import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import Navigationbar from "../Components/Navbar";
import LoginPage from "../Pages/LoginPage";
import HomePage from "../Pages/HomePage";
import JobsPage from "../Pages/JobsPage";

export const router = createBrowserRouter ([
    {
        path: "/",
        element: <App/>,
        children: [
            {path: "", element: <JobsPage/>},
            {path: "login", element: <LoginPage/>}
        ],
    }
])