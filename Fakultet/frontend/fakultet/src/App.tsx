import React, { useState } from 'react';
import './App.css';
import Navbar from './Navbar';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Home from './Home';
import ErrorPage from './ErrorPage';
import Form from './Form';
function App() {
  
  return (
    <Router>
      <div className="App">
        <Navbar paragraf={"Sta mai"} title="Dje ste mangupi"/>
        <div className="content">
          <Switch>
            <Route exact path="/">
              <Home />
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
