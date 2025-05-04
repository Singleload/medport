<template>
    <nav class="navbar">
      <div class="container">
        <div class="navbar-content">
          <router-link to="/" class="navbar-brand">
            <img src="@/assets/images/logo.png" alt="Svensk Hälsovård" />
          </router-link>
          
          <div class="nav-links">
            <router-link to="/" class="nav-link">Hem</router-link>
            <router-link to="/contact" class="nav-link">Kontakt</router-link>
          </div>
          
          <div class="navbar-actions">
            <div class="cart-icon" @click="toggleCart">
              <i class="fas fa-shopping-cart"></i>
              <span v-if="cartItemCount > 0" class="badge">{{ cartItemCount }}</span>
              
              <!-- Cart Dropdown -->
              <div v-if="showCartDropdown" class="cart-dropdown">
                <div v-if="cartItems.length === 0" class="empty-cart">
                  <p>Din varukorg är tom</p>
                </div>
                <div v-else>
                  <div class="cart-items">
                    <div v-for="item in cartItems" :key="item.id" class="cart-item">
                      <div class="item-info">
                        <h4>{{ item.service.name }}</h4>
                        <p>
                          {{ item.purchaseType === 'subscription' ? 'Prenumeration' : 'Engångsköp' }}
                        </p>
                      </div>
                      <div class="item-price">
                        <p>{{ formatPrice(item.service.discountedPrice || item.service.price) }} kr</p>
                        <p v-if="item.quantity > 1">x{{ item.quantity }}</p>
                      </div>
                    </div>
                  </div>
                  <div class="cart-footer">
                    <div class="cart-total">
                      <p>Totalt: {{ formatPrice(cartTotal) }} kr</p>
                    </div>
                    <div class="cart-actions">
                      <router-link to="/cart" class="btn btn-secondary" @click="showCartDropdown = false">
                        Visa varukorg
                      </router-link>
                      <router-link to="/checkout" class="btn btn-primary" @click="showCartDropdown = false">
                        Till kassan
                      </router-link>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="user-icon">
              <i class="fas fa-user"></i>
            </div>
          </div>
        </div>
      </div>
    </nav>
  </template>
  
  <script>
  import { ref, computed, onMounted, onBeforeUnmount } from 'vue';
  import { useStore } from 'vuex';
  
  export default {
    name: 'Navbar',
    setup() {
      const store = useStore();
      const showCartDropdown = ref(false);
      
      const cartItems = computed(() => store.getters['cart/cartItems']);
      const cartItemCount = computed(() => store.getters['cart/cartItemCount']);
      const cartTotal = computed(() => store.getters['cart/cartTotal']);
      
      const toggleCart = () => {
        showCartDropdown.value = !showCartDropdown.value;
      };
      
      const closeCart = (e) => {
        if (!e.target.closest('.cart-icon')) {
          showCartDropdown.value = false;
        }
      };
      
      const formatPrice = (price) => {
        return new Intl.NumberFormat('sv-SE').format(price);
      };
      
      // Close cart dropdown when clicking outside
      onMounted(() => {
        document.addEventListener('click', closeCart);
      });
      
      onBeforeUnmount(() => {
        document.removeEventListener('click', closeCart);
      });
      
      return {
        showCartDropdown,
        cartItems,
        cartItemCount,
        cartTotal,
        toggleCart,
        formatPrice
      };
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .navbar {
    position: sticky;
    top: 0;
    z-index: 1000;
    
    .navbar-content {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 1rem 0;
    }
    
    .navbar-brand {
      img {
        height: 40px;
      }
    }
    
    .nav-links {
      display: flex;
      gap: 1.5rem;
      
      .nav-link {
        color: $dark;
        font-weight: 500;
        
        &:hover, &.router-link-active {
          color: $primary;
        }
      }
    }
    
    .navbar-actions {
      display: flex;
      align-items: center;
      gap: 1.5rem;
      
      .cart-icon, .user-icon {
        position: relative;
        font-size: 1.25rem;
        cursor: pointer;
        
        &:hover {
          color: $primary;
        }
        
        .badge {
          position: absolute;
          top: -8px;
          right: -8px;
          background-color: $primary;
          color: white;
          border-radius: 50%;
          width: 20px;
          height: 20px;
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: 0.75rem;
        }
      }
    }
    
    .cart-dropdown {
      position: absolute;
      top: 100%;
      right: 0;
      width: 320px;
      background-color: white;
      border-radius: $border-radius;
      box-shadow: $box-shadow;
      padding: 1rem;
      margin-top: 0.5rem;
      z-index: 1000;
      
      .empty-cart {
        text-align: center;
        padding: 1rem;
      }
      
      .cart-items {
        max-height: 300px;
        overflow-y: auto;
        
        .cart-item {
          display: flex;
          justify-content: space-between;
          align-items: center;
          padding: 0.75rem 0;
          border-bottom: 1px solid #eee;
          
          &:last-child {
            border-bottom: none;
          }
          
          .item-info {
            flex: 1;
            
            h4 {
              font-size: 0.9rem;
              margin-bottom: 0.25rem;
            }
            
            p {
              font-size: 0.8rem;
              color: #666;
              margin: 0;
            }
          }
          
          .item-price {
            text-align: right;
            
            p {
              margin: 0;
              font-size: 0.9rem;
            }
          }
        }
      }
      
      .cart-footer {
        margin-top: 1rem;
        
        .cart-total {
          display: flex;
          justify-content: space-between;
          font-weight: bold;
          margin-bottom: 1rem;
        }
        
        .cart-actions {
          display: flex;
          gap: 0.5rem;
          
          .btn {
            flex: 1;
            font-size: 0.9rem;
          }
        }
      }
    }
  }
  
  @media (max-width: $breakpoint-md) {
    .navbar {
      .nav-links {
        display: none;
      }
    }
  }
  </style>