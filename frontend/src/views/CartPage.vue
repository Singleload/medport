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
        <router-link to="/" class="btn btn-primary">
          <i class="ri-arrow-left-line"></i> Fortsätt handla
        </router-link>
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
              <p class="purchase-type">
                <span class="badge" :class="item.purchaseType === 'subscription' ? 'badge-primary' : 'badge-secondary'">
                  {{ item.purchaseType === 'subscription' ? 'Prenumeration' : 'Engångsköp' }}
                </span>
              </p>
              <p v-if="item.purchaseType === 'subscription' && item.service.subscriptionInterval" class="subscription-interval">
                {{ item.service.subscriptionInterval }}
              </p>
            </div>
            
            <div class="item-quantity">
              <div class="quantity-control">
                <button 
                  @click="decreaseQuantity(item)" 
                  :disabled="item.quantity <= 1"
                  class="quantity-btn"
                  aria-label="Minska antal"
                >
                  <i class="ri-subtract-line"></i>
                </button>
                <input 
                  type="number" 
                  v-model.number="item.quantity" 
                  min="1" 
                  max="10"
                  @change="updateQuantity(item)"
                  aria-label="Antal"
                />
                <button 
                  @click="increaseQuantity(item)"
                  :disabled="item.quantity >= 10"
                  class="quantity-btn"
                  aria-label="Öka antal"
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
              <button @click="removeItem(item.id)" class="remove-btn" aria-label="Ta bort från kundvagnen">
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
            
            <div class="summary-badges">
              <div class="summary-badge">
                <i class="ri-shield-check-line"></i>
                <span>Säker betalning</span>
              </div>
              <div class="summary-badge">
                <i class="ri-truck-line"></i>
                <span>Snabb leverans</span>
              </div>
            </div>
            
            <div class="summary-actions">
              <router-link to="/" class="btn btn-outline-primary">
                <i class="ri-arrow-left-line"></i> Fortsätt handla
              </router-link>
              <router-link to="/checkout" class="btn btn-primary">
                <i class="ri-arrow-right-line"></i> Till kassan
              </router-link>
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

