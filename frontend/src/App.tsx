import React from 'react';
import './styles/global.css';
import './input.css';
import './styles/tailwind.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/home';
import Header from './layouts/header';
import Sidebar from './layouts/sidebar';

function App() {
  return (
    <Router>
      <Header />
      <div style={{ paddingTop: '10%' }} className="app-container">
        <div className="content">
          <Routes>
            <Route path="/" element={<Home />} />
          </Routes>
        </div>
        <Sidebar />
      </div>
    </Router>

  );
}

export default App;
