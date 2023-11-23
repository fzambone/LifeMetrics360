import axios from 'axios';

const signUpUser = async (email, password, firstName, lastName) => {
    try {
        const response = await axios.post('http://localhost:8081/users', {
            email,
            password,
            firstName,
            lastName,
            roles: ['user'],
        });
        return response.data;
    } catch(error) {
        console.error('Registration error:', error.response.data);
        throw error;
    }
    // throw new Error("Signup failed for testing purposes");
};

export default signUpUser;