<template>
    <div class="checkout-page">
      <div class="container">
        <h1>Kassa</h1>
  
        <div v-if="cartItems.length === 0" class="empty-checkout">
          <div class="icon">
            <i class="fas fa-shopping-cart"></i>
          </div>
          <h2>Din kundvagn är tom</h2>
          <p>Du behöver lägga till artiklar i din kundvagn innan du kan gå till kassan.</p>
          <router-link to="/" class="btn btn-primary">Tillbaka till butiken</router-link>
        </div>
  
        <div v-else-if="checkoutComplete" class="checkout-success">
          <div class="success-icon">
            <i class="fas fa-check-circle"></i>
          </div>
          <h2>Tack för din beställning!</h2>
          <p>Din order har mottagits och behandlas nu.</p>
          <p>Orderbekräftelse har skickats till din e-post.</p>
          <div class="order-details">
            <p><strong>Ordernummer:</strong> {{ bookingId }}</p>
          </div>
          <router-link to="/" class="btn btn-primary">Tillbaka till butiken</router-link>
        </div>
  
        <div v-else class="checkout-content">
          <div class="checkout-form-section">
            <h2>Dina uppgifter</h2>
  
            <form @submit.prevent="submitCheckout" class="checkout-form">
              <div class="form-row">
                <div class="form-group">
                  <label for="firstName">Förnamn *</label>
                  <input 
                    type="text" 
                    id="firstName" 
                    v-model="customer.firstName" 
                    class="form-control" 
                    required
                  >
                </div>
                <div class="form-group">
                  <label for="lastName">Efternamn *</label>
                  <input 
                    type="text" 
                    id="lastName" 
                    v-model="customer.lastName" 
                    class="form-control" 
                    required
                  >
                </div>
              </div>
  
              <div class="form-group">
                <label for="email">E-post *</label>
                <input 
                  type="email" 
                  id="email" 
                  v-model="customer.email" 
                  class="form-control" 
                  required
                >
              </div>
  
              <div class="form-group">
                <label for="phone">Telefon *</label>
                <input 
                  type="tel" 
                  id="phone" 
                  v-model="customer.phone" 
                  class="form-control" 
                  required
                >
              </div>
  
              <div class="form-group">
                <label for="streetAddress">Gatuadress *</label>
                <input 
                  type="text" 
                  id="streetAddress" 
                  v-model="customer.streetAddress" 
                  class="form-control" 
                  required
                >
              </div>
  
              <div class="form-row">
                <div class="form-group">
                  <label for="postalCode">Postnummer *</label>
                  <input 
                    type="text" 
                    id="postalCode" 
                    v-model="customer.postalCode" 
                    class="form-control" 
                    required
                  >
                </div>
                <div class="form-group">
                  <label for="city">Ort *</label>
                  <input 
                    type="text" 
                    id="city" 
                    v-model="customer.city" 
                    class="form-control" 
                    required
                  >
                </div>
              </div>
  
              <div class="form-group">
                <label for="additionalInfo">Ytterligare information (valfritt)</label>
                <textarea 
                  id="additionalInfo" 
                  v-model="customer.additionalInfo" 
                  class="form-control"
                  rows="3"
                ></textarea>
              </div>
  
              <div class="form-group payment-selection">
                <h3>Betalningsmetod</h3>
                <p>All betalning hanteras säkert via Svea Ekonomi.</p>
                <div class="payment-logos">
                  <span>Svea Ekonomi</span>
                </div>
              </div>
  
              <div v-if="checkoutError" class="alert alert-danger">
                {{ checkoutError }}
              </div>
  
              <div class="form-actions">
                <button type="submit" class="btn btn-primary" :disabled="isLoading">
                  <span v-if="isLoading">
                    <i class="fas fa-spinner fa-spin"></i> Bearbetar...
                  </span>
                  <span v-else>
                    Slutför beställning
                  </span>
                </button>
              </div>
            </form>
          </div>
  
          <div class="checkout-summary">
            <h2>Ordersammanfattning</h2>
            
            <div class="summary-content">
              <div class="summary-items">
                <div v-for="item in cartItems" :key="item.id" class="summary-item">
                  <div class="item-info">
                    <h4>{{ item.service.name }}</h4>
                    <p>{{ item.purchaseType === 'subscription' ? 'Prenumeration' : 'Engångsköp' }}</p>
                    <p v-if="item.quantity > 1">{{ item.quantity }} st</p>
                  </div>
                  <div class="item-price">
                    {{ formatPrice(getItemPrice(item)) }} kr
                  </div>
                </div>
              </div>
              
              <div class="summary-totals">
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
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { ref, computed, onMounted, watch } from 'vue';
  import { useStore } from 'vuex';
  import { useRouter } from 'vue-router';
  
  export default {
    name: 'CheckoutPage',
    setup() {
      const store = useStore();
      const router = useRouter();
      
      const customer = ref({
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        streetAddress: '',
        postalCode: '',
        city: '',
        additionalInfo: ''
      });
      
      const checkoutError = ref(null);
      const isLoading = computed(() => store.getters.isLoading);
      const cartItems = computed(() => store.getters['cart/cartItems']);
      const cartTotal = computed(() => store.getters['cart/cartTotal']);
      const paymentStatus = computed(() => store.getters['checkout/paymentStatus']);
      const bookingStatus = computed(() => store.getters['checkout/bookingStatus']);
      const bookingId = computed(() => store.getters['checkout/bookingId']);
      const checkoutComplete = computed(() => store.getters['checkout/isCheckoutComplete']);
      
      // Load customer data from store if available
      onMounted(() => {
        const savedCustomer = store.getters['checkout/customer'];
        if (savedCustomer.firstName) {
          customer.value = { ...savedCustomer };
        }
        
        // Redirect to home if cart is empty
        if (cartItems.value.length === 0 && !checkoutComplete.value) {
          router.push('/');
        }
      });
      
      // Watch for payment or booking errors
      watch(() => store.getters['checkout/paymentError'], (error) => {
        if (error) {
          checkoutError.value = `Betalningsfel: ${error}`;
        }
      });
      
      watch(() => store.getters['checkout/bookingError'], (error) => {
        if (error) {
          checkoutError.value = `Bokningsfel: ${error}`;
        }
      });
      
      const formatPrice = (price) => {
        return new Intl.NumberFormat('sv-SE').format(price);
      };
      
      const getItemPrice = (item) => {
        const unitPrice = item.service.discountedPrice || item.service.price;
        return unitPrice * item.quantity;
      };
      
      const submitCheckout = async () => {
        try {
          checkoutError.value = null;
          
          // Save customer data to store
          await store.dispatch('checkout/updateCustomer', customer.value);
          
          // Initiate checkout
          const checkoutResponse = await store.dispatch('checkout/initiateCheckout');
          
          // Process payment with Svea Ekonomi
          // In a real application, this would redirect to Svea's payment page
          // For this implementation, we'll simulate a successful payment
          await store.dispatch('checkout/processPayment', {
            paymentMethod: 'card', // This would normally come from Svea
            paymentReference: checkoutResponse.paymentId
          });
          
          // At this point, the backend would create the booking after payment confirmation
          // This is handled in the checkout store module's processPayment action
          
          // Set document title
          document.title = 'Beställning bekräftad - Svensk Hälsovård';
          
        } catch (error) {
          console.error('Checkout error:', error);
          checkoutError.value = error.message || 'Ett fel uppstod vid beställningen. Försök igen senare.';
        }
      };
      
      return {
        customer,
        cartItems,
        cartTotal,
        checkoutError,
        isLoading,
        paymentStatus,
        bookingStatus,
        bookingId,
        checkoutComplete,
        formatPrice,
        getItemPrice,
        submitCheckout
      };
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .checkout-page {
    padding: 2rem 0;
    
    h1 {
      margin-bottom: 2rem;
      color: $primary;
    }
    
    .empty-checkout, .checkout-success {
      text-align: center;
      padding: 3rem 0;
      
      .icon, .success-icon {
        font-size: 4rem;
        margin-bottom: 1.5rem;
      }
      
      .icon {
        color: #dee2e6;
      }
      
      .success-icon {
        color: $success;
      }
      
      h2 {
        margin-bottom: 1rem;
      }
      
      p {
        margin-bottom: 0.5rem;
        color: #6c757d;
      }
      
      .order-details {
        margin: 2rem 0;
        padding: 1.5rem;
        background-color: #f8f9fa;
        border-radius: $border-radius;
        display: inline-block;
        
        p {
          margin: 0;
        }
      }
      
      .btn {
        margin-top: 1.5rem;
      }
    }
    
    .checkout-content {
      display: flex;
      flex-wrap: wrap;
      gap: 2rem;
      
      .checkout-form-section {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-lg) {
          flex: 0 0 calc(60% - 1rem);
          max-width: calc(60% - 1rem);
        }
        
        h2 {
          margin-bottom: 1.5rem;
          color: $primary;
          font-size: 1.5rem;
        }
        
        .checkout-form {
          background-color: white;
          border-radius: $border-radius;
          box-shadow: $box-shadow-sm;
          padding: 1.5rem;
          
          .form-row {
            display: flex;
            flex-wrap: wrap;
            gap: 1rem;
            
            .form-group {
              flex: 1;
              min-width: 0;
            }
          }
          
          .form-group {
            margin-bottom: 1.5rem;
            
            label {
              display: block;
              margin-bottom: 0.5rem;
              font-weight: 500;
            }
            
            .form-control {
              width: 100%;
              padding: 0.75rem;
              border: 1px solid #dee2e6;
              border-radius: $border-radius-sm;
              transition: $transition-base;
              
              &:focus {
                border-color: $primary;
                box-shadow: 0 0 0 0.2rem rgba($primary, 0.25);
                outline: none;
              }
            }
          }
          
          .payment-selection {
            margin-top: 2rem;
            
            h3 {
              margin-bottom: 0.5rem;
              font-size: 1.25rem;
            }
            
            .payment-logos {
              display: flex;
              align-items: center;
              margin-top: 1rem;
              
              span {
                padding: 0.5rem 1rem;
                border: 1px solid #dee2e6;
                border-radius: $border-radius-sm;
                font-weight: 500;
              }
            }
          }
          
          .alert {
            margin-bottom: 1.5rem;
          }
          
          .form-actions {
            .btn {
              width: 100%;
              padding: 0.75rem;
              font-size: 1.1rem;
            }
          }
        }
      }
      
      .checkout-summary {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-lg) {
          flex: 0 0 calc(40% - 1rem);
          max-width: calc(40% - 1rem);
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
          
          .summary-items {
            .summary-item {
              display: flex;
              justify-content: space-between;
              padding: 0.75rem 0;
              border-bottom: 1px solid #f2f2f2;
              
              .item-info {
                flex: 1;
                min-width: 0;
                
                h4 {
                  margin: 0 0 0.25rem;
                  font-size: 1rem;
                }
                
                p {
                  margin: 0;
                  color: #6c757d;
                  font-size: 0.9rem;
                }
              }
              
              .item-price {
                font-weight: 500;
                margin-left: 1rem;
              }
            }
          }
          
          .summary-totals {
            margin-top: 1.5rem;
            
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
          }
        }
      }
    }
  }
  </style>