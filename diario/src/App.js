import React from 'react';
import { Paginas } from "./componentes/Paginas.js";
import { BrowserRouter as Router } from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <Router>
        <Paginas />
      </Router>
    </div>
  );
}

export default App;
