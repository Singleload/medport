<template>
    <div class="service-list">
      <div class="row">
        <div v-for="service in services" :key="service.id" class="service-item">
          <div class="service-card">
            <div class="service-image">
              <img :src="service.image" :alt="service.name" />
              <div v-if="service.discountedPrice" class="discount-badge">
                -{{ calculateDiscount(service.price, service.discountedPrice) }}%
              </div>
            </div>
            <div class="service-info">
              <h3>{{ service.name }}</h3>
              <p>{{ service.shortDescription }}</p>
              <div class="service-price">
                <span v-if="service.discountedPrice" class="original-price">{{ formatPrice(service.price) }} kr</span>
                <span class="current-price">{{ formatPrice(service.discountedPrice || service.price) }} kr</span>
                <span v-if="service.isSubscription" class="subscription-interval">{{ service.subscriptionInterval }}</span>
              </div>
              <div class="service-actions">
                <router-link :to="`/service/${service.id}`" class="btn btn-primary">Läs mer</router-link>
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
      
      return {
        calculateDiscount,
        formatPrice
      };
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .service-list {
    .row {
      display: flex;
      flex-wrap: wrap;
      margin: 0 -1rem;
    }
    
    .service-item {
      flex: 0 0 100%;
      max-width: 100%;
      padding: 0 1rem;
      margin-bottom: 2rem;
      
      @media (min-width: $breakpoint-sm) {
        flex: 0 0 50%;
        max-width: 50%;
      }
      
      @media (min-width: $breakpoint-lg) {
        flex: 0 0 33.333333%;
        max-width: 33.333333%;
      }
    }
    
    .service-card {
      height: 100%;
      border-radius: $border-radius;
      overflow: hidden;
      box-shadow: $box-shadow-sm;
      background-color: white;
      transition: transform 0.3s ease, box-shadow 0.3s ease;
      
      &:hover {
        transform: translateY(-5px);
        box-shadow: $box-shadow;
      }
      
      .service-image {
        position: relative;
        height: 200px;
        overflow: hidden;
        
        img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
        
        .discount-badge {
          position: absolute;
          top: 10px;
          right: 10px;
          background-color: $primary;
          color: white;
          padding: 0.25rem 0.5rem;
          border-radius: 0.25rem;
          font-weight: bold;
          font-size: 0.9rem;
        }
      }
      
      .service-info {
        padding: 1.5rem;
        
        h3 {
          margin-bottom: 0.5rem;
          font-size: 1.25rem;
        }
        
        p {
          color: #6c757d;
          margin-bottom: 1rem;
        }
        
        .service-price {
          margin-bottom: 1.5rem;
          
          .original-price {
            text-decoration: line-through;
            color: #6c757d;
            margin-right: 0.5rem;
            font-size: 0.9rem;
          }
          
          .current-price {
            font-weight: bold;
            font-size: 1.2rem;
            color: $primary;
          }
          
          .subscription-interval {
            display: block;
            font-size: 0.9rem;
            color: #6c757d;
            margin-top: 0.25rem;
          }
        }
        
        .service-actions {
          display: flex;
          justify-content: center;
        }
      }
    }
  }
  </style>