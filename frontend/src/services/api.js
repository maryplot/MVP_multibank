import axios from 'axios';
import { authService } from './auth';

const API_BASE = '';

export const api = axios.create({
  baseURL: API_BASE,
});

// Текущий активный токен
let currentToken = localStorage.getItem('token');

// Функция обновления токена
const refreshToken = async () => {
  try {
    console.log('🔄 Attempting token refresh...');
    const response = await authService.login('testuser', 'password123');
    const newToken = response.token;
    
    localStorage.setItem('token', newToken);
    currentToken = newToken;
    console.log('✅ Token refreshed successfully');
    return newToken;
  } catch (error) {
    console.error('❌ Token refresh failed:', error);
    localStorage.removeItem('token');
    window.location.href = '/login';
    throw error;
  }
};

// Интерцептор для добавления токена к запросам
api.interceptors.request.use((config) => {
  if (currentToken) {
    config.headers.Authorization = `Bearer ${currentToken}`;
  }
  return config;
});

// Интерцептор для обработки ошибок и обновления токена
api.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;
    
    // Если ошибка 401 и это не запрос логина
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      
      try {
        const newToken = await refreshToken();
        originalRequest.headers.Authorization = `Bearer ${newToken}`;
        return api(originalRequest); // Повторяем запрос
      } catch (refreshError) {
        console.error('❌ Cannot refresh token, redirecting to login');
        localStorage.removeItem('token');
        window.location.href = '/login';
        return Promise.reject(refreshError);
      }
    }
    
    return Promise.reject(error);
  }
);

// Функция для ручного обновления токена
export const refreshTokenManually = async () => {
  return await refreshToken();
};