import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import styles from '../../styles/AuthForm.module.css';
import { loginUser } from '../../services/authService';
import Spinner from '../../components/Spinner/Spinner';

const SignIn = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [loginSuccess, setLoginSuccess] = useState(false);
    const [loginError, setLoginError] = useState('');
    const location = useLocation();
    const signUpMessage = location.state?.message;

    const handleLogin = async(event) => {
        event.preventDefault();
        setIsLoading(true);
        try {
            await loginUser(email, password);
            setIsLoading(false);
            setLoginSuccess(true);
            // TODO: redirect to next page
        } catch(error) {
            setIsLoading(false);
            setLoginSuccess(false);
            setLoginError('Invalid credentials!');
            // TODO: Handle login error
        }
    };

    return (
        <div className={styles.authForm}>
            <h2 className={styles.authFormH2}>Hello there, welcome back</h2>

            {signUpMessage && (
                <div className={styles.authFormSuccess}>
                    {signUpMessage}
                </div>
            )}

            {loginError && (
                <div className={styles.authFormError}>
                    {loginError}
                </div>
            )}

            {isLoading ? <Spinner /> : (
                <form onSubmit={handleLogin}>
                    <input type='email' placeholder='E-mail' value={email} onChange={(e) => setEmail(e.target.value)} required />
                    <input type='password' placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)} required />
                    <Link to='/forgot-password' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}>Forgot your password?</Link>
                    <button type='submit'>Sign In</button>
                    <Link to='/signup' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}>New here? Sign Up instead</Link>
                </form>
            )}            
        </div>
    );
};

export default SignIn;