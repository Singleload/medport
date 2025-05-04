<template>
    <div class="cart-page">
      <div class="container">
        <h1>Kundvagn</h1>
        
        <div v-if="cartItems.length === 0" class="empty-cart">
          <div class="icon">
            <i class="fas fa-shopping-cart"></i>
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
                    <i class="fas fa-minus"></i>
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
                    <i class="fas fa-plus"></i>
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
                  <i class="fas fa-trash-alt"></i>
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
  
  <style lang="scss" scoped>
  .cart-page {
    padding: 2rem 0;
    
    h1 {
      margin-bottom: 2rem;
      color: $primary;
    }
    
    .empty-cart {
      text-align: center;
      padding: 3rem 0;
      
      .icon {
        font-size: 4rem;
        color: #dee2e6;
        margin-bottom: 1.5rem;
      }
      
      h2 {
        margin-bottom: 1rem;
      }
      
      p {
        margin-bottom: 2rem;
        color: #6c757d;
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
        }
        
        .cart-item {
          display: flex;
          flex-wrap: wrap;
          align-items: center;
          padding: 1.5rem;
          margin-bottom: 1rem;
          background-color: white;
          border-radius: $border-radius;
          box-shadow: $box-shadow-sm;
          
          .item-image {
            flex: 0 0 80px;
            max-width: 80px;
            margin-right: 1rem;
            
            img {
              width: 100%;
              border-radius: $border-radius-sm;
            }
          }
          
          .item-details {
            flex: 1;
            min-width: 0;
            padding-right: 1rem;
            
            h3 {
              margin-bottom: 0.25rem;
              font-size: 1.2rem;
            }
            
            .purchase-type {
              color: #6c757d;
              margin-bottom: 0.25rem;
            }
          }
          
          .item-quantity {
            flex: 0 0 120px;
            max-width: 120px;
            
            .quantity-control {
              display: flex;
              align-items: center;
              
              .quantity-btn {
                width: 30px;
                height: 30px;
                display: flex;
                align-items: center;
                justify-content: center;
                background-color: #f8f9fa;
                border: 1px solid #dee2e6;
                cursor: pointer;
                
                &:hover:not(:disabled) {
                  background-color: #e9ecef;
                }
                
                &:disabled {
                  cursor: not-allowed;
                  opacity: 0.5;
                }
                
                &:first-child {
                  border-radius: $border-radius-sm 0 0 $border-radius-sm;
                }
                
                &:last-child {
                  border-radius: 0 $border-radius-sm $border-radius-sm 0;
                }
              }
              
              input {
                width: 40px;
                height: 30px;
                border: 1px solid #dee2e6;
                border-left: none;
                border-right: none;
                text-align: center;
                -moz-appearance: textfield;
                appearance: textfield;
                
                &::-webkit-outer-spin-button,
                &::-webkit-inner-spin-button {
                  -webkit-appearance: none;
                  margin: 0;
                }
              }
            }
          }
          
          .item-price {
            flex: 0 0 120px;
            max-width: 120px;
            text-align: right;
            
            .price {
              font-weight: 500;
              margin-bottom: 0.25rem;
            }
            
            .unit-price {
              font-size: 0.9rem;
              color: #6c757d;
            }
          }
          
          .item-remove {
            flex: 0 0 40px;
            max-width: 40px;
            text-align: center;
            
            .remove-btn {
              background: none;
              border: none;
              color: #dc3545;
              cursor: pointer;
              font-size: 1.1rem;
              
              &:hover {
                color: darken(#dc3545, 10%);
              }
            }
          }
          
          @media (max-width: $breakpoint-md) {
            .item-image {
              flex: 0 0 60px;
              max-width: 60px;
            }
            
            .item-details {
              flex: 0 0 calc(100% - 60px - 1rem);
              max-width: calc(100% - 60px - 1rem);
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
        
        @media (min-width: $breakpoint-lg) {
          flex: 0 0 calc(35% - 1rem);
          max-width: calc(35% - 1rem);
        }
        
        h2 {
          margin-bottom: 1.5rem;
          color: $primary;
          font-size: 1.5rem;
        }
        
        .summary-content {
          background-color: white;
          border-radius: $border-radius;
          box-shadow: $box-shadow-sm;
          padding: 1.5rem;
          
          .summary-row {
            display: flex;
            justify-content: space-between;
            padding: 0.75rem 0;
            
            &:not(:last-child) {
              border-bottom: 1px solid #f2f2f2;
            }
            
            &.total {
              font-weight: bold;
              font-size: 1.2rem;
              margin-top: 0.5rem;
              color: $primary;
            }
          }
          
          .summary-actions {
            margin-top: 2rem;
            display: flex;
            flex-direction: column;
            gap: 1rem;
            
            .btn {
              width: 100%;
              padding: 0.75rem;
            }
          }
        }
      }
    }
  }
  </style>