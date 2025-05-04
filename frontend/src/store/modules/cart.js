export default {
    namespaced: true,
    state: {
      items: []
    },
    mutations: {
      ADD_TO_CART(state, payload) {
        const { service, purchaseType } = payload;
        
        // Check if item already exists in cart
        const existingItemIndex = state.items.findIndex(
          item => item.service.id === service.id && item.purchaseType === purchaseType
        );
        
        if (existingItemIndex !== -1) {
          // Update quantity if item exists
          state.items[existingItemIndex].quantity += 1;
        } else {
          // Add new item if it doesn't exist
          state.items.push({
            id: Date.now(), // Unique ID for cart item
            service,
            purchaseType,
            quantity: 1
          });
        }
      },
      REMOVE_FROM_CART(state, cartItemId) {
        state.items = state.items.filter(item => item.id !== cartItemId);
      },
      UPDATE_QUANTITY(state, { cartItemId, quantity }) {
        const item = state.items.find(item => item.id === cartItemId);
        if (item) {
          item.quantity = quantity;
        }
      },
      CLEAR_CART(state) {
        state.items = [];
      }
    },
    actions: {
      addToCart({ commit }, payload) {
        commit('ADD_TO_CART', payload);
      },
      removeFromCart({ commit }, cartItemId) {
        commit('REMOVE_FROM_CART', cartItemId);
      },
      updateQuantity({ commit }, payload) {
        commit('UPDATE_QUANTITY', payload);
      },
      clearCart({ commit }) {
        commit('CLEAR_CART');
      }
    },
    getters: {
      cartItems: state => state.items,
      cartTotal: state => {
        return state.items.reduce((total, item) => {
          const price = item.service.discountedPrice || item.service.price;
          return total + (price * item.quantity);
        }, 0);
      },
      cartItemCount: state => {
        return state.items.reduce((count, item) => count + item.quantity, 0);
      },
      hasSubscriptionItems: state => {
        return state.items.some(item => item.purchaseType === 'subscription');
      }
    }
  };