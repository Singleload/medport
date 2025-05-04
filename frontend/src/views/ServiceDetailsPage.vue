<template>
  <div v-if="service" class="service-details-page">
    <SEOHead
      :title="`${service.name} - Svensk Hälsovård`"
      :description="service.shortDescription"
      :schemaType="'Product'"
      :schemaData="schemaData"
    />
    
    <section class="page-header bg-primary-lighter">
      <div class="container">
        <div class="breadcrumbs">
          <router-link to="/">
            <i class="ri-home-line"></i> Hem
          </router-link>
          <span class="separator">
            <i class="ri-arrow-right-s-line"></i>
          </span>
          <router-link to="/#services">Tjänster</router-link>
          <span class="separator">
            <i class="ri-arrow-right-s-line"></i>
          </span>
          <span class="current">{{ service.name }}</span>
        </div>
      </div>
    </section>
    
    <div class="container">
      <div class="service-details">
        <div class="service-image img-decoration decoration-1">
          <img :src="service.image" :alt="service.name" />
          <div v-if="service.discountedPrice" class="discount-badge">
            <i class="ri-price-tag-3-line"></i>
            -{{ calculateDiscount(service.price, service.discountedPrice) }}%
          </div>
        </div>
        
        <div class="service-info">
          <h1>{{ service.name }}</h1>
          <p class="service-short-description">{{ service.shortDescription }}</p>
          
          <div class="service-highlights">
            <div class="highlight">
              <div class="highlight-icon">
                <i class="ri-shield-check-line"></i>
              </div>
              <div class="highlight-text">
                Kvalitetssäkrade analyser
              </div>
            </div>
            <div class="highlight">
              <div class="highlight-icon">
                <i class="ri-time-line"></i>
              </div>
              <div class="highlight-text">
                Snabba resultat
              </div>
            </div>
            <div class="highlight">
              <div class="highlight-icon">
                <i class="ri-user-heart-line"></i>
              </div>
              <div class="highlight-text">
                Erfarna specialister
              </div>
            </div>
          </div>
          
          <div class="service-price">
            <div class="price-container">
              <span v-if="service.discountedPrice" class="original-price">{{ formatPrice(service.price) }} kr</span>
              <span class="current-price">{{ formatPrice(service.discountedPrice || service.price) }} kr</span>
              <span v-if="service.isSubscription" class="subscription-interval">{{ service.subscriptionInterval }}</span>
            </div>
            
            <div class="service-rating">
              <div class="stars">
                <i class="ri-star-fill"></i>
                <i class="ri-star-fill"></i>
                <i class="ri-star-fill"></i>
                <i class="ri-star-fill"></i>
                <i class="ri-star-fill"></i>
              </div>
              <span>(24 recensioner)</span>
            </div>
          </div>
          
          <div class="purchase-options">
            <PurchaseOptions :service="service" />
          </div>
        </div>
      </div>
      
      <div class="service-description-section">
        <div class="service-description">
          <div class="description-tabs">
            <div 
              class="tab" 
              :class="{ active: activeTab === 'description' }" 
              @click="activeTab = 'description'"
            >
              <i class="ri-information-line"></i> Beskrivning
            </div>
            <div 
              class="tab" 
              :class="{ active: activeTab === 'features' }" 
              @click="activeTab = 'features'"
            >
              <i class="ri-list-check"></i> Inkluderar
            </div>
            <div 
              class="tab" 
              :class="{ active: activeTab === 'faq' }" 
              @click="activeTab = 'faq'"
            >
              <i class="ri-question-line"></i> Vanliga frågor
            </div>
          </div>
          
          <div class="tab-content">
            <div v-if="activeTab === 'description'" class="description-content">
              <p>{{ service.description }}</p>
              <p>Våra tjänster utförs av erfarna sjuksköterskor och läkare med hög kompetens inom området. Vi använder de senaste metoderna och teknologierna för att säkerställa högsta kvalitet på våra analyser.</p>
              <p>Efter genomförd undersökning eller prov får du en detaljerad rapport med dina resultat och rekommendationer från våra medicinska experter.</p>
              
              <div class="process-steps">
                <h3>Så här går det till</h3>
                <div class="steps">
                  <div class="step">
                    <div class="step-number">1</div>
                    <div class="step-icon">
                      <i class="ri-shopping-cart-2-line"></i>
                    </div>
                    <h4>Beställ</h4>
                    <p>Välj tjänst och genomför beställningen via vår säkra betalningslösning.</p>
                  </div>
                  
                  <div class="step">
                    <div class="step-number">2</div>
                    <div class="step-icon">
                      <i class="ri-calendar-check-line"></i>
                    </div>
                    <h4>Boka tid</h4>
                    <p>Efter beställning kontaktar vi dig för att boka en tid som passar dig.</p>
                  </div>
                  
                  <div class="step">
                    <div class="step-number">3</div>
                    <div class="step-icon">
                      <i class="ri-test-tube-line"></i>
                    </div>
                    <h4>Provtagning</h4>
                    <p>Besök någon av våra kliniker för provtagning eller undersökning.</p>
                  </div>
                  
                  <div class="step">
                    <div class="step-number">4</div>
                    <div class="step-icon">
                      <i class="ri-file-chart-line"></i>
                    </div>
                    <h4>Resultat</h4>
                    <p>Ta del av dina resultat digitalt och få personlig uppföljning.</p>
                  </div>
                </div>
              </div>
            </div>
            
            <div v-if="activeTab === 'features'" class="features-content">
              <ul class="feature-list">
                <li v-for="(feature, index) in service.features" :key="index">
                  <i class="ri-check-line"></i>
                  <span>{{ feature }}</span>
                </li>
              </ul>
              
              <div class="feature-highlight">
                <div class="highlight-icon">
                  <i class="ri-customer-service-2-line"></i>
                </div>
                <div class="highlight-content">
                  <h3>Professionell service</h3>
                  <p>Alla våra prover och undersökningar utförs av legitimerad sjukvårdspersonal med lång erfarenhet. Vi står för kvalitet och pålitlighet.</p>
                </div>
              </div>
            </div>
            
            <div v-if="activeTab === 'faq'" class="faq-content">
              <div class="faq-item">
                <div class="faq-question">
                  <i class="ri-question-line"></i> Hur går processen till?
                </div>
                <div class="faq-answer">
                  <p>Efter att du har bokat och betalat för din tjänst kommer du att få en bekräftelse via e-post med information om din bokning. Du kommer att kontaktas för att bekräfta en tid för ditt besök, eller för att få instruktioner för självprovtagning om det är tillämpligt.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">
                  <i class="ri-question-line"></i> Hur lång tid tar testet/undersökningen?
                </div>
                <div class="faq-answer">
                  <p>Tiden varierar beroende på vilken tjänst du har valt. En standard blodprovstagning tar normalt 15-20 minuter, medan en fullständig hälsokontroll kan ta upp till 1-2 timmar.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">
                  <i class="ri-question-line"></i> När får jag mina resultat?
                </div>
                <div class="faq-answer">
                  <p>Resultaten från blodprover är normalt tillgängliga inom 2-3 arbetsdagar. För mer omfattande hälsokontroller kan det ta upp till 5 arbetsdagar innan alla resultat är sammanställda.</p>
                </div>
              </div>
              <div class="faq-item">
                <div class="faq-question">
                  <i class="ri-question-line"></i> Kan jag avboka eller ändra min bokning?
                </div>
                <div class="faq-answer">
                  <p>Du kan avboka eller ändra din bokning upp till 24 timmar före din bokade tid utan kostnad. För avbokningar med kortare varsel debiteras en avgift på 50% av tjänstens pris.</p>
                </div>
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
                  <i class="ri-price-tag-3-line"></i>
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
                  <router-link :to="`/service/${relatedService.id}`" class="btn btn-primary">
                    <i class="ri-arrow-right-line"></i>
                    <span>Läs mer</span>
                  </router-link>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="service-cta">
        <div class="cta-content">
          <div class="cta-icon">
            <i class="ri-chat-smile-3-line"></i>
          </div>
          <div class="cta-text">
            <h3>Har du frågor om våra tjänster?</h3>
            <p>Vi hjälper dig gärna att hitta den tjänst som passar dig bäst.</p>
          </div>
          <div class="cta-actions">
            <router-link to="/contact" class="btn btn-primary">
              <i class="ri-customer-service-line"></i>
              Kontakta oss
            </router-link>
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
  .page-header {
    padding: 2rem 0;
    background-color: $primary-lighter;
    margin-bottom: 3rem;
  }
  
  .breadcrumbs {
    display: flex;
    align-items: center;
    font-size: 0.95rem;
    
    a {
      color: $primary;
      display: flex;
      align-items: center;
      transition: $transition-base;
      
      &:hover {
        color: darken($primary, 10%);
      }
      
      i {
        margin-right: 0.3rem;
      }
    }
    
    .separator {
      margin: 0 0.5rem;
      color: #64748b;
    }
    
    .current {
      color: #64748b;
      font-weight: 500;
    }
  }
  
  .service-details {
    display: flex;
    flex-wrap: wrap;
    gap: 3rem;
    margin-bottom: 4rem;
    
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
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        transition: $transition-base;
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
        display: flex;
        align-items: center;
        gap: 0.5rem;
        box-shadow: 0 4px 8px rgba($dark, 0.1);
        
        i {
          font-size: 1.1rem;
        }
      }
    }
    
    .service-info {
      flex: 0 0 100%;
      max-width: 100%;
      
      @media (min-width: $breakpoint-md) {
        flex: 0 0 calc(60% - 3rem);
        max-width: calc(60% - 3rem);
      }
      
      h1 {
        color: $primary;
        margin-bottom: 1rem;
        font-weight: 800;
        font-size: 2.5rem;
        
        @media (max-width: $breakpoint-md) {
          font-size: 2rem;
        }
      }
      
      .service-short-description {
        font-size: 1.2rem;
        margin-bottom: 1.75rem;
        color: #4a5568;
        line-height: 1.6;
      }
      
      .service-highlights {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        margin-bottom: 2rem;
        
        .highlight {
          flex: 1;
          min-width: 120px;
          display: flex;
          align-items: center;
          gap: 0.75rem;
          background-color: $primary-lighter;
          padding: 0.75rem 1rem;
          border-radius: $border-radius;
          
          .highlight-icon {
            color: $primary;
            font-size: 1.5rem;
            line-height: 1;
          }
          
          .highlight-text {
            font-weight: 500;
            font-size: 0.9rem;
          }
        }
      }
      
      .service-price {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        align-items: flex-end;
        margin-bottom: 2rem;
        
        .price-container {
          display: flex;
          flex-wrap: wrap;
          align-items: baseline;
          
          .original-price {
            text-decoration: line-through;
            color: #64748b;
            margin-right: 0.75rem;
            font-size: 1.1rem;
          }
          
          .current-price {
            font-weight: 800;
            font-size: 2.2rem;
            color: $primary;
            
            @media (max-width: $breakpoint-md) {
              font-size: 1.8rem;
            }
          }
          
          .subscription-interval {
            display: block;
            width: 100%;
            font-size: 0.95rem;
            color: #64748b;
            margin-top: 0.375rem;
          }
        }
        
        .service-rating {
          display: flex;
          flex-direction: column;
          align-items: flex-end;
          
          .stars {
            color: $accent;
            margin-bottom: 0.25rem;
            font-size: 1.2rem;
          }
          
          span {
            color: #64748b;
            font-size: 0.9rem;
          }
        }
      }
      
      .purchase-options {
        margin-top: 2rem;
      }
    }
  }
  
  .service-description-section {
    margin-bottom: 4rem;
  }
  
  .service-description {
    background-color: $light;
    border-radius: $border-radius-lg;
    box-shadow: $box-shadow;
    overflow: hidden;
    
    .description-tabs {
      display: flex;
      border-bottom: 1px solid #e2e8f0;
      
      .tab {
        padding: 1.25rem 1.5rem;
        cursor: pointer;
        font-weight: 600;
        color: #64748b;
        transition: $transition-base;
        display: flex;
        align-items: center;
        gap: 0.5rem;
        
        i {
          font-size: 1.2rem;
        }
        
        &:hover {
          color: $primary;
          background-color: rgba($primary, 0.05);
        }
        
        &.active {
          color: $primary;
          border-bottom: 2px solid $primary;
        }
      }
    }
    
    .tab-content {
      padding: 2rem;
      
      .description-content {
        p {
          margin-bottom: 1.5rem;
          line-height: 1.7;
          color: #4a5568;
        }
        
        .process-steps {
          margin-top: 3rem;
          
          h3 {
            text-align: center;
            color: $primary;
            margin-bottom: 2rem;
            font-size: 1.5rem;
          }
          
          .steps {
            display: grid;
            grid-template-columns: repeat(2, 1fr);
            gap: 2rem;
            
            @media (min-width: $breakpoint-lg) {
              grid-template-columns: repeat(4, 1fr);
            }
            
            @media (max-width: $breakpoint-sm) {
              grid-template-columns: 1fr;
            }
            
            .step {
              position: relative;
              text-align: center;
              padding: 2.5rem 1rem 1.5rem;
              background-color: $primary-lighter;
              border-radius: $border-radius-lg;
              
              .step-number {
                position: absolute;
                top: -15px;
                left: 50%;
                transform: translateX(-50%);
                width: 30px;
                height: 30px;
                background-color: $accent;
                color: $dark;
                font-weight: 700;
                border-radius: 50%;
                display: flex;
                align-items: center;
                justify-content: center;
                box-shadow: 0 4px 6px rgba($dark, 0.15);
              }
              
              .step-icon {
                width: 60px;
                height: 60px;
                margin: 0 auto 1rem;
                background-color: white;
                color: $primary;
                font-size: 1.75rem;
                border-radius: 50%;
                display: flex;
                align-items: center;
                justify-content: center;
                box-shadow: 0 4px 6px rgba($primary, 0.1);
              }
              
              h4 {
                margin-bottom: 0.75rem;
                font-size: 1.1rem;
                color: $primary;
              }
              
              p {
                margin-bottom: 0;
                font-size: 0.95rem;
                color: #4a5568;
              }
            }
          }
        }
      }
      
      .features-content {
        .feature-list {
          list-style: none;
          padding: 0;
          margin-bottom: 3rem;
          
          li {
            margin-bottom: 1.25rem;
            display: flex;
            align-items: flex-start;
            padding-bottom: 1.25rem;
            border-bottom: 1px dashed #e2e8f0;
            
            &:last-child {
              border-bottom: none;
            }
            
            i {
              flex-shrink: 0;
              width: 26px;
              height: 26px;
              background-color: $primary-lighter;
              color: $primary;
              border-radius: 50%;
              display: flex;
              align-items: center;
              justify-content: center;
              margin-right: 1rem;
              margin-top: 0.125rem;
            }
            
            span {
              flex: 1;
              color: #4a5568;
            }
          }
        }
        
        .feature-highlight {
          background-color: $primary-lighter;
          padding: 2rem;
          border-radius: $border-radius-lg;
          display: flex;
          align-items: flex-start;
          gap: 1.5rem;
          
          .highlight-icon {
            flex-shrink: 0;
            width: 60px;
            height: 60px;
            background-color: $primary;
            color: white;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 1.75rem;
          }
          
          .highlight-content {
            flex: 1;
            
            h3 {
              color: $primary;
              margin-bottom: 0.75rem;
              font-size: 1.3rem;
            }
            
            p {
              margin: 0;
              color: #4a5568;
            }
          }
        }
      }
      
      .faq-content {
        .faq-item {
          margin-bottom: 1.75rem;
          padding-bottom: 1.75rem;
          border-bottom: 1px solid #e2e8f0;
          
          &:last-child {
            margin-bottom: 0;
            padding-bottom: 0;
            border-bottom: none;
          }
          
          .faq-question {
            font-weight: 700;
            margin-bottom: 0.75rem;
            color: $primary;
            display: flex;
            align-items: center;
            gap: 0.75rem;
            
            i {
              background-color: $primary-lighter;
              color: $primary;
              width: 30px;
              height: 30px;
              border-radius: 50%;
              display: flex;
              align-items: center;
              justify-content: center;
              flex-shrink: 0;
            }
          }
          
          .faq-answer {
            color: #4a5568;
            line-height: 1.7;
            padding-left: calc(30px + 0.75rem);
            
            p {
              margin: 0;
            }
          }
        }
      }
    }
  }
  
  .related-services {
    margin-bottom: 4rem;
    
    h2 {
      color: $primary;
      margin-bottom: 2.5rem;
      font-weight: 700;
      text-align: center;
      position: relative;
      padding-bottom: 1rem;
      
      &::after {
        content: '';
        position: absolute;
        bottom: 0;
        left: 50%;
        transform: translateX(-50%);
        width: 80px;
        height: 3px;
        background-color: $primary-light;
      }
    }
  }
  
  .service-cta {
    margin-bottom: 4rem;
    
    .cta-content {
      background-color: $primary-lighter;
      border-radius: $border-radius-lg;
      padding: 2.5rem;
      display: flex;
      align-items: center;
      gap: 2rem;
      
      @media (max-width: $breakpoint-md) {
        flex-direction: column;
        text-align: center;
      }
      
      .cta-icon {
        width: 80px;
        height: 80px;
        background-color: $light;
        color: $primary;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 2.5rem;
        flex-shrink: 0;
        box-shadow: $box-shadow;
      }
      
      .cta-text {
        flex: 1;
        
        h3 {
          color: $primary;
          margin-bottom: 0.5rem;
          font-size: 1.5rem;
          
          @media (max-width: $breakpoint-md) {
            font-size: 1.3rem;
          }
        }
        
        p {
          margin: 0;
          color: #4a5568;
          font-size: 1.1rem;
          
          @media (max-width: $breakpoint-md) {
            font-size: 1rem;
          }
        }
      }
      
      .cta-actions {
        flex-shrink: 0;
        
        .btn {
          padding: 0.75rem 1.5rem;
          font-weight: 600;
          display: flex;
          align-items: center;
          gap: 0.5rem;
          
          i {
            font-size: 1.2rem;
          }
        }
      }
    }
  }
}
</style>