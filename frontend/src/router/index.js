import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/HomePage.vue'),
    meta: {
      title: 'Svensk Hälsovård - Din väg till bättre hälsa',
      metaTags: [
        {
          name: 'description',
          content: 'Boka blodprov och hälsokontroller enkelt och snabbt. Vi erbjuder professionella hälsotjänster för alla.'
        },
        {
          property: 'og:title',
          content: 'Svensk Hälsovård - Din väg till bättre hälsa'
        },
        {
          property: 'og:description',
          content: 'Boka blodprov och hälsokontroller enkelt och snabbt. Vi erbjuder professionella hälsotjänster för alla.'
        }
      ]
    }
  },
  {
    path: '/contact',
    name: 'Contact',
    component: () => import('@/views/ContactPage.vue'),
    meta: {
      title: 'Kontakta Oss - Svensk Hälsovård',
      metaTags: [
        {
          name: 'description',
          content: 'Kontakta oss för frågor om våra tjänster eller boka tid för blodprov och hälsokontroller.'
        }
      ]
    }
  },
  {
    path: '/service/:id',
    name: 'ServiceDetails',
    component: () => import('@/views/ServiceDetailsPage.vue'),
    meta: {
      title: 'Tjänstedetaljer - Svensk Hälsovård',
      metaTags: [
        {
          name: 'description',
          content: 'Detaljerad information om våra hälsotjänster, blodprover och hälsokontroller.'
        }
      ]
    }
  },
  {
    path: '/cart',
    name: 'Cart',
    component: () => import('@/views/CartPage.vue'),
    meta: {
      title: 'Kundvagn - Svensk Hälsovård'
    }
  },
  {
    path: '/checkout',
    name: 'Checkout',
    component: () => import('@/views/CheckoutPage.vue'),
    meta: {
      title: 'Kassa - Svensk Hälsovård'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFoundPage.vue'),
    meta: {
      title: 'Sidan hittades inte - Svensk Hälsovård'
    }
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0 };
    }
  }
});

// Update document title and meta tags
router.beforeEach((to, from, next) => {
  // Update document title
  document.title = to.meta.title || 'Svensk Hälsovård';
  
  // Remove existing meta tags
  const metaTags = document.querySelectorAll('[data-vue-router-controlled]');
  metaTags.forEach(tag => tag.parentNode.removeChild(tag));
  
  // Add new meta tags
  if (to.meta.metaTags) {
    to.meta.metaTags.forEach(tagDef => {
      const tag = document.createElement('meta');
      Object.keys(tagDef).forEach(key => {
        tag.setAttribute(key, tagDef[key]);
      });
      tag.setAttribute('data-vue-router-controlled', '');
      document.head.appendChild(tag);
    });
  }
  
  next();
});

export default router;