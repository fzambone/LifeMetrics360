import React, { Suspense, lazy } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';

const SignIn = lazy(() => import('./SignIn'));
const SignUp = lazy(() => import('./SignUp'));

function App() {
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/" element={<SignIn />} />
        </Routes>
      </Suspense>
    </Router>
  );
}

export default App;
