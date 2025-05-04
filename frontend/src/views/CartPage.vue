<template>
  <div class="cart-page">
    <div class="container">
      <h1>Kundvagn</h1>
      
      <div v-if="cartItems.length === 0" class="empty-cart">
        <div class="icon">
          <i class="ri-shopping-cart-2-line"></i>
        </div>
        <h2>Din kundvagn är tom</h2>
        <p>Du har inte lagt till några artiklar i din kundvagn än.</p>
        <router-link to="/" class="btn btn-primary">Fortsätt handla</router-link>
      </div>
      
      <div v-else class="cart-content">
        <div class="cart-items">
          <h2>Artiklar</h2>
          
          <div v-for="item in cartItems" :key="item.id" class="cart-item">
            <div class="item-image">
              <img :src="item.service.image" :alt="item.service.name" />
            </div>
            
            <div class="item-details">
              <h3>{{ item.service.name }}</h3>
              <p class="purchase-type">{{ item.purchaseType === 'subscription' ? 'Prenumeration' : 'Engångsköp' }}</p>
              <p v-if="item.purchaseType === 'subscription' && item.service.subscriptionInterval">
                {{ item.service.subscriptionInterval }}
              </p>
            </div>
            
            <div class="item-quantity">
              <div class="quantity-control">
                <button 
                  @click="decreaseQuantity(item)" 
                  :disabled="item.quantity <= 1"
                  class="quantity-btn"
                >
                  <i class="ri-subtract-line"></i>
                </button>
                <input 
                  type="number" 
                  v-model.number="item.quantity" 
                  min="1" 
                  max="10"
                  @change="updateQuantity(item)"
                />
                <button 
                  @click="increaseQuantity(item)"
                  :disabled="item.quantity >= 10"
                  class="quantity-btn"
                >
                  <i class="ri-add-line"></i>
                </button>
              </div>
            </div>
            
            <div class="item-price">
              <p class="price">{{ formatPrice(getItemPrice(item)) }} kr</p>
              <p v-if="item.quantity > 1" class="unit-price">
                {{ formatPrice(item.service.discountedPrice || item.service.price) }} kr / st
              </p>
            </div>
            
            <div class="item-remove">
              <button @click="removeItem(item.id)" class="remove-btn">
                <i class="ri-delete-bin-line"></i>
              </button>
            </div>
          </div>
        </div>
        
        <div class="cart-summary">
          <h2>Sammanfattning</h2>
          
          <div class="summary-content">
            <div class="summary-row">
              <span>Delsumma</span>
              <span>{{ formatPrice(cartTotal) }} kr</span>
            </div>
            
            <div class="summary-row">
              <span>Leverans</span>
              <span>Ingår</span>
            </div>
            
            <div class="summary-row total">
              <span>Totalt</span>
              <span>{{ formatPrice(cartTotal) }} kr</span>
            </div>
            
            <div class="summary-actions">
              <router-link to="/" class="btn btn-secondary">Fortsätt handla</router-link>
              <router-link to="/checkout" class="btn btn-primary">Till kassan</router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, watch } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'CartPage',
  setup() {
    const store = useStore();
    
    const cartItems = computed(() => store.getters['cart/cartItems']);
    const cartTotal = computed(() => store.getters['cart/cartTotal']);
    
    const formatPrice = (price) => {
      return new Intl.NumberFormat('sv-SE').format(price);
    };
    
    const getItemPrice = (item) => {
      const unitPrice = item.service.discountedPrice || item.service.price;
      return unitPrice * item.quantity;
    };
    
    const increaseQuantity = (item) => {
      if (item.quantity < 10) {
        store.dispatch('cart/updateQuantity', {
          cartItemId: item.id,
          quantity: item.quantity + 1
        });
      }
    };
    
    const decreaseQuantity = (item) => {
      if (item.quantity > 1) {
        store.dispatch('cart/updateQuantity', {
          cartItemId: item.id,
          quantity: item.quantity - 1
        });
      }
    };
    
    const updateQuantity = (item) => {
      // Ensure quantity is within limits
      let quantity = item.quantity;
      if (quantity < 1) quantity = 1;
      if (quantity > 10) quantity = 10;
      
      store.dispatch('cart/updateQuantity', {
        cartItemId: item.id,
        quantity
      });
    };
    
    const removeItem = (itemId) => {
      store.dispatch('cart/removeFromCart', itemId);
    };
    
    // Watch for empty cart and set document title
    watch(cartItems, (items) => {
      document.title = items.length > 0 
        ? `Kundvagn (${items.length}) - Svensk Hälsovård` 
        : 'Kundvagn - Svensk Hälsovård';
    }, { immediate: true });
    
    return {
      cartItems,
      cartTotal,
      formatPrice,
      getItemPrice,
      increaseQuantity,
      decreaseQuantity,
      updateQuantity,
      removeItem
    };
  }
};
</script>