import { Link } from "react-router-dom";

const Home = () => {
    return (  
        <div className="home">
            <h1>Home</h1>
            <button className="button">
                <Link to="/form" className="button-link">Forma</Link>
            </button>
        </div>
    );
}
 
export default Home;