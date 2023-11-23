import React, { useState } from 'react';
import {Link, useNavigate} from 'react-router-dom';
import styles from '../../styles/AuthForm.module.css';
import signUpUser from '../../services/registrationService';
import Spinner from '../../components/Spinner/Spinner';

const SignUp = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    // const [signUpSuccess, setSignUpSuccess] = useState(false);
    const [signUpError, setSignUpError] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const navigate = useNavigate();

    const handleSignUp = async(event) => {
        event.preventDefault();
        setIsLoading(true);
        try {
            const userData = await signUpUser(email, password, firstName, lastName);
            // setSignUpSuccess(true);
            setSignUpError('');
            navigate('/signin', {state: { fromSignUp: true, message: 'Registration sucessful. Please log in.'} });
        } catch(error) {
            setSignUpError('Signup failed. Please try again.');
            // setSignUpSuccess(false);
        } finally {
            setIsLoading(false);
        }
    };

    // if(signUpSuccess) {
    //     return (
    //         <div className={styles.authFormSuccess}>
    //             Registration successful! Please check your email to confirm your accocunt.
    //             {/* TODO: Add button or link to resend email */}
    //         </div>
    //     );
    // }
    
    return (
        <div className={styles.authForm}>
            <h2 className={styles.authFormH2}>Get on Board</h2>
            
            {signUpError && (
                <div className={styles.authFormError}>
                    {signUpError}
                </div>
            )}

            {isLoading ? <Spinner /> : (
                <form onSubmit={handleSignUp}>
                    <input type='text' placeholder='First Name' value={firstName} onChange={(e) => setFirstName(e.target.value)} required />
                    <input type='text' placeholder='Last Name' value={lastName} onChange={(e) => setLastName(e.target.value)} required />
                    <input type='email' placeholder='E-mail' value={email} onChange={(e) => setEmail(e.target.value)} required />
                    <input type='password' placeholder='Enter Password' value={password} onChange={(e) => setPassword(e.target.value)} required />
                    <input type='password' placeholder='Confirm Password' required />
                    <p className={styles.authFormP}>
                        By creating an account, you agree to the
                        <Link to='/terms' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}> Terms and Use</Link> and
                        <Link to='/privacy-policy' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}> Privacy Policy</Link>
                    </p>
                    <button type='submit'>Sign Up</button>
                    <Link to='/signin' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}>I am already a member</Link>
                </form>
            )}
            
        </div>
    );
};

export default SignUp;