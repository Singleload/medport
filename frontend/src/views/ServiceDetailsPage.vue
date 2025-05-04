<template>
    <div v-if="service" class="service-details-page">
      <SEOHead
        :title="`${service.name} - Svensk Hälsovård`"
        :description="service.shortDescription"
        :schemaType="'Product'"
        :schemaData="schemaData"
      />
      
      <div class="container">
        <div class="breadcrumbs">
          <router-link to="/">Hem</router-link>
          <span class="separator">/</span>
          <router-link to="/#services">Tjänster</router-link>
          <span class="separator">/</span>
          <span class="current">{{ service.name }}</span>
        </div>
        
        <div class="service-details">
          <div class="service-image">
            <img :src="service.image" :alt="service.name" />
            <div v-if="service.discountedPrice" class="discount-badge">
              -{{ calculateDiscount(service.price, service.discountedPrice) }}%
            </div>
          </div>
          
          <div class="service-info">
            <h1>{{ service.name }}</h1>
            <p class="service-short-description">{{ service.shortDescription }}</p>
            
            <div class="service-price">
              <div class="price-container">
                <span v-if="service.discountedPrice" class="original-price">{{ formatPrice(service.price) }} kr</span>
                <span class="current-price">{{ formatPrice(service.discountedPrice || service.price) }} kr</span>
                <span v-if="service.isSubscription" class="subscription-interval">{{ service.subscriptionInterval }}</span>
              </div>
              
              <div class="service-rating">
                <div class="stars">
                  <i class="fas fa-star"></i>
                  <i class="fas fa-star"></i>
                  <i class="fas fa-star"></i>
                  <i class="fas fa-star"></i>
                  <i class="fas fa-star"></i>
                </div>
                <span>(24 recensioner)</span>
              </div>
            </div>
            
            <div class="purchase-options">
              <PurchaseOptions :service="service" />
            </div>
          </div>
        </div>
        
        <div class="service-description">
          <div class="description-tabs">
            <div 
              class="tab" 
              :class="{ active: activeTab === 'description' }" 
              @click="activeTab = 'description'"
            >
              Beskrivning
            </div>
            <div 
              class="tab" 
              :class="{ active: activeTab === 'features' }" 
              @click="activeTab = 'features'"
            >
              Inkluderar
            </div>
            <div 
              class="tab" 
              :class="{ active: activeTab === 'faq' }" 
              @click="activeTab = 'faq'"
            >
              Vanliga frågor
            </div>
          </div>
          
          <div class="tab-content">
            <div v-if="activeTab === 'description'" class="description-content">
              <p>{{ service.description }}</p>
              <p>Våra tjänster utförs av erfarna sjuksköterskor och läkare med hög kompetens inom området. Vi använder de senaste metoderna och teknologierna för att säkerställa högsta kvalitet på våra analyser.</p>
              <p>Efter genomförd undersökning eller prov får du en detaljerad rapport med dina resultat och rekommendationer från våra medicinska experter.</p>
            </div>
            
            <div v-if="activeTab === 'features'" class="features-content">
              <ul class="feature-list">
                <li v-for="(feature, index) in service.features" :key="index">
                  <i class="fas fa-check"></i>
                  <span>{{ feature }}</span>
                </li>
              </ul>
            </div>
            
            <div v-if="activeTab === 'faq'" class="faq-content">
              <div class="faq-item">
                <div class="faq-question">Hur går processen till?</div>
                <div class="faq-answer">
                  <p>Efter att du har bokat och betalat för din tjänst kommer du att få en bekräftelse via e-post med information om din bokning. Du kommer att kontaktas för att bekräfta en tid för ditt besök, eller för att få instruktioner för självprovtagning om det är tillämpligt.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">Hur lång tid tar testet/undersökningen?</div>
                <div class="faq-answer">
                  <p>Tiden varierar beroende på vilken tjänst du har valt. En standard blodprovstagning tar normalt 15-20 minuter, medan en fullständig hälsokontroll kan ta upp till 1-2 timmar.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">När får jag mina resultat?</div>
                <div class="faq-answer">
                  <p>Resultaten från blodprover är normalt tillgängliga inom 2-3 arbetsdagar. För mer omfattande hälsokontroller kan det ta upp till 5 arbetsdagar innan alla resultat är sammanställda.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">Kan jag avboka eller ändra min bokning?</div>
                <div class="faq-answer">
                  <p>Du kan avboka eller ändra din bokning upp till 24 timmar före din bokade tid utan kostnad. För avbokningar med kortare varsel debiteras en avgift på 50% av tjänstens pris.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="related-services">
          <h2>Relaterade tjänster</h2>
          <div class="related-services-list">
            <div 
              v-for="relatedService in relatedServices" 
              :key="relatedService.id" 
              class="related-service-item"
            >
              <div class="service-card">
                <div class="service-image">
                  <img :src="relatedService.image" :alt="relatedService.name" />
                  <div v-if="relatedService.discountedPrice" class="discount-badge">
                    -{{ calculateDiscount(relatedService.price, relatedService.discountedPrice) }}%
                  </div>
                </div>
                <div class="service-info">
                  <h3>{{ relatedService.name }}</h3>
                  <p>{{ relatedService.shortDescription }}</p>
                  <div class="service-price">
                    <span v-if="relatedService.discountedPrice" class="original-price">{{ formatPrice(relatedService.price) }} kr</span>
                    <span class="current-price">{{ formatPrice(relatedService.discountedPrice || relatedService.price) }} kr</span>
                  </div>
                  <div class="service-actions">
                    <router-link :to="`/service/${relatedService.id}`" class="btn btn-primary">Läs mer</router-link>
                  </div>
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
  import { useRoute } from 'vue-router';
  import { useStore } from 'vuex';
  import SEOHead from '@/components/common/SEOHead.vue';
  import PurchaseOptions from '@/components/service/PurchaseOptions.vue';
  
  export default {
    name: 'ServiceDetailsPage',
    components: {
      SEOHead,
      PurchaseOptions
    },
    setup() {
      const route = useRoute();
      const store = useStore();
      
      const activeTab = ref('description');
      
      const service = computed(() => {
        return store.getters.getServiceById(route.params.id);
      });
      
      const relatedServices = computed(() => {
        if (!service.value) return [];
        
        // Get all services except the current one
        const otherServices = store.getters.services.filter(s => s.id !== service.value.id);
        
        // Return up to 3 related services
        return otherServices.slice(0, 3);
      });
      
      const calculateDiscount = (originalPrice, discountPrice) => {
        return Math.round((1 - discountPrice / originalPrice) * 100);
      };
      
      const formatPrice = (price) => {
        return new Intl.NumberFormat('sv-SE').format(price);
      };
      
      const schemaData = computed(() => {
        if (!service.value) return {};
        
        return {
          name: service.value.name,
          description: service.value.description,
          image: service.value.image,
          offers: {
            '@type': 'Offer',
            price: service.value.discountedPrice || service.value.price,
            priceCurrency: 'SEK',
            availability: 'https://schema.org/InStock'
          },
          brand: {
            '@type': 'Brand',
            name: 'Svensk Hälsovård'
          }
        };
      });
      
      // Reset active tab when service changes
      watch(() => route.params.id, () => {
        activeTab.value = 'description';
        window.scrollTo(0, 0);
      });
      
      onMounted(() => {
        window.scrollTo(0, 0);
      });
      
      return {
        service,
        relatedServices,
        activeTab,
        calculateDiscount,
        formatPrice,
        schemaData
      };
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .service-details-page {
    padding: 2rem 0;
    
    .breadcrumbs {
      display: flex;
      align-items: center;
      margin-bottom: 2rem;
      font-size: 0.9rem;
      
      a {
        color: #6c757d;
        
        &:hover {
          color: $primary;
        }
      }
      
      .separator {
        margin: 0 0.5rem;
        color: #adb5bd;
      }
      
      .current {
        color: $primary;
        font-weight: 500;
      }
    }
    
    .service-details {
      display: flex;
      flex-wrap: wrap;
      gap: 2rem;
      margin-bottom: 3rem;
      
      .service-image {
        position: relative;
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-md) {
          flex: 0 0 40%;
          max-width: 40%;
        }
        
        img {
          width: 100%;
          border-radius: $border-radius;
          box-shadow: $box-shadow-sm;
        }
        
        .discount-badge {
          position: absolute;
          top: 1rem;
          right: 1rem;
          background-color: $primary;
          color: white;
          padding: 0.5rem 0.75rem;
          border-radius: 0.25rem;
          font-weight: bold;
        }
      }
      
      .service-info {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-md) {
          flex: 0 0 calc(60% - 2rem);
          max-width: calc(60% - 2rem);
        }
        
        h1 {
          color: $primary;
          margin-bottom: 1rem;
        }
        
        .service-short-description {
          font-size: 1.1rem;
          margin-bottom: 1.5rem;
          color: #495057;
        }
        
        .service-price {
          display: flex;
          flex-wrap: wrap;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 1.5rem;
          
          .price-container {
            display: flex;
            flex-wrap: wrap;
            align-items: baseline;
            
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
            
            .subscription-interval {
              display: block;
              width: 100%;
              font-size: 0.9rem;
              color: #6c757d;
              margin-top: 0.25rem;
            }
          }
          
          .service-rating {
            display: flex;
            align-items: center;
            
            .stars {
              color: #ffc107;
              margin-right: 0.5rem;
            }
            
            span {
              color: #6c757d;
              font-size: 0.9rem;
            }
          }
        }
      }
    }
    
    .service-description {
      margin-bottom: 3rem;
      
      .description-tabs {
        display: flex;
        border-bottom: 1px solid #dee2e6;
        margin-bottom: 2rem;
        
        .tab {
          padding: 1rem 1.5rem;
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
      
      .tab-content {
        .description-content {
          p {
            margin-bottom: 1rem;
            line-height: 1.6;
          }
        }
        
        .features-content {
          .feature-list {
            list-style: none;
            padding: 0;
            
            li {
              margin-bottom: 1rem;
              display: flex;
              align-items: flex-start;
              
              i {
                color: $primary;
                margin-right: 0.75rem;
                margin-top: 0.25rem;
              }
            }
          }
        }
        
        .faq-content {
          .faq-item {
            margin-bottom: 1.5rem;
            
            .faq-question {
              font-weight: 500;
              margin-bottom: 0.5rem;
              color: $primary;
            }
            
            .faq-answer {
              color: #495057;
            }
          }
        }
      }
    }
    
    .related-services {
      h2 {
        color: $primary;
        margin-bottom: 2rem;
      }
      
      .related-services-list {
        display: flex;
        flex-wrap: wrap;
        gap: 2rem;
        
        .related-service-item {
          flex: 0 0 100%;
          max-width: 100%;
          
          @media (min-width: $breakpoint-sm) {
            flex: 0 0 calc(50% - 1rem);
            max-width: calc(50% - 1rem);
          }
          
          @media (min-width: $breakpoint-lg) {
            flex: 0 0 calc(33.333333% - 1.33rem);
            max-width: calc(33.333333% - 1.33rem);
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
              height: 180px;
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
                font-size: 1.2rem;
              }
              
              p {
                color: #6c757d;
                margin-bottom: 1rem;
                font-size: 0.9rem;
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
              }
              
              .service-actions {
                display: flex;
                justify-content: center;
              }
            }
          }
        }
      }
    }
  }
  </style>