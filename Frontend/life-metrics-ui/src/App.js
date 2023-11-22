import './App.css';
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SignIn from './containers/SignIn/SignIn';
import SignUp from './containers/SignUp/SignUp';

function App() {
  return (
    <div>
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
