import { api } from './api';

export const authService = {
  login: async (username, password) => {
    const response = await api.post('/api/login', {  // ← Добавляем /api/
      username, 
      password 
    });
    return response.data;
  },
  
  register: async (username, email, password) => {
    const response = await api.post('/api/register', {  // ← Добавляем /api/
      username, 
      email, 
      password 
    });
    return response.data;
  }
};