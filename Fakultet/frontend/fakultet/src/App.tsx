import Navbar from './Navbar';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import ErrorPage from './ErrorPage';
import Form from './Form';
import Login from './login/Login';
import Competitions from './competitions/Competitions';
import CreateCompetition from './createCompetition/CreateCompetition';
import Competition from './comepetition/Competiton';
import ExamResults from './examResults/ExamResults';
import ExamResult from './examResult/ExamResult';
import Diploma from './diploma/Diploma';
import Departments from './departments/Departments';
import StudyPrograms from './studyPrograms/StudyPrograms';
import StudyProgram from './studyProgram/StudyProgram';
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
              <CreateCompetition/>
            </Route>
            <Route exact path="/competitions">
              <Competitions/>
            </Route>
            <Route exact path="/competition/:id">
              <Competition/>
            </Route>
            <Route exact path="/form">
              <Form/>
            </Route>
            <Route exact path="/examResults/:id">
              <ExamResults/>
            </Route>
            <Route exact path="/examResult">
              <ExamResult/>
            </Route>
            <Route exact path="/diploma/:id">
              <Diploma/>
            </Route>
            <Route exact path="/departments">
              <Departments/>
            </Route>
            <Route exact path="/studyPrograms">
              <StudyPrograms/>
            </Route>
            <Route exact path="/studyProgram/:id">
              <StudyProgram/>
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
