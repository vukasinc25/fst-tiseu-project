import React from 'react';
import './App.css';
import Navigationbar from './Components/Navbar';
import { Outlet } from 'react-router';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer } from 'react-toastify';
import Footer from './Components/Footer';

function App() {
  return (
    <div className="App">
      <Navigationbar/>
      <ToastContainer />
      <Outlet/>
      <Footer/>
    </div>
  );
}

export default App;
