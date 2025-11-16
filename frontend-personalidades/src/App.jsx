import React from 'react';
import './App.css';
import logo from './assets/logo.svg';
import logogo from './assets/logo-go.png';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logogo} className="App-logo" alt="Logo Go" />
        <h1>Personalidades</h1>
        <img src={logo} className="App-logo" alt="React Logo" />
      </header>
    </div>
  );
}

export default App;
