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
import Navbar from './navbar/Navbar';
import CreateStudyProgram from './createStudyProgram/CreateStudyProgram';
import ProtectedRoute from './role-base/Auth';
import DiplomaRequest from './diplomaRequests/DiplomaRequests';
import DiplomaRequests from './diplomaRequests/DiplomaRequests';
import CreateDepartments from './createDepartments/CreateDepartments';
import CompetitionRequests from './competitionRequests/CompetitionRequests';
function App() {
  
  return (
    <Router>
      <div className="App">
        <Navbar/>
        <div className="content">
          <Switch>
            {/* <Route exact path="/"> */}
              {/* <Login/> */}
            {/* </Route> */}
            <ProtectedRoute exact path="/competition" component={CreateCompetition} roles={['ADMIN']} />
            <ProtectedRoute exact path="/competitions" component={Competitions} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/competition/:id" component={Competition} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/form" component={Form} roles={['ADMIN']} />
            <ProtectedRoute exact path="/examResults/:id" component={ExamResults} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/examResult" component={ExamResult} roles={['ADMIN']} />
            <ProtectedRoute exact path="/diploma" component={Diploma} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/departments" component={Departments} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/studyPrograms" component={StudyPrograms} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/studyProgram/:id" component={StudyProgram} roles={['ADMIN','STUDENT']} />
            <ProtectedRoute exact path="/studyProgram" component={CreateStudyProgram} roles={['ADMIN']} />
            <ProtectedRoute exact path="/diplomaRequests" component={DiplomaRequests} roles={['ADMIN']} />
            <ProtectedRoute exact path="/department" component={CreateDepartments} roles={['ADMIN']} />
            <ProtectedRoute exact path="/competitionRequests/:id" component={CompetitionRequests} roles={['ADMIN']} />
            {/* <Route exact path="/competition">
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
            <Route exact path="/diploma">
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
            <Route exact path="/studyProgram">
              <CreateStudyProgram/>
            </Route> */}
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
