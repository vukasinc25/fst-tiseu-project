import { Link } from "react-router-dom";

const ErrorPage = () => {
    return (
        <div className="errorPage">
            <h1>Stranica ne postoji</h1>
            <Link to="/competitions">Nazad</Link>
        </div>
    );
}
 
export default ErrorPage;