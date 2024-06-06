import Navbar from './Navbar';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import ErrorPage from './ErrorPage';
import Form from './Form';
import Login from './login/Login';
import Competition from './competition/Competition';
import Competitions from './competitions/Competitions';
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
            <Route exact path="/competition">
              <Competition/>
            </Route>
            <Route exact path="/competitions">
              <Competitions/>
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
