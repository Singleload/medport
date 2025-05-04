import { createStore } from 'vuex';
import VuexPersistence from 'vuex-persist';
import cart from './modules/cart';
import checkout from './modules/checkout';

// Setup Vuex persistence for cart
const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['cart']
});

export default createStore({
  state: {
    loading: false,
    error: null,
    services: [
      {
        id: 1,
        name: 'Hälsokontroll - Kvinna',
        shortDescription: 'Omfattande hälsokontroll anpassad för kvinnor',
        description: 'Vår omfattande hälsokontroll för kvinnor innehåller alla viktiga hälsokontroller som rekommenderas för kvinnor i alla åldrar. Inkluderar blodprover, gynekologisk undersökning, och personlig konsultation.',
        price: 3495,
        discountedPrice: 2995,
        isSubscription: false,
        image: '/assets/images/health-check-women.jpg',
        features: [
          'Komplett blodprov (inkl. hormoner)',
          'Gynekologisk undersökning',
          'Bröstundersökning',
          'BMI och kroppssammansättning',
          'Blodtryck och hjärthälsa',
          'Läkarkonsultation med genomgång av resultat',
          'Personlig hälsoplan'
        ]
      },
      {
        id: 2,
        name: 'Hälsokontroll - Man',
        shortDescription: 'Omfattande hälsokontroll anpassad för män',
        description: 'Vår omfattande hälsokontroll för män innehåller alla viktiga hälsokontroller som rekommenderas för män i alla åldrar. Inkluderar blodprover, prostataundersökning, och personlig konsultation.',
        price: 3495,
        discountedPrice: 2995,
        isSubscription: false,
        image: '/assets/images/health-check-men.jpg',
        features: [
          'Komplett blodprov (inkl. hormoner)',
          'Prostataundersökning (PSA)',
          'BMI och kroppssammansättning',
          'Blodtryck och hjärthälsa',
          'Läkarkonsultation med genomgång av resultat',
          'Personlig hälsoplan'
        ]
      },
      {
        id: 3,
        name: 'Blodprov - Bas',
        shortDescription: 'Grundläggande blodprover för allmän hälsokontroll',
        description: 'Vårt baspaket för blodprover ger dig en bra grund för att kontrollera din allmänna hälsostatus. Perfekt för regelbundna kontroller.',
        price: 995,
        discountedPrice: null,
        isSubscription: false,
        image: '/assets/images/blood-test-basic.jpg',
        features: [
          'Hemoglobin',
          'Blodsocker',
          'Kolesterol',
          'Lever- och njurvärden',
          'Digitala resultat inom 2 arbetsdagar',
          'Kort tolkning av resultaten'
        ]
      },
      {
        id: 4,
        name: 'Blodprov - Premium',
        shortDescription: 'Omfattande blodanalys med över 20 parametrar',
        description: 'Vårt premiumpaket för blodprover ger en omfattande analys av din hälsa med över 20 olika parametrar. Inkluderar personlig läkarkonsultation för genomgång av resultaten.',
        price: 2495,
        discountedPrice: 1995,
        isSubscription: false,
        image: '/assets/images/blood-test-premium.jpg',
        features: [
          'Över 20 olika blodparametrar',
          'Hormonnivåer',
          'Vitaminer och mineraler',
          'Immunfunktion',
          'Inflammationsmarkörer',
          'Läkarkonsultation för genomgång av resultat',
          'Personlig hälsoplan'
        ]
      },
      {
        id: 5,
        name: 'Blodprov - Prenumeration',
        shortDescription: 'Regelbundna blodprover för kontinuerlig hälsokontroll',
        description: 'Vår prenumerationstjänst för blodprover ger dig möjlighet att regelbundet kontrollera din hälsa. Perfekt för dig som vill följa din hälsoutveckling över tid.',
        price: 695,
        discountedPrice: null,
        isSubscription: true,
        subscriptionInterval: 'Var 3:e månad',
        image: '/assets/images/blood-test-subscription.jpg',
        features: [
          'Bas-blodprov var 3:e månad',
          'Personlig hälsotrend över tid',
          'Digital resultatgenomgång',
          'Prioriterad bokning',
          'Förmånligt prenumerationspris'
        ]
      }
    ]
  },
  mutations: {
    SET_LOADING(state, status) {
      state.loading = status;
    },
    SET_ERROR(state, error) {
      state.error = error;
    },
    CLEAR_ERROR(state) {
      state.error = null;
    }
  },
  actions: {
    setLoading({ commit }, status) {
      commit('SET_LOADING', status);
    },
    setError({ commit }, error) {
      commit('SET_ERROR', error);
    },
    clearError({ commit }) {
      commit('CLEAR_ERROR');
    }
  },
  getters: {
    isLoading: state => state.loading,
    error: state => state.error,
    services: state => state.services,
    getServiceById: state => id => {
      return state.services.find(service => service.id === parseInt(id));
    }
  },
  modules: {
    cart,
    checkout
  },
  plugins: [vuexLocal.plugin]
});