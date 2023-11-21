import React from 'react';
import {Link} from 'react-router-dom';

const SignUp = () => {
    return (
        <div className='auth-form'>
            <h2>Get on Board</h2>
            <form>
                <input type='text' placeholder='Name' required />
                <input type='email' placeholder='E-mail' required />
                <input type='password' placeholder='Enter Password' required />
                <input type='password' placeholder='Confirm Password' required />
                <p>
                    By creating an account, you agree to the
                    <Link to='/terms' className='form-link'> Terms and Use</Link> and
                    <Link to='/privacy-policy' className='form-link'> Privacy Policy</Link>.
                </p>
                <button type='submit'>Sign Up</button>
                <Link to='/signin' className='form-link'>I am already a member</Link>
            </form>
        </div>
    );
};

export default SignUp