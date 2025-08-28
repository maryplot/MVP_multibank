import axios from 'axios';

const API_BASE = '';

export const api = axios.create({
  baseURL: API_BASE,
});

// Логирование всех запросов
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  console.log('🔄 API Request:', config.method?.toUpperCase(), config.url);
  console.log('📦 Request Data:', config.data);
  console.log('🔑 Token:', token ? 'Present' : 'Missing');
  
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Логирование всех ответов
api.interceptors.response.use(
  (response) => {
    console.log('✅ API Response:', response.status, response.data);
    return response;
  },
  (error) => {
    console.error('❌ API Error:', error.response?.status, error.response?.data);
    return Promise.reject(error);
  }
);