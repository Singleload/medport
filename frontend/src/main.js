import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './assets/styles/main.scss';

// Add the remixicon CSS globally
import 'remixicon/fonts/remixicon.css';

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

// Add a directive for smooth scroll
app.directive('smooth-scroll', {
  mounted(el, binding) {
    el.addEventListener('click', (e) => {
      e.preventDefault();
      const targetId = binding.value || el.getAttribute('href');
      if (targetId && targetId.startsWith('#')) {
        const targetElement = document.querySelector(targetId);
        if (targetElement) {
          // Smooth scroll to the element
          targetElement.scrollIntoView({ 
            behavior: 'smooth',
            block: 'start'
          });
          
          // Update URL hash without scrolling
          history.pushState(null, null, targetId);
        }
      }
    });
  }
});

// Add global methods
app.config.globalProperties.$formatPrice = (price) => {
  return new Intl.NumberFormat('sv-SE').format(price);
};

// Apply scroll event to handle navbar styling
let scrollTimeout;
window.addEventListener('scroll', () => {
  clearTimeout(scrollTimeout);
  scrollTimeout = setTimeout(() => {
    const navbar = document.querySelector('.navbar');
    if (navbar) {
      if (window.scrollY > 50) {
        navbar.classList.add('scrolled');
      } else {
        navbar.classList.remove('scrolled');
      }
    }
  }, 100);
});

app.mount('#app');