import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/styles/main.scss';

const app = createApp(App);

app.use(store);
app.use(router);

// Global error handler
app.config.errorHandler = (err, vm, info) => {
  console.error('Vue Error:', err);
  console.error('Info:', info);
  // Log to server in production
  if (process.env.NODE_ENV === 'production') {
    // Send to error tracking service
  }
};

app.mount('#app');