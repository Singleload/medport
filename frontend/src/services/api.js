import axios from 'axios';
import store from '@/store';

// Create axios instance with default config
const api = axios.create({
  baseURL: process.env.VUE_APP_API_URL || '/api',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  timeout: 30000 // 30 seconds
});

// Request interceptor
api.interceptors.request.use(
  config => {
    // Set loading state
    store.dispatch('setLoading', true);
    // Clear previous errors
    store.dispatch('clearError');
    return config;
  },
  error => {
    store.dispatch('setLoading', false);
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  response => {
    store.dispatch('setLoading', false);
    return response;
  },
  error => {
    store.dispatch('setLoading', false);
    
    // Handle specific error cases
    if (error.response) {
      // Server responded with a status code that falls out of the range of 2xx
      const { status, data } = error.response;
      
      switch (status) {
        case 400:
          store.dispatch('setError', {
            code: 'BAD_REQUEST',
            message: data.message || 'Ogiltig förfrågan'
          });
          break;
        case 401:
          store.dispatch('setError', {
            code: 'UNAUTHORIZED',
            message: 'Obehörig åtkomst'
          });
          break;
        case 404:
          store.dispatch('setError', {
            code: 'NOT_FOUND',
            message: 'Resursen kunde inte hittas'
          });
          break;
        case 422:
          store.dispatch('setError', {
            code: 'VALIDATION_ERROR',
            message: data.message || 'Valideringsfel',
            errors: data.errors
          });
          break;
        case 500:
          store.dispatch('setError', {
            code: 'SERVER_ERROR',
            message: 'Ett serverfel inträffade'
          });
          break;
        default:
          store.dispatch('setError', {
            code: 'UNKNOWN_ERROR',
            message: 'Ett okänt fel inträffade'
          });
      }
    } else if (error.request) {
      // The request was made but no response was received
      store.dispatch('setError', {
        code: 'NETWORK_ERROR',
        message: 'Kunde inte ansluta till servern. Kontrollera din internetanslutning.'
      });
    } else {
      // Something happened in setting up the request that triggered an Error
      store.dispatch('setError', {
        code: 'REQUEST_ERROR',
        message: error.message || 'Ett fel uppstod vid förfrågan'
      });
    }
    
    return Promise.reject(error);
  }
);

export default api;