import React from "react";

import { BrowserRouter as Router } from "react-router-dom";

import ConfigRoutes from "./routes/ConfigRoutes";

import "./App.css";


const App: React.FC = () => { // App: React.FC คือการระบุว่า App เป็น Functional Component ของ React ด้วย TypeScript

  return (

    <Router>

      <ConfigRoutes />

    </Router>

  );

};


export default App;