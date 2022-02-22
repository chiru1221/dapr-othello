import React from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
}
from "react-router-dom";
import Top from './top';
import Game from './game';
import './App.css';

const App = () => {
  return (
    <div>
      <Router>
        <Routes>
          <Route path="/" element={<Top />} />
          <Route path="/game" element={<Game />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
