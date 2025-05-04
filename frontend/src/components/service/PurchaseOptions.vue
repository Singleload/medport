<template>
    <div class="purchase-options">
      <div class="option-tabs">
        <div 
          class="tab" 
          :class="{ active: selectedOption === 'one-time' }" 
          @click="selectedOption = 'one-time'"
        >
          Engångsköp
        </div>
        <div 
          v-if="service.isSubscription"
          class="tab" 
          :class="{ active: selectedOption === 'subscription' }" 
          @click="selectedOption = 'subscription'"
        >
          Prenumerera
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
            <button @click="addToCart('one-time')" class="btn btn-primary">Lägg i kundvagn</button>
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
                <i class="fas fa-check"></i>
                <span>Automatisk leverans</span>
              </li>
              <li>
                <i class="fas fa-check"></i>
                <span>Förmånligt prenumerationspris</span>
              </li>
              <li>
                <i class="fas fa-check"></i>
                <span>Avsluta när du vill</span>
              </li>
            </ul>
          </div>
          <div class="actions">
            <button @click="addToCart('subscription')" class="btn btn-primary">Prenumerera nu</button>
          </div>
        </div>
      </div>
      
      <div v-if="showNotification" class="notification" :class="notificationType">
        <i v-if="notificationType === 'success'" class="fas fa-check-circle"></i>
        <i v-else class="fas fa-exclamation-circle"></i>
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
      border-bottom: 1px solid #dee2e6;
      margin-bottom: 1.5rem;
      
      .tab {
        padding: 0.75rem 1.25rem;
        cursor: pointer;
        font-weight: 500;
        color: #6c757d;
        transition: $transition-base;
        
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
        padding: 1.5rem;
        border: 1px solid #dee2e6;
        border-radius: $border-radius;
        margin-bottom: 1.5rem;
        
        .price {
          margin-bottom: 1rem;
          
          .original-price {
            text-decoration: line-through;
            color: #6c757d;
            margin-right: 0.5rem;
            font-size: 1.1rem;
          }
          
          .current-price {
            font-weight: bold;
            font-size: 1.8rem;
            color: $primary;
          }
          
          .interval {
            font-size: 1rem;
            color: #6c757d;
          }
        }
        
        .description {
          margin-bottom: 1.5rem;
          
          p {
            margin-bottom: 1rem;
          }
          
          .benefits {
            list-style: none;
            padding: 0;
            
            li {
              display: flex;
              align-items: center;
              margin-bottom: 0.75rem;
              
              i {
                color: $primary;
                margin-right: 0.75rem;
              }
            }
          }
        }
        
        .actions {
          .btn {
            width: 100%;
            padding: 0.75rem;
          }
        }
      }
      
      .subscription-option {
        background-color: #f8f9fa;
      }
    }
    
    .notification {
      position: absolute;
      top: -3.5rem;
      left: 0;
      right: 0;
      padding: 1rem;
      border-radius: $border-radius;
      display: flex;
      align-items: center;
      animation: slideDown 0.3s ease;
      
      &.success {
        background-color: lighten($success, 40%);
        border: 1px solid lighten($success, 30%);
        color: darken($success, 10%);
      }
      
      &.error {
        background-color: lighten($danger, 40%);
        border: 1px solid lighten($danger, 30%);
        color: darken($danger, 10%);
      }
      
      i {
        margin-right: 0.75rem;
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