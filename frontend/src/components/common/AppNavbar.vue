<template>
  <div>
    <nav class="navbar" :class="{ 'scrolled': scrolled }">
      <div class="container">
        <div class="navbar-content">
          <router-link to="/" class="navbar-brand">
            <img src="@/assets/images/logo.png" alt="Svensk Hälsovård" />
          </router-link>
          
          <div class="nav-links">
            <router-link to="/" class="nav-link">Hem</router-link>
            <router-link to="/#services" class="nav-link">Tjänster</router-link>
            <router-link to="/contact" class="nav-link">Kontakt</router-link>
          </div>
          
          <div class="navbar-actions">
            <div class="cart-icon" @click="toggleCart">
              <i class="ri-shopping-cart-2-line"></i>
              <span v-if="cartItemCount > 0" class="badge">{{ cartItemCount }}</span>
              
              <!-- Cart Dropdown -->
              <div v-if="showCartDropdown" class="cart-dropdown">
                <div v-if="cartItems.length === 0" class="empty-cart">
                  <i class="ri-shopping-cart-line empty-cart-icon"></i>
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
                      <p>Totalt:</p>
                      <p>{{ formatPrice(cartTotal) }} kr</p>
                    </div>
                    <div class="cart-actions">
                      <router-link to="/cart" class="btn btn-outline-primary" @click="closeCartDropdown">
                        <i class="ri-shopping-cart-line"></i> Visa varukorg
                      </router-link>
                      <router-link to="/checkout" class="btn btn-primary" @click="closeCartDropdown">
                        <i class="ri-arrow-right-line"></i> Till kassan
                      </router-link>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <div class="user-icon">
              <i class="ri-user-3-line"></i>
            </div>

            <div class="mobile-menu-toggle" @click="toggleMobileMenu">
              <i class="ri-menu-line"></i>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <!-- Mobile Menu -->
    <div class="mobile-menu" :class="{ 'active': showMobileMenu }">
      <div class="mobile-menu-content">
        <div class="mobile-menu-header">
          <img src="@/assets/images/logo.png" alt="Svensk Hälsovård" height="40" />
          <div class="mobile-menu-close" @click="toggleMobileMenu">
            <i class="ri-close-line"></i>
          </div>
        </div>
        
        <div class="mobile-menu-links">
          <router-link to="/" class="nav-link" @click="showMobileMenu = false">
            <i class="ri-home-line"></i> Hem
          </router-link>
          <router-link to="/#services" class="nav-link" @click="showMobileMenu = false">
            <i class="ri-heart-pulse-line"></i> Tjänster
          </router-link>
          <router-link to="/contact" class="nav-link" @click="showMobileMenu = false">
            <i class="ri-customer-service-line"></i> Kontakt
          </router-link>
        </div>
        
        <div class="mobile-menu-actions">
          <router-link to="/cart" class="btn btn-outline-primary" @click="showMobileMenu = false">
            <i class="ri-shopping-cart-2-line"></i> Varukorg
            <span v-if="cartItemCount > 0" class="badge badge-primary">{{ cartItemCount }}</span>
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, onBeforeUnmount, watch } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';

export default {
  name: 'AppNavbar',
  setup() {
    const store = useStore();
    const route = useRoute();
    const showCartDropdown = ref(false);
    const showMobileMenu = ref(false);
    const scrolled = ref(false);
    
    const cartItems = computed(() => store.getters['cart/cartItems']);
    const cartItemCount = computed(() => store.getters['cart/cartItemCount']);
    const cartTotal = computed(() => store.getters['cart/cartTotal']);
    
    const toggleCart = (e) => {
      e.stopPropagation();
      showCartDropdown.value = !showCartDropdown.value;
      
      if (showCartDropdown.value) {
        showMobileMenu.value = false;
      }
    };
    
    const closeCartDropdown = () => {
      showCartDropdown.value = false;
    };
    
    const toggleMobileMenu = () => {
      showMobileMenu.value = !showMobileMenu.value;
      
      if (showMobileMenu.value) {
        showCartDropdown.value = false;
      }
    };
    
    const closeCart = (e) => {
      if (!e.target.closest('.cart-icon')) {
        showCartDropdown.value = false;
      }
    };
    
    const formatPrice = (price) => {
      return new Intl.NumberFormat('sv-SE').format(price);
    };
    
    const handleScroll = () => {
      if (window.scrollY > 50) {
        scrolled.value = true;
      } else {
        scrolled.value = false;
      }
    };
    
    // Close cart dropdown when clicking outside
    onMounted(() => {
      document.addEventListener('click', closeCart);
      window.addEventListener('scroll', handleScroll);
      handleScroll(); // Check initial scroll position
    });
    
    onBeforeUnmount(() => {
      document.removeEventListener('click', closeCart);
      window.removeEventListener('scroll', handleScroll);
    });
    
    // Close mobile menu and cart dropdown when route changes
    watch(() => route.path, () => {
      showMobileMenu.value = false;
      showCartDropdown.value = false;
    });
    
    return {
      showCartDropdown,
      showMobileMenu,
      scrolled,
      cartItems,
      cartItemCount,
      cartTotal,
      toggleCart,
      closeCartDropdown,
      toggleMobileMenu,
      formatPrice
    };
  }
};
</script>

