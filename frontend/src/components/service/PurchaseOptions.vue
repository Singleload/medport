<template>
  <div class="purchase-options">
    <div class="option-tabs">
      <div 
        class="tab" 
        :class="{ active: selectedOption === 'one-time' }" 
        @click="selectedOption = 'one-time'"
      >
        <i class="ri-shopping-cart-2-line"></i> Engångsköp
      </div>
      <div 
        v-if="service.isSubscription"
        class="tab" 
        :class="{ active: selectedOption === 'subscription' }" 
        @click="selectedOption = 'subscription'"
      >
        <i class="ri-repeat-line"></i> Prenumerera
      </div>
    </div>
    
    <div class="option-content">
      <div v-if="selectedOption === 'one-time'" class="one-time-option">
        <div class="price">
          <span v-if="service.discountedPrice" class="original-price">{{ formatPrice(service.price) }} kr</span>
          <span class="current-price">{{ formatPrice(service.discountedPrice || service.price) }} kr</span>
        </div>
        <div class="description">
          <p>Beställ nu för engångsanvändning. Passa på att göra en hälsokontroll för att få en nulägesbild av din hälsa.</p>
        </div>
        <div class="actions">
          <button @click="addToCart('one-time')" class="btn btn-primary">
            <i class="ri-shopping-cart-2-line"></i> Lägg i kundvagn
          </button>
        </div>
      </div>
      
      <div v-if="selectedOption === 'subscription' && service.isSubscription" class="subscription-option">
        <div class="price">
          <span class="current-price">{{ formatPrice(service.price) }} kr</span>
          <span class="interval">/ {{ service.subscriptionInterval }}</span>
        </div>
        <div class="description">
          <p>Prenumerera för regelbundna kontroller. Perfekt för dig som vill följa din hälsoutveckling över tid.</p>
          <ul class="benefits">
            <li>
              <i class="ri-check-line"></i>
              <span>Automatisk leverans</span>
            </li>
            <li>
              <i class="ri-check-line"></i>
              <span>Förmånligt prenumerationspris</span>
            </li>
            <li>
              <i class="ri-check-line"></i>
              <span>Avsluta när du vill</span>
            </li>
          </ul>
        </div>
        <div class="actions">
          <button @click="addToCart('subscription')" class="btn btn-primary">
            <i class="ri-repeat-line"></i> Prenumerera nu
          </button>
        </div>
      </div>
    </div>
    
    <div v-if="showNotification" class="notification" :class="notificationType">
      <i v-if="notificationType === 'success'" class="ri-check-line"></i>
      <i v-else class="ri-error-warning-line"></i>
      <span>{{ notificationMessage }}</span>
    </div>
  </div>
</template>

<script>
import { ref, watch } from 'vue';
import { useStore } from 'vuex';

export default {
  name: 'PurchaseOptions',
  props: {
    service: {
      type: Object,
      required: true
    }
  },
  setup(props) {
    const store = useStore();
    const selectedOption = ref('one-time');
    const showNotification = ref(false);
    const notificationMessage = ref('');
    const notificationType = ref('');
    
    // Reset selected option if service changes
    watch(() => props.service, () => {
      selectedOption.value = 'one-time';
    });
    
    const formatPrice = (price) => {
      return new Intl.NumberFormat('sv-SE').format(price);
    };
    
    const addToCart = (purchaseType) => {
      store.dispatch('cart/addToCart', {
        service: props.service,
        purchaseType
      });
      
      // Show success notification
      notificationMessage.value = 'Produkten har lagts till i kundvagnen';
      notificationType.value = 'success';
      showNotification.value = true;
      
      // Hide notification after 3 seconds
      setTimeout(() => {
        showNotification.value = false;
      }, 3000);
    };
    
    return {
      selectedOption,
      showNotification,
      notificationMessage,
      notificationType,
      formatPrice,
      addToCart
    };
  }
};
</script>

<style lang="scss" scoped>
.purchase-options {
  position: relative;
  
  .option-tabs {
    display: flex;
    border-bottom: 1px solid #e2e8f0;
    margin-bottom: 1.5rem;
    
    .tab {
      display: flex;
      align-items: center;
      gap: 0.5rem;
      padding: 1rem 1.5rem;
      cursor: pointer;
      font-weight: 600;
      color: #64748b;
      transition: $transition-base;
      
      i {
        font-size: 1.25rem;
      }
      
      &:hover {
        color: $primary;
      }
      
      &.active {
        color: $primary;
        border-bottom: 2px solid $primary;
      }
    }
  }
  
  .option-content {
    .one-time-option, .subscription-option {
      padding: 1.75rem;
      border-radius: $border-radius-lg;
      margin-bottom: 1.5rem;
      box-shadow: $box-shadow;
      transition: $transition-base;
      
      &:hover {
        box-shadow: $box-shadow-hover;
      }
      
      .price {
        margin-bottom: 1.25rem;
        
        .original-price {
          text-decoration: line-through;
          color: #64748b;
          margin-right: 0.75rem;
          font-size: 1.1rem;
        }
        
        .current-price {
          font-weight: 700;
          font-size: 1.8rem;
          color: $primary;
        }
        
        .interval {
          font-size: 1rem;
          color: #64748b;
        }
      }
      
      .description {
        margin-bottom: 1.75rem;
        
        p {
          margin-bottom: 1.25rem;
          color: #4a5568;
          line-height: 1.6;
        }
        
        .benefits {
          list-style: none;
          padding: 0;
          
          li {
            display: flex;
            align-items: flex-start;
            margin-bottom: 0.875rem;
            
            i {
              color: $primary;
              margin-right: 0.875rem;
              margin-top: 0.25rem;
              font-size: 1.1rem;
            }
            
            span {
              flex: 1;
              color: #4a5568;
            }
          }
        }
      }
      
      .actions {
        .btn {
          width: 100%;
          padding: 0.875rem;
          font-weight: 600;
          display: flex;
          align-items: center;
          justify-content: center;
          gap: 0.5rem;
          
          i {
            font-size: 1.2rem;
          }
        }
      }
    }
    
    .one-time-option {
      background-color: $light;
    }
    
    .subscription-option {
      background-color: $primary-lighter;
    }
  }
  
  .notification {
    position: absolute;
    top: -4rem;
    left: 0;
    right: 0;
    padding: 1rem 1.25rem;
    border-radius: $border-radius-lg;
    display: flex;
    align-items: center;
    gap: 0.75rem;
    animation: slideDown 0.3s ease;
    box-shadow: $box-shadow;
    
    &.success {
      background-color: lighten($success, 40%);
      border-left: 4px solid $success;
      color: darken($success, 10%);
      
      i {
        color: $success;
      }
    }
    
    &.error {
      background-color: lighten($danger, 40%);
      border-left: 4px solid $danger;
      color: darken($danger, 10%);
      
      i {
        color: $danger;
      }
    }
    
    i {
      font-size: 1.25rem;
    }
  }
}

@keyframes slideDown {
  from {
    transform: translateY(-20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>