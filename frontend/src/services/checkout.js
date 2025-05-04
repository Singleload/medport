import api from './api';

const checkoutService = {
  /**
   * Initiates the checkout process by sending customer and cart data to the backend
   * @param {Object} checkoutData - Customer and cart data
   * @returns {Promise} - Promise with payment initiation data
   */
  async initiateCheckout(checkoutData) {
    try {
      const response = await api.post('/api/checkout/initiate', checkoutData);
      return response.data;
    } catch (error) {
      console.error('Checkout initiation error:', error);
      throw error;
    }
  },

  /**
   * Processes the payment with Svea Ekonomi via our backend
   * @param {Object} paymentData - Payment data including paymentId and payment details
   * @returns {Promise} - Promise with payment result
   */
  async processPayment(paymentData) {
    try {
      const response = await api.post('/api/checkout/payment', paymentData);
      return response.data;
    } catch (error) {
      console.error('Payment processing error:', error);
      throw error;
    }
  },

  /**
   * Verifies the payment status with the backend
   * @param {string} paymentId - The payment ID to verify
   * @returns {Promise} - Promise with payment status
   */
  async verifyPayment(paymentId) {
    try {
      const response = await api.get(`/api/checkout/verify/${paymentId}`);
      return response.data;
    } catch (error) {
      console.error('Payment verification error:', error);
      throw error;
    }
  }
};

export default checkoutService;