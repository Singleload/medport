<template>
  <div class="checkout-page">
    <div class="container">
      <h1>Kassa</h1>

      <div v-if="cartItems.length === 0" class="empty-checkout">
        <div class="icon">
          <i class="ri-shopping-cart-2-line"></i>
        </div>
        <h2>Din kundvagn är tom</h2>
        <p>Du behöver lägga till artiklar i din kundvagn innan du kan gå till kassan.</p>
        <router-link to="/" class="btn btn-primary">Tillbaka till butiken</router-link>
      </div>

      <div v-else-if="checkoutComplete" class="checkout-success">
        <div class="success-icon">
          <i class="ri-check-line"></i>
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
                <i class="ri-bank-card-line"></i>
                <i class="ri-mastercard-line"></i>
                <i class="ri-visa-line"></i>
              </div>
            </div>

            <div v-if="checkoutError" class="alert alert-danger">
              <i class="ri-error-warning-line alert-icon"></i>
              <div class="alert-content">
                {{ checkoutError }}
              </div>
            </div>

            <div class="form-actions">
              <button type="submit" class="btn btn-primary" :disabled="isLoading">
                <span v-if="isLoading">
                  <i class="ri-loader-4-line ri-spin"></i> Bearbetar...
                </span>
                <span v-else>
                  <i class="ri-check-double-line"></i> Slutför beställning
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

            <div class="summary-additional">
              <div class="summary-badge">
                <i class="ri-shield-check-line"></i>
                <span>Säker betalning</span>
              </div>
              <div class="summary-badge">
                <i class="ri-lock-line"></i>
                <span>Krypterad information</span>
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
  padding: 3rem 0;
  background-color: $background;
  
  h1 {
    margin-bottom: 2rem;
    color: $primary;
    font-weight: 700;
  }
  
  .empty-checkout, .checkout-success {
    text-align: center;
    padding: 4rem 0;
    background-color: $light;
    border-radius: $border-radius-lg;
    box-shadow: $box-shadow;
    margin-bottom: 2rem;
    
    .icon, .success-icon {
      font-size: 5rem;
      margin-bottom: 1.5rem;
    }
    
    .icon {
      color: #cbd5e0;
    }
    
    .success-icon {
      color: $success;
    }
    
    h2 {
      margin-bottom: 1rem;
      font-weight: 700;
      color: $primary;
    }
    
    p {
      margin-bottom: 0.5rem;
      color: #4a5568;
      
      &:last-of-type {
        margin-bottom: 2rem;
      }
    }
    
    .order-details {
      margin: 2rem 0;
      padding: 1.5rem;
      background-color: $primary-lighter;
      border-radius: $border-radius;
      display: inline-block;
      
      p {
        margin: 0;
        color: $primary;
        font-weight: 500;
      }
    }
    
    .btn {
      padding: 0.75rem 2rem;
      font-weight: 600;
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      box-shadow: $box-shadow;
      
      &:hover {
        transform: translateY(-3px);
        box-shadow: $box-shadow-hover;
      }
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
        font-weight: 700;
      }
      
      .checkout-form {
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        padding: 2rem;
        
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
            font-weight: $font-weight-medium;
            color: $dark;
          }
          
          .form-control {
            width: 100%;
            padding: 0.75rem 1rem;
            border: 1px solid #e2e8f0;
            border-radius: $border-radius;
            transition: $transition-base;
            
            &:focus {
              border-color: $primary;
              box-shadow: 0 0 0 3px rgba($primary, 0.15);
              outline: none;
            }
          }
          
          textarea.form-control {
            resize: vertical;
            min-height: 100px;
          }
        }
        
        .payment-selection {
          margin-top: 2rem;
          border-top: 1px solid #e2e8f0;
          padding-top: 1.5rem;
          
          h3 {
            margin-bottom: 0.75rem;
            font-size: 1.25rem;
            font-weight: 700;
            color: $primary;
          }
          
          p {
            margin-bottom: 1rem;
            color: #4a5568;
          }
          
          .payment-logos {
            display: flex;
            align-items: center;
            gap: 1rem;
            background-color: #f8f9fa;
            padding: 1rem;
            border-radius: $border-radius;
            
            span {
              font-weight: 500;
              color: $dark;
            }
            
            i {
              font-size: 1.5rem;
              color: #64748b;
            }
          }
        }
        
        .alert {
          display: flex;
          align-items: flex-start;
          margin-bottom: 1.5rem;
          
          .alert-icon {
            margin-right: 0.75rem;
            margin-top: 0.2rem;
            flex-shrink: 0;
            font-size: 1.25rem;
          }
        }
        
        .form-actions {
          .btn {
            width: 100%;
            padding: 0.875rem;
            font-size: 1.1rem;
            font-weight: 600;
            box-shadow: $box-shadow;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 0.5rem;
            
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
        font-weight: 700;
      }
      
      .summary-content {
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        padding: 2rem;
        
        .summary-items {
          margin-bottom: 1.5rem;
          
          .summary-item {
            display: flex;
            justify-content: space-between;
            padding: 0.75rem 0;
            border-bottom: 1px solid #f1f5f9;
            
            &:last-child {
              border-bottom: none;
            }
            
            .item-info {
              flex: 1;
              min-width: 0;
              
              h4 {
                margin: 0 0 0.375rem;
                font-size: 1rem;
                font-weight: $font-weight-bold;
                color: $dark;
              }
              
              p {
                margin: 0;
                color: #64748b;
                font-size: 0.9rem;
              }
            }
            
            .item-price {
              font-weight: $font-weight-bold;
              margin-left: 1rem;
              color: $primary;
              white-space: nowrap;
            }
          }
        }
        
        .summary-totals {
          padding-top: 1rem;
          margin-bottom: 1.5rem;
          border-top: 1px solid #e2e8f0;
          
          .summary-row {
            display: flex;
            justify-content: space-between;
            padding: 0.5rem 0;
            
            &.total {
              font-weight: $font-weight-bold;
              font-size: 1.25rem;
              margin-top: 0.5rem;
              color: $primary;
              padding-top: 0.5rem;
              border-top: 1px solid #e2e8f0;
            }
          }
        }
        
        .summary-additional {
          border-top: 1px solid #e2e8f0;
          padding-top: 1.5rem;
          
          .summary-badge {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            margin-bottom: 0.75rem;
            
            i {
              color: $primary;
              font-size: 1.25rem;
            }
            
            span {
              color: #4a5568;
              font-size: 0.95rem;
            }
          }
        }
      }
    }
  }
}
</style>