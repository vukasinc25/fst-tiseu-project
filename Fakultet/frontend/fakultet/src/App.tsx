import './App.css';
import Navbar from './Navbar';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import ErrorPage from './ErrorPage';
import Form from './Form';
import Login from './Login';
function App() {
  
  return (
    <Router>
      <div className="App">
        <Navbar paragraf={"Sta mai"} title="Dje ste mangupi"/>
        <div className="content">
          <Switch>
            <Route exact path="/">
              {/* <Home /> */}
              <Login/>
            </Route>
            <Route exact path="/form">
              <Form/>
            </Route>
            <Route path="*">
              <ErrorPage/>
            </Route>
          </Switch>
        </div>
      </div>
    </Router>
  );
}

export default App;
