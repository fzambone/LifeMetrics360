import React from 'react';
import { Link } from 'react-router-dom';

const SignUp = () => {
    return (
        <div className='auth-container'>
            <h2>Get on Board</h2>
            <form>
                <input type='text' placeholder='First Name' required />
                <input type='text' placeholder='Last Name' required />
                <input type='email' placeholder='E-mail' required />
                <input type='password' placeholder='Enter Password' required />
                <input type='password' placeholder='Confirm Password' required />
                <div className='terms'>
                    By creating an account, you agree to the
                    <Link to="/terms"> Terms and Use</Link> and
                    <Link to="/privacy-policy"> Privacy Policy</Link>.
                </div>
                <button type='submit'>Sign Up</button>
                <div className='auth-links'>
                    <Link to="/signin">I am already a member</Link>
                </div>
            </form>
        </div>
    );
};

export default SignUp;