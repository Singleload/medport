<template>
    <div class="home-page">
      <SEOHead
        title="Svensk Hälsovård - Din väg till bättre hälsa"
        description="Boka blodprov och hälsokontroller enkelt och snabbt. Vi erbjuder professionella hälsotjänster för alla."
        :schemaType="'HealthAndBeautyBusiness'"
        :schemaData="schemaData"
      />
      
      <section class="hero">
        <div class="container">
          <div class="hero-content">
            <h1>Din väg till bättre hälsa</h1>
            <p>Vi erbjuder professionella hälsokontroller och blodprover för att hjälpa dig ta kontroll över din hälsa.</p>
            <div class="hero-cta">
              <router-link to="#services" class="btn btn-primary">Se våra tjänster</router-link>
            </div>
          </div>
        </div>
      </section>
      
      <section id="about" class="about-section">
        <div class="container">
          <div class="about-content">
            <h2>Om Svensk Hälsovård</h2>
            <p>Svensk Hälsovård är en ledande leverantör av hälsokontroller och blodprovsanalyser. Vi erbjuder både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa.</p>
            <p>Vårt erfarna team av läkare och sjuksköterskor garanterar högkvalitativ service och professionell rådgivning.</p>
            <div class="about-features">
              <div class="feature">
                <div class="feature-icon">
                  <i class="fas fa-certificate"></i>
                </div>
                <h3>Kvalitetssäkrad</h3>
                <p>Våra laboratoriepartners är ackrediterade enligt svenska kvalitetsstandarder.</p>
              </div>
              <div class="feature">
                <div class="feature-icon">
                  <i class="fas fa-lock"></i>
                </div>
                <h3>Säker hantering</h3>
                <p>Din hälsodata hanteras med högsta sekretess och säkerhet.</p>
              </div>
              <div class="feature">
                <div class="feature-icon">
                  <i class="fas fa-user-md"></i>
                </div>
                <h3>Expertis</h3>
                <p>Vårt team består av erfarna läkare och sjuksköterskor.</p>
              </div>
            </div>
          </div>
        </div>
      </section>
      
      <section id="services" class="services-section">
        <div class="container">
          <h2>Våra tjänster</h2>
          <ServiceList :services="services" />
        </div>
      </section>
      
      <section class="testimonials-section">
        <div class="container">
          <h2>Vad våra kunder säger</h2>
          <div class="testimonials">
            <div class="testimonial">
              <div class="testimonial-content">
                <p>"Smidigt och professionellt! Jag fick en grundlig genomgång av mina resultat och känner mig trygg med min hälsa."</p>
              </div>
              <div class="testimonial-author">
                <p><strong>Anna Johansson</strong></p>
                <p>Hälsokontroll - Kvinna</p>
              </div>
            </div>
            <div class="testimonial">
              <div class="testimonial-content">
                <p>"Prenumerationstjänsten är perfekt för mig som vill hålla koll på mina värden regelbundet. Rekommenderas!"</p>
              </div>
              <div class="testimonial-author">
                <p><strong>Erik Lindberg</strong></p>
                <p>Blodprov - Prenumeration</p>
              </div>
            </div>
            <div class="testimonial">
              <div class="testimonial-content">
                <p>"Enkelt att boka och betala. Fick mina resultat snabbt och med tydliga förklaringar."</p>
              </div>
              <div class="testimonial-author">
                <p><strong>Maria Svensson</strong></p>
                <p>Blodprov - Premium</p>
              </div>
            </div>
          </div>
        </div>
      </section>
      
      <section class="cta-section">
        <div class="container">
          <div class="cta-content">
            <h2>Ta kontroll över din hälsa idag</h2>
            <p>Boka en hälsokontroll eller blodprov och investera i din framtida hälsa.</p>
            <router-link to="#services" class="btn btn-primary">Se våra tjänster</router-link>
          </div>
        </div>
      </section>
    </div>
  </template>
  
  <script>
  import { computed } from 'vue';
  import { useStore } from 'vuex';
  import SEOHead from '@/components/common/SEOHead.vue';
  import ServiceList from '@/components/home/ServiceList.vue';
  
  export default {
    name: 'HomePage',
    components: {
      SEOHead,
      ServiceList
    },
    setup() {
      const store = useStore();
      
      const services = computed(() => store.getters.services);
      
      const schemaData = {
        address: {
          '@type': 'PostalAddress',
          streetAddress: 'Hälsogatan 123',
          addressLocality: 'Stockholm',
          postalCode: '123 45',
          addressCountry: 'SE'
        },
        telephone: '+46801234567',
        email: 'info@svenskhalsovard.se',
        openingHoursSpecification: [
          {
            '@type': 'OpeningHoursSpecification',
            dayOfWeek: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
            opens: '08:00',
            closes: '17:00'
          }
        ],
        priceRange: '995 kr - 3495 kr'
      };
      
      return {
        services,
        schemaData
      };
    }
  };
  </script>
  
  <style lang="scss" scoped>
  .home-page {
    .hero {
      background-color: #f8f9fa;
      background-image: linear-gradient(rgba(255, 255, 255, 0.8), rgba(255, 255, 255, 0.8)), url('/assets/images/hero-bg.jpg');
      background-size: cover;
      background-position: center;
      padding: 5rem 0;
      margin-bottom: 3rem;
      
      .hero-content {
        max-width: 600px;
        
        h1 {
          font-size: 2.8rem;
          margin-bottom: 1.5rem;
          color: $primary;
        }
        
        p {
          font-size: 1.2rem;
          margin-bottom: 2rem;
        }
        
        .hero-cta {
          .btn {
            padding: 0.75rem 2rem;
            font-size: 1.1rem;
          }
        }
      }
    }
    
    .about-section {
      padding: 4rem 0;
      
      .about-content {
        max-width: 800px;
        margin: 0 auto;
        text-align: center;
        
        h2 {
          color: $primary;
          margin-bottom: 1.5rem;
        }
        
        p {
          font-size: 1.1rem;
          margin-bottom: 1.5rem;
        }
        
        .about-features {
          display: flex;
          flex-wrap: wrap;
          justify-content: center;
          gap: 2rem;
          margin-top: 3rem;
          
          .feature {
            flex: 0 0 100%;
            max-width: 100%;
            
            @media (min-width: $breakpoint-sm) {
              flex: 0 0 calc(33.333333% - 2rem);
              max-width: calc(33.333333% - 2rem);
            }
            
            .feature-icon {
              font-size: 2.5rem;
              color: $primary;
              margin-bottom: 1rem;
            }
            
            h3 {
              margin-bottom: 0.5rem;
            }
            
            p {
              font-size: 1rem;
            }
          }
        }
      }
    }
    
    .services-section {
      padding: 4rem 0;
      background-color: #f8f9fa;
      
      h2 {
        text-align: center;
        color: $primary;
        margin-bottom: 2rem;
      }
    }
    
    .testimonials-section {
      padding: 4rem 0;
      
      h2 {
        text-align: center;
        color: $primary;
        margin-bottom: 2rem;
      }
      
      .testimonials {
        display: flex;
        flex-wrap: wrap;
        gap: 2rem;
        
        .testimonial {
          flex: 0 0 100%;
          max-width: 100%;
          background-color: #f8f9fa;
          border-radius: $border-radius;
          padding: 1.5rem;
          
          @media (min-width: $breakpoint-md) {
            flex: 0 0 calc(33.333333% - 2rem);
            max-width: calc(33.333333% - 2rem);
          }
          
          .testimonial-content {
            margin-bottom: 1rem;
            
            p {
              font-style: italic;
              position: relative;
              margin: 0;
              
              &::before, &::after {
                content: '"';
                color: $primary;
                font-size: 1.5rem;
              }
            }
          }
          
          .testimonial-author {
            text-align: right;
            
            p {
              margin: 0;
            }
          }
        }
      }
    }
    
    .cta-section {
      padding: 4rem 0;
      background-color: $primary;
      color: white;
      
      .cta-content {
        text-align: center;
        max-width: 600px;
        margin: 0 auto;
        
        h2 {
          margin-bottom: 1rem;
        }
        
        p {
          margin-bottom: 2rem;
          font-size: 1.1rem;
        }
        
        .btn {
          background-color: white;
          color: $primary;
          border-color: white;
          
          &:hover {
            background-color: darken(white, 10%);
            border-color: darken(white, 10%);
          }
        }
      }
    }
  }
  </style>