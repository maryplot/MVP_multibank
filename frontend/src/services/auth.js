import { authApi } from './api';

export const authService = {
  login: async (username, password) => {
    const response = await authApi.post('/login', {
      username, 
      password 
    });
    return response.data;
  },
  
  register: async (username, email, password) => {
    const response = await authApi.post('/register', {
      username, 
      email, 
      password 
    });
    return response.data;
  },
  
  // Проверка валидности токена
  validateToken: async (token) => {
    try {
      const response = await authApi.get('/validate', {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      return response.data.valid;
    } catch (error) {
      return false;
    }
  }
};