import React from "react";

import "./App.css";

import { Outlet } from "react-router";

function App() {
  return (
    <div className="App">
      <h1>APR CROSO</h1>
      <Outlet />
      <button> Create Firm</button>
    </div>
  );
}

export default App;