<style lang="scss" scoped>
.navbar {
  background-color: rgba($light, 0.98);
  backdrop-filter: blur(10px);
  box-shadow: $box-shadow-sm;
  padding: 1rem 0;
  position: sticky;
  top: 0;
  z-index: 1000;
  transition: $transition-base;
  
  &.scrolled {
    padding: 0.75rem 0;
    box-shadow: $box-shadow;
  }
  
  .navbar-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  
  .navbar-brand {
    display: flex;
    align-items: center;
    
    img {
      height: 45px;
      transition: $transition-base;
    }
  }
  
  .nav-links {
    display: flex;
    align-items: center;
    gap: 2rem;
    
    @media (max-width: $breakpoint-md) {
      display: none;
    }
    
    .nav-link {
      color: $dark;
      font-weight: 500;
      position: relative;
      padding: 0.5rem 0;
      
      &::after {
        content: '';
        position: absolute;
        bottom: 0;
        left: 0;
        width: 0;
        height: 2px;
        background-color: $primary;
        transition: $transition-base;
      }
      
      &:hover, &.router-link-active {
        color: $primary;
        
        &::after {
          width: 100%;
        }
      }
    }
  }
  
  .navbar-actions {
    display: flex;
    align-items: center;
    gap: 1.5rem;
    
    .cart-icon, .user-icon {
      position: relative;
      font-size: 1.5rem;
      cursor: pointer;
      color: $dark;
      transition: $transition-base;
      
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
        width: 22px;
        height: 22px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 0.7rem;
        font-weight: 700;
        box-shadow: 0 2px 5px rgba($dark, 0.2);
      }
    }
    
    .cart-dropdown {
      position: absolute;
      top: calc(100% + 0.75rem);
      right: -1rem;
      width: 350px;
      background-color: white;
      border-radius: $border-radius-lg;
      box-shadow: $box-shadow;
      padding: 1.5rem;
      margin-top: 0.5rem;
      z-index: 1000;
      transform-origin: top right;
      animation: dropdown-fade 0.2s ease;
      
      @keyframes dropdown-fade {
        from {
          opacity: 0;
          transform: translateY(-10px);
        }
        to {
          opacity: 1;
          transform: translateY(0);
        }
      }
      
      .empty-cart {
        text-align: center;
        padding: 1.5rem 0;
        
        p {
          color: #64748b;
          margin-bottom: 0;
        }
      }
      
      .cart-items {
        max-height: 300px;
        overflow-y: auto;
        margin-bottom: 1.5rem;
        
        &::-webkit-scrollbar {
          width: 6px;
        }
        
        &::-webkit-scrollbar-track {
          background: #f1f1f1;
          border-radius: 10px;
        }
        
        &::-webkit-scrollbar-thumb {
          background: #cdcdcd;
          border-radius: 10px;
        }
        
        .cart-item {
          display: flex;
          align-items: center;
          padding: 0.75rem 0;
          border-bottom: 1px solid #f1f5f9;
          
          &:last-child {
            border-bottom: none;
          }
          
          .item-info {
            flex: 1;
            min-width: 0;
            
            h4 {
              font-size: 0.95rem;
              margin-bottom: 0.25rem;
              font-weight: 600;
            }
            
            p {
              margin: 0;
              color: #64748b;
              font-size: 0.85rem;
            }
          }
          
          .item-price {
            text-align: right;
            font-weight: 600;
            white-space: nowrap;
            margin-left: 1rem;
            
            p {
              margin: 0;
              font-size: 0.95rem;
              
              &:last-child {
                font-size: 0.8rem;
                color: #64748b;
                font-weight: normal;
              }
            }
          }
        }
      }
      
      .cart-footer {
        .cart-total {
          display: flex;
          justify-content: space-between;
          font-weight: 700;
          margin-bottom: 1.5rem;
          padding-top: 0.75rem;
          border-top: 1px solid #e2e8f0;
          font-size: 1.1rem;
        }
        
        .cart-actions {
          display: flex;
          gap: 0.75rem;
          
          .btn {
            flex: 1;
            font-size: 0.9rem;
            padding: 0.75rem 1rem;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 0.5rem;
            
            i {
              font-size: 1.1rem;
            }
          }
        }
      }
    }
  }

  .mobile-menu-toggle {
    display: none;
    font-size: 1.75rem;
    color: $dark;
    cursor: pointer;
    transition: $transition-base;
    
    @media (max-width: $breakpoint-md) {
      display: block;
    }
    
    &:hover {
      color: $primary;
    }
  }
}

.mobile-menu {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba($dark, 0.5);
  backdrop-filter: blur(5px);
  z-index: 1001;
  opacity: 0;
  visibility: hidden;
  transition: $transition-base;
  
  &.active {
    opacity: 1;
    visibility: visible;
    
    .mobile-menu-content {
      transform: translateX(0);
    }
  }
  
  .mobile-menu-content {
    position: absolute;
    top: 0;
    right: 0;
    width: 300px;
    height: 100%;
    background-color: $light;
    padding: 2rem;
    overflow-y: auto;
    transform: translateX(100%);
    transition: transform 0.3s ease-in-out;
    
    .mobile-menu-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 2rem;
      
      .mobile-menu-close {
        font-size: 1.5rem;
        color: $dark;
        cursor: pointer;
        transition: $transition-base;
        
        &:hover {
          color: $primary;
        }
      }
    }
    
    .mobile-menu-links {
      margin-bottom: 2rem;
      
      a {
        display: block;
        padding: 0.75rem 0;
        font-weight: 500;
        border-bottom: 1px solid #e2e8f0;
        transition: $transition-base;
        
        &:hover, &.router-link-active {
          color: $primary;
          padding-left: 0.5rem;
        }
      }
    }
    
    .mobile-menu-actions {
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }
  }
}
</style>