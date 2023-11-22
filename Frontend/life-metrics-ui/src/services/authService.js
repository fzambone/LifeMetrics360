import axios from 'axios';

const API_URL = 'http://localhost:8081/login';

const loginUser = async (email, password) => {
    try {
        const response = await axios.post(API_URL, {email, password });
        if(response.data.token) {
            localStorage.setItem('userToken', response.data.token);
        }
        return response.data;
    } catch(error) {
        console.error('Login error:', error.response.data);
        throw error;
    }
};

const logoutUser = () => {
    localStorage.removeItem('userToken');
};

export { loginUser, logoutUser };