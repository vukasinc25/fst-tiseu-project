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
function App() {
  
  return (
    <Router>
      <div className="App">
        <Navbar/>
        <div className="content">
          <Switch>
            <Route exact path="/">
              {/* <Home /> */}
              <Login/>
            </Route>
            <ProtectedRoute exact path="/competition" component={CreateCompetition} roles={['ADMIN']} />
            <ProtectedRoute exact path="/competitions" component={Competitions} roles={['ADMIN']} />
            <ProtectedRoute exact path="/competition/:id" component={Competition} roles={['ADMIN']} />
            <ProtectedRoute exact path="/form" component={Form} roles={['ADMIN']} />
            <ProtectedRoute exact path="/examResults/:id" component={ExamResults} roles={['ADMIN']} />
            <ProtectedRoute exact path="/examResult" component={ExamResult} roles={['ADMIN']} />
            <ProtectedRoute exact path="/diploma" component={Diploma} roles={['ADMIN']} />
            <ProtectedRoute exact path="/departments" component={Departments} roles={['ADMIN']} />
            <ProtectedRoute exact path="/studyPrograms" component={StudyPrograms} roles={['ADMIN']} />
            <ProtectedRoute exact path="/studyProgram/:id" component={StudyProgram} roles={['ADMIN']} />
            <ProtectedRoute exact path="/studyProgram" component={CreateStudyProgram} roles={['ADMIN']} />
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
