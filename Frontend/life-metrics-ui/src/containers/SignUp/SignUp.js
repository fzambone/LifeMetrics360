import React from 'react';
import {Link} from 'react-router-dom';
import styles from '../../styles/AuthForm.module.css';

const SignUp = () => {
    return (
        <div className={styles.authForm}>
            <h2 className={styles.authFormH2}>Get on Board</h2>
            <form>
                <input type='text' placeholder='Name' required />
                <input type='email' placeholder='E-mail' required />
                <input type='password' placeholder='Enter Password' required />
                <input type='password' placeholder='Confirm Password' required />
                <p className={styles.authFormP}>
                    By creating an account, you agree to the
                    <Link to='/terms' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}> Terms and Use</Link> and
                    <Link to='/privacy-policy' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}> Privacy Policy</Link>
                </p>
                <button type='submit'>Sign Up</button>
                <Link to='/signin' className={`${styles.authFormLink} ${styles.authFormLinkHover}`}>I am already a member</Link>
            </form>
        </div>
    );
};

export default SignUp;