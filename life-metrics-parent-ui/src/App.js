import React, { Suspense, lazy } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';

const SignInMFE = lazy(() => import('auth_mfe/SingIn'));
const SignUpMFE = lazy(() => import('auth_mfe/SingUp'));

function App() {
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <nav>
          <Link to="/signin">Sign In</Link>
          <Link to="/singup">Sign Up</Link>
        </nav>
        <Routes>
          <Route path="/signin" element={<SignInMFE />}/>
          <Route path="/signup" element={<SignUpMFE />}/>
        </Routes>
      </Suspense>
    </Router>
  );
}

export default App;
