import { Link } from "react-router-dom";

const ErrorPage = () => {
    return (
        <div className="errorPage">
            <h1>Glabaj ga majmune</h1>
            <Link to="/">Nazad</Link>
        </div>
    );
}
 
export default ErrorPage;