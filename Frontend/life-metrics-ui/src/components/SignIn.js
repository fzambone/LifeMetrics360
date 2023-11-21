import React from 'react';
import { Link } from 'react-router-dom';

const SignIn = () => {
    return (
        <div className='auth-form'>
            <h2>Hello there, welcome back</h2>
            <form>
                <input type='email' placeholder='E-mail' required />
                <input type='password' placeholder='Password' required />
                <Link to='/forgot-password' className='form-link'>Forgot your password?</Link>
                <button type='submit'>Sign In</button>
                <Link to='/signup' className='form-link'>New here? Sign Up instead</Link>
            </form>
        </div>
    );
};

export default SignIn;