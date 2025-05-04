import api from '@/services/api';
import checkoutService from '@/services/checkout';

export default {
  namespaced: true,
  state: {
    customer: {
      firstName: '',
      lastName: '',
      email: '',
      phone: '',
      streetAddress: '',
      postalCode: '',
      city: ''
    },
    payment: {
      status: null, // null, 'pending', 'success', 'failed'
      paymentId: null,
      errorMessage: null
    },
    booking: {
      status: null, // null, 'pending', 'confirmed', 'failed'
      bookingId: null,
      errorMessage: null
    }
  },
  mutations: {
    UPDATE_CUSTOMER(state, customerData) {
      state.customer = { ...state.customer, ...customerData };
    },
    SET_PAYMENT_STATUS(state, status) {
      state.payment.status = status;
    },
    SET_PAYMENT_ID(state, paymentId) {
      state.payment.paymentId = paymentId;
    },
    SET_PAYMENT_ERROR(state, errorMessage) {
      state.payment.errorMessage = errorMessage;
    },
    SET_BOOKING_STATUS(state, status) {
      state.booking.status = status;
    },
    SET_BOOKING_ID(state, bookingId) {
      state.booking.bookingId = bookingId;
    },
    SET_BOOKING_ERROR(state, errorMessage) {
      state.booking.errorMessage = errorMessage;
    },
    RESET_CHECKOUT(state) {
      state.payment = {
        status: null,
        paymentId: null,
        errorMessage: null
      };
      state.booking = {
        status: null,
        bookingId: null,
        errorMessage: null
      };
    }
  },
  actions: {
    updateCustomer({ commit }, customerData) {
      commit('UPDATE_CUSTOMER', customerData);
    },
    async initiateCheckout({ commit, state, rootGetters }) {
      try {
        // Set payment status to pending
        commit('SET_PAYMENT_STATUS', 'pending');
        
        // Prepare checkout data
        const checkoutData = {
          customer: state.customer,
          items: rootGetters['cart/cartItems'].map(item => ({
            serviceId: item.service.id,
            quantity: item.quantity,
            purchaseType: item.purchaseType,
            price: item.service.discountedPrice || item.service.price
          })),
          totalAmount: rootGetters['cart/cartTotal']
        };
        
        // Send checkout request to the backend
        const response = await checkoutService.initiateCheckout(checkoutData);
        
        // Set payment ID from response
        commit('SET_PAYMENT_ID', response.paymentId);
        
        return response;
      } catch (error) {
        commit('SET_PAYMENT_STATUS', 'failed');
        commit('SET_PAYMENT_ERROR', error.message || 'Ett fel uppstod vid betalningsförsöket');
        throw error;
      }
    },
    async processPayment({ commit, state, dispatch }, paymentDetails) {
      try {
        // Process payment
        const response = await checkoutService.processPayment({
          paymentId: state.payment.paymentId,
          ...paymentDetails
        });
        
        // Update payment status
        commit('SET_PAYMENT_STATUS', 'success');
        
        // Create booking
        await dispatch('createBooking');
        
        return response;
      } catch (error) {
        commit('SET_PAYMENT_STATUS', 'failed');
        commit('SET_PAYMENT_ERROR', error.message || 'Ett fel uppstod vid betalningen');
        throw error;
      }
    },
    async createBooking({ commit, state, rootGetters, dispatch }) {
      try {
        // Set booking status to pending
        commit('SET_BOOKING_STATUS', 'pending');
        
        // Prepare booking data
        const bookingData = {
          paymentId: state.payment.paymentId,
          customer: state.customer,
          items: rootGetters['cart/cartItems'].map(item => ({
            serviceId: item.service.id,
            quantity: item.quantity,
            purchaseType: item.purchaseType
          }))
        };
        
        // Send booking request to backend
        const response = await api.post('/api/bookings', bookingData);
        
        // Update booking status and ID
        commit('SET_BOOKING_STATUS', 'confirmed');
        commit('SET_BOOKING_ID', response.data.bookingId);
        
        // Clear cart after successful booking
        dispatch('cart/clearCart', null, { root: true });
        
        return response.data;
      } catch (error) {
        commit('SET_BOOKING_STATUS', 'failed');
        commit('SET_BOOKING_ERROR', error.message || 'Ett fel uppstod vid bokningen');
        throw error;
      }
    },
    resetCheckout({ commit }) {
      commit('RESET_CHECKOUT');
    }
  },
  getters: {
    customer: state => state.customer,
    paymentStatus: state => state.payment.status,
    paymentId: state => state.payment.paymentId,
    paymentError: state => state.payment.errorMessage,
    bookingStatus: state => state.booking.status,
    bookingId: state => state.booking.bookingId,
    bookingError: state => state.booking.errorMessage,
    isCheckoutComplete: state => state.booking.status === 'confirmed'
  }
};