<style lang="scss" scoped>
.cart-page {
  padding: 3rem 0;
  background-color: $background;
  
  h1 {
    margin-bottom: 2rem;
    color: $primary;
    font-weight: 700;
  }
  
  .empty-cart {
    text-align: center;
    padding: 4rem 0;
    background-color: $light;
    border-radius: $border-radius-lg;
    box-shadow: $box-shadow;
    margin-bottom: 2rem;
    
    .icon {
      font-size: 5rem;
      margin-bottom: 1.5rem;
      color: #cbd5e0;
      animation: float 3s ease-in-out infinite;
      
      @keyframes float {
        0% { transform: translateY(0); }
        50% { transform: translateY(-10px); }
        100% { transform: translateY(0); }
      }
    }
    
    h2 {
      margin-bottom: 1rem;
      font-weight: 700;
      color: $primary;
    }
    
    p {
      margin-bottom: 2rem;
      color: #4a5568;
      font-size: 1.1rem;
    }
    
    .btn {
      padding: 0.75rem 2rem;
      font-weight: 600;
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      box-shadow: $box-shadow;
      
      i {
        font-size: 1.2rem;
      }
      
      &:hover {
        transform: translateY(-3px);
        box-shadow: $box-shadow-hover;
      }
    }
  }
  
  .cart-content {
    display: flex;
    flex-wrap: wrap;
    gap: 2rem;
    
    .cart-items {
      flex: 0 0 100%;
      max-width: 100%;
      
      @media (min-width: $breakpoint-lg) {
        flex: 0 0 calc(65% - 1rem);
        max-width: calc(65% - 1rem);
      }
      
      h2 {
        margin-bottom: 1.5rem;
        color: $primary;
        font-size: 1.5rem;
        font-weight: 700;
      }
      
      .cart-item {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        padding: 1.5rem;
        margin-bottom: 1.5rem;
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        transition: $transition-base;
        
        &:hover {
          box-shadow: $box-shadow-hover;
          transform: translateY(-3px);
        }
        
        .item-image {
          flex: 0 0 100px;
          max-width: 100px;
          margin-right: 1.5rem;
          
          img {
            width: 100%;
            height: 100px;
            object-fit: cover;
            border-radius: $border-radius;
          }
        }
        
        .item-details {
          flex: 1;
          min-width: 0;
          padding-right: 1rem;
          
          h3 {
            margin-bottom: 0.5rem;
            font-size: 1.25rem;
            font-weight: 700;
            color: $dark;
          }
          
          .purchase-type {
            margin-bottom: 0.25rem;
            
            .badge {
              padding: 0.35em 0.75em;
              border-radius: 50rem;
              font-size: 0.75rem;
              font-weight: 600;
            }
          }
          
          .subscription-interval {
            color: #64748b;
            font-size: 0.9rem;
            margin: 0;
          }
        }
        
        .item-quantity {
          flex: 0 0 150px;
          max-width: 150px;
          
          .quantity-control {
            display: flex;
            align-items: stretch;
            border: 1px solid #e2e8f0;
            border-radius: $border-radius;
            overflow: hidden;
            width: fit-content;
            
            .quantity-btn {
              width: 36px;
              height: 36px;
              display: flex;
              align-items: center;
              justify-content: center;
              background-color: #f8f9fa;
              border: none;
              padding: 0;
              cursor: pointer;
              transition: $transition-base;
              flex-shrink: 0;
              
              &:hover:not(:disabled) {
                background-color: #e2e8f0;
                color: $primary;
              }
              
              &:disabled {
                cursor: not-allowed;
                opacity: 0.5;
              }
            }
            
            input {
              width: 50px;
              height: 36px;
              border: none;
              border-left: 1px solid #e2e8f0;
              border-right: 1px solid #e2e8f0;
              text-align: center;
              font-weight: 600;
              color: $dark;
              -moz-appearance: textfield;
              appearance: textfield;
              margin: 0;
              padding: 0;
              
              &::-webkit-outer-spin-button,
              &::-webkit-inner-spin-button {
                -webkit-appearance: none;
                margin: 0;
              }
              
              &:focus {
                outline: none;
              }
            }
          }
        }
        
        .item-price {
          flex: 0 0 120px;
          max-width: 120px;
          text-align: right;
          
          .price {
            font-weight: 700;
            margin-bottom: 0.25rem;
            color: $primary;
            font-size: 1.1rem;
          }
          
          .unit-price {
            font-size: 0.9rem;
            color: #64748b;
            margin: 0;
          }
        }
        
        .item-remove {
          flex: 0 0 40px;
          max-width: 40px;
          text-align: center;
          
          .remove-btn {
            background: none;
            border: none;
            color: #cbd5e0;
            cursor: pointer;
            font-size: 1.25rem;
            transition: $transition-base;
            width: 36px;
            height: 36px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            
            &:hover {
              color: $danger;
              background-color: rgba($danger, 0.1);
              transform: scale(1.1);
            }
          }
        }
        
        @media (max-width: $breakpoint-md) {
          .item-image {
            flex: 0 0 80px;
            max-width: 80px;
            margin-right: 1rem;
            
            img {
              height: 80px;
            }
          }
          
          .item-details {
            flex: 0 0 calc(100% - 80px - 1rem);
            max-width: calc(100% - 80px - 1rem);
            margin-bottom: 1rem;
          }
          
          .item-quantity, .item-price {
            flex: 0 0 calc(50% - 20px);
            max-width: calc(50% - 20px);
          }
        }
      }
    }
    
    .cart-summary {
      flex: 0 0 100%;
      max-width: 100%;
      align-self: flex-start;
      
      @media (min-width: $breakpoint-lg) {
        flex: 0 0 calc(35% - 1rem);
        max-width: calc(35% - 1rem);
      }
      
      h2 {
        margin-bottom: 1.5rem;
        color: $primary;
        font-size: 1.5rem;
        font-weight: 700;
      }
      
      .summary-content {
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        padding: 1.75rem;
        
        .summary-row {
          display: flex;
          justify-content: space-between;
          padding: 0.75rem 0;
          
          &:not(:last-child) {
            border-bottom: 1px solid #f1f5f9;
          }
          
          &.total {
            font-weight: 700;
            font-size: 1.25rem;
            margin-top: 0.5rem;
            color: $primary;
            border-top: 1px solid #e2e8f0;
            border-bottom: none;
            padding-top: 1rem;
          }
        }
        
        .summary-badges {
          margin: 1.5rem 0;
          display: flex;
          flex-direction: column;
          gap: 0.75rem;
          
          .summary-badge {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            padding: 0.75rem;
            background-color: $primary-lighter;
            border-radius: $border-radius;
            
            i {
              color: $primary;
              font-size: 1.25rem;
            }
            
            span {
              color: #4a5568;
              font-weight: 500;
            }
          }
        }
        
        .summary-actions {
          margin-top: 2rem;
          display: flex;
          flex-direction: column;
          gap: 1rem;
          
          .btn {
            width: 100%;
            padding: 0.875rem;
            font-weight: 600;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 0.5rem;
            box-shadow: $box-shadow;
            
            i {
              font-size: 1.2rem;
            }
            
            &:hover {
              transform: translateY(-3px);
              box-shadow: $box-shadow-hover;
            }
          }
        }
      }
    }
  }
}
</style>