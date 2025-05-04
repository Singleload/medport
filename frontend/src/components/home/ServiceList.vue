<template>
  <div class="service-list">
    <div class="services-grid">
      <div v-for="service in services" :key="service.id" class="service-item">
        <div class="service-card">
          <div class="service-image">
            <img :src="service.image" :alt="service.name" />
            <div v-if="service.discountedPrice" class="discount-badge">
              -{{ calculateDiscount(service.price, service.discountedPrice) }}%
            </div>
            <div v-if="service.isSubscription" class="subscription-badge">
              <i class="ri-repeat-line"></i> Prenumeration
            </div>
          </div>
          <div class="service-info">
            <h3>{{ service.name }}</h3>
            <p>{{ service.shortDescription }}</p>
            <div class="service-features">
              <div v-for="(feature, index) in getFeaturePreview(service)" :key="index" class="feature-item">
                <i class="ri-check-line"></i>
                <span>{{ feature }}</span>
              </div>
            </div>
            <div class="service-price">
              <div class="price-info">
                <span v-if="service.discountedPrice" class="original-price">{{ formatPrice(service.price) }} kr</span>
                <span class="current-price">{{ formatPrice(service.discountedPrice || service.price) }} kr</span>
                <span v-if="service.isSubscription" class="subscription-interval">{{ service.subscriptionInterval }}</span>
              </div>
              <div class="service-actions">
                <router-link :to="`/service/${service.id}`" class="btn btn-primary">
                  <i class="ri-arrow-right-line"></i>
                  <span>Läs mer</span>
                </router-link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ServiceList',
  props: {
    services: {
      type: Array,
      required: true
    }
  },
  setup() {
    const calculateDiscount = (originalPrice, discountPrice) => {
      return Math.round((1 - discountPrice / originalPrice) * 100);
    };
    
    const formatPrice = (price) => {
      return new Intl.NumberFormat('sv-SE').format(price);
    };
    
    const getFeaturePreview = (service) => {
      // Return only first 3 features for preview
      return service.features ? service.features.slice(0, 3) : [];
    };
    
    return {
      calculateDiscount,
      formatPrice,
      getFeaturePreview
    };
  }
};
</script>

<style lang="scss" scoped>
.service-list {
  .services-grid {
    display: grid;
    grid-template-columns: repeat(1, 1fr);
    gap: 2rem;
    
    @media (min-width: $breakpoint-sm) {
      grid-template-columns: repeat(2, 1fr);
    }
    
    @media (min-width: $breakpoint-lg) {
      grid-template-columns: repeat(3, 1fr);
    }
    
    .service-item {
      display: flex;
      height: 100%;
      
      .service-card {
        display: flex;
        flex-direction: column;
        width: 100%;
        background-color: $light;
        border-radius: $border-radius-lg;
        overflow: hidden;
        box-shadow: $box-shadow;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-10px);
          box-shadow: $box-shadow-hover;
          
          .service-image img {
            transform: scale(1.05);
          }
        }
        
        .service-image {
          position: relative;
          height: 220px;
          overflow: hidden;
          
          img {
            width: 100%;
            height: 100%;
            object-fit: cover;
            transition: transform 0.5s ease;
          }
          
          .discount-badge {
            position: absolute;
            top: 1rem;
            right: 1rem;
            background-color: $accent;
            color: $dark;
            padding: 0.5rem 0.75rem;
            border-radius: $border-radius-sm;
            font-weight: 700;
            font-size: 0.9rem;
            box-shadow: 0 2px 5px rgba($dark, 0.15);
            z-index: 1;
          }
          
          .subscription-badge {
            position: absolute;
            top: 1rem;
            left: 1rem;
            background-color: rgba($primary, 0.9);
            color: $light;
            padding: 0.5rem 0.75rem;
            border-radius: $border-radius-sm;
            font-weight: 500;
            font-size: 0.85rem;
            display: flex;
            align-items: center;
            gap: 0.5rem;
            box-shadow: 0 2px 5px rgba($dark, 0.15);
            z-index: 1;
          }
        }
        
        .service-info {
          padding: 1.75rem;
          display: flex;
          flex-direction: column;
          flex-grow: 1;
          
          h3 {
            margin-bottom: 0.75rem;
            font-size: 1.35rem;
            font-weight: 700;
            color: $dark;
          }
          
          p {
            color: #4a5568;
            margin-bottom: 1.5rem;
            flex-grow: 1;
          }
          
          .service-features {
            margin-bottom: 1.5rem;
            
            .feature-item {
              display: flex;
              align-items: flex-start;
              margin-bottom: 0.75rem;
              
              i {
                color: $primary;
                margin-right: 0.75rem;
                margin-top: 0.25rem;
              }
              
              span {
                flex: 1;
                color: #4a5568;
                font-size: 0.95rem;
              }
            }
          }
          
          .service-price {
            margin-top: auto;
            display: flex;
            flex-direction: column;
            gap: 1rem;
            
            .price-info {
              display: flex;
              flex-direction: column;
              
              .original-price {
                text-decoration: line-through;
                color: #64748b;
                font-size: 1rem;
              }
              
              .current-price {
                font-weight: 700;
                font-size: 1.5rem;
                color: $primary;
              }
              
              .subscription-interval {
                font-size: 0.9rem;
                color: #64748b;
                margin-top: 0.25rem;
              }
            }
            
            .service-actions {
              .btn {
                width: 100%;
                display: flex;
                align-items: center;
                justify-content: center;
                gap: 0.5rem;
                padding: 0.75rem 1rem;
                font-weight: 600;
                
                i {
                  font-size: 1.2rem;
                  transition: $transition-base;
                }
                
                &:hover i {
                  transform: translateX(3px);
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>