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
                      <router-link to="/cart" class="btn btn-outline-primary" @click="showCartDropdown = false">
                        <i class="ri-shopping-cart-line"></i> Visa varukorg
                      </router-link>
                      <router-link to="/checkout" class="btn btn-primary" @click="showCartDropdown = false">
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
    
    // Close mobile menu when route changes
    watch(() => route.path, () => {
      showMobileMenu.value = false;
    });
    
    return {
      showCartDropdown,
      showMobileMenu,
      scrolled,
      cartItems,
      cartItemCount,
      cartTotal,
      toggleCart,
      toggleMobileMenu,
      formatPrice
    };
  }
};
</script>