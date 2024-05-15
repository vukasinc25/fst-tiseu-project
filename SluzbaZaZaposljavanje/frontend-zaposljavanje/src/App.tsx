import React from 'react';
import './App.css';
import Navigationbar from './Components/Navbar';
import LoginPage from './Pages/LoginPage';
import { Outlet } from 'react-router';

function App() {
  return (
    <div className="App">
      <Navigationbar/>
      <Outlet/>
    </div>
  );
}

export default App;
