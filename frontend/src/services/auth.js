import { api } from './api';  // ← ДОБАВИТЬ ЭТОТ ИМПОРТ!

export const authService = {
  login: async (username, password) => {
    const response = await api.post('/api/login', { 
      username, 
      password 
    });
    return response.data;
  },
  
  register: async (username, email, password) => {
    const response = await api.post('/api/register', { 
      username, 
      email, 
      password 
    });
    return response.data;
  },
  
  // Проверка валидности токена
  validateToken: async (token) => {
    try {
      const response = await api.post('/api/validate', { token });
      return response.data.valid;
    } catch (error) {
      return false;
    }
  }
};