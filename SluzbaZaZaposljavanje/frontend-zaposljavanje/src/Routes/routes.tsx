import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import Navigationbar from "../Components/Navbar";
import LoginPage from "../Pages/LoginPage";
import HomePage from "../Pages/HomePage";

export const router = createBrowserRouter ([
    {
        path: "/",
        element: <App/>,
        children: [
            {path: "", element: <HomePage/>},
            {path: "login", element: <LoginPage/>}
        ],
    }
])