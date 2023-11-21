import './App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignIn from './components/SignIn';
import SignUp from './components/SignUp';
import Header from './components/Header';

function App() {
  return (
    <div>
      <Header/>
      <Router>
        <Routes>
          <Route path='/signin' element={<SignIn />} />
          <Route path='/signup' element={<SignUp />} />
          <Route path='/' element={<SignIn />} />
        </Routes>
      </Router>
    </div>

  );
}

export default App;
