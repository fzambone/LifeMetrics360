import React from 'react';
import { Link } from 'react-router-dom';

const SignIn = () => {
    return (
        <div className="auth-container">
            <h2>Hello there, welcome back</h2>
            <form>
                <input type='email' placeholder='E-mail' required />
                <input type='password' placeholder='Password' required />
                <div className='auth-links'>
                    <Link to="/reset-password">Forgot your Password?</Link>
                </div>
                <button type='submit'>Sign In</button>
                <div className='auth-links'>
                    <Link to="/signup">New here? Sign Up instead</Link>
                </div>
            </form>
        </div>
    );
};

export default SignIn;