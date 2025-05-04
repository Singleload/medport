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
          <h1>Din väg till <span class="text-primary">bättre hälsa</span></h1>
          <p>Vi erbjuder professionella hälsokontroller och blodprover för att hjälpa dig ta kontroll över din hälsa.</p>
          <div class="hero-cta">
            <a href="#services" class="btn btn-primary">Se våra tjänster</a>
            <a href="#about" class="btn btn-outline-primary">Om oss</a>
          </div>
          <div class="hero-features">
            <div class="feature">
              <div class="feature-icon">
                <i class="ri-shield-check-line"></i>
              </div>
              <div class="feature-text">
                <span>Ackrediterad</span>
              </div>
            </div>
            <div class="feature">
              <div class="feature-icon">
                <i class="ri-user-heart-line"></i>
              </div>
              <div class="feature-text">
                <span>Expertis</span>
              </div>
            </div>
            <div class="feature">
              <div class="feature-icon">
                <i class="ri-heart-pulse-line"></i>
              </div>
              <div class="feature-text">
                <span>Omfattande</span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="hero-decoration">
        <div class="blob blob-1"></div>
        <div class="blob blob-2"></div>
      </div>
    </section>
    
    <section id="about" class="section about-section">
      <div class="container">
        <div class="section-title">
          <h2>Om Svensk Hälsovård</h2>
          <p>Vi hjälper dig ta kontroll över din hälsa med moderna och tillförlitliga metoder</p>
        </div>
        
        <div class="about-content">
          <div class="about-image img-decoration decoration-1">
            <img src="@/assets/images/company-intro.jpg" alt="Svensk Hälsovård Team" />
          </div>
          
          <div class="about-text">
            <p>Svensk Hälsovård är en ledande leverantör av hälsokontroller och blodprovsanalyser. Vi erbjuder både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa.</p>
            <p>Vårt erfarna team av läkare och sjuksköterskor garanterar högkvalitativ service och professionell rådgivning.</p>
            
            <div class="about-stats">
              <div class="stat-item">
                <div class="stat-value">10<span>+</span></div>
                <div class="stat-label">Års erfarenhet</div>
              </div>
              <div class="stat-item">
                <div class="stat-value">15 000<span>+</span></div>
                <div class="stat-label">Nöjda kunder</div>
              </div>
              <div class="stat-item">
                <div class="stat-value">99<span>%</span></div>
                <div class="stat-label">Kundnöjdhet</div>
              </div>
            </div>
            
            <div class="about-features">
              <div class="feature-item">
                <div class="feature-icon">
                  <i class="ri-medal-line"></i>
                </div>
                <div class="feature-content">
                  <h3>Kvalitetssäkrad</h3>
                  <p>Våra laboratoriepartners är ackrediterade enligt svenska kvalitetsstandarder.</p>
                </div>
              </div>
              
              <div class="feature-item">
                <div class="feature-icon">
                  <i class="ri-lock-line"></i>
                </div>
                <div class="feature-content">
                  <h3>Säker hantering</h3>
                  <p>Din hälsodata hanteras med högsta sekretess och säkerhet.</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    
    <section id="services" class="section services-section bg-primary-lighter">
      <div class="container">
        <div class="section-title">
          <h2>Våra tjänster</h2>
          <p>Vi erbjuder ett brett utbud av hälsokontroller och blodprover för att möta dina behov</p>
        </div>
        
        <ServiceList :services="services" />
      </div>
    </section>
    
    <section class="section how-it-works-section">
      <div class="container">
        <div class="section-title">
          <h2>Hur det fungerar</h2>
          <p>Enkel process från bokning till resultat</p>
        </div>
        
        <div class="steps">
          <div class="step">
            <div class="step-number">1</div>
            <div class="step-icon">
              <i class="ri-shopping-cart-2-line"></i>
            </div>
            <h3>Välj och boka</h3>
            <p>Välj den tjänst du är intresserad av och genomför köpet online.</p>
          </div>
          
          <div class="step">
            <div class="step-number">2</div>
            <div class="step-icon">
              <i class="ri-calendar-check-line"></i>
            </div>
            <h3>Bekräftelse</h3>
            <p>Du får en bekräftelse via e-post med information om din bokning.</p>
          </div>
          
          <div class="step">
            <div class="step-number">3</div>
            <div class="step-icon">
              <i class="ri-test-tube-line"></i>
            </div>
            <h3>Provtagning</h3>
            <p>Besök en av våra kliniker för din bokning eller få hembesök.</p>
          </div>
          
          <div class="step">
            <div class="step-number">4</div>
            <div class="step-icon">
              <i class="ri-file-list-3-line"></i>
            </div>
            <h3>Resultat</h3>
            <p>Få tillgång till dina resultat digitalt och personlig uppföljning.</p>
          </div>
        </div>
      </div>
    </section>
    
    <section class="section testimonials-section bg-primary-lighter">
      <div class="container">
        <div class="section-title">
          <h2>Vad våra kunder säger</h2>
          <p>Lyssna på vad våra nöjda kunder har att säga om våra tjänster</p>
        </div>
        
        <div class="testimonials">
          <div class="testimonial">
            <div class="testimonial-content">
              <p>"Smidigt och professionellt! Jag fick en grundlig genomgång av mina resultat och känner mig trygg med min hälsa."</p>
            </div>
            <div class="testimonial-author">
              <div class="author-avatar">
                <img src="@/assets/images/testimonial-1.jpg" alt="Anna Johansson" />
              </div>
              <div class="author-info">
                <p><strong>Anna Johansson</strong></p>
                <p>Hälsokontroll - Kvinna</p>
              </div>
            </div>
          </div>
          
          <div class="testimonial">
            <div class="testimonial-content">
              <p>"Prenumerationstjänsten är perfekt för mig som vill hålla koll på mina värden regelbundet. Rekommenderas!"</p>
            </div>
            <div class="testimonial-author">
              <div class="author-avatar">
                <img src="@/assets/images/testimonial-2.jpg" alt="Erik Lindberg" />
              </div>
              <div class="author-info">
                <p><strong>Erik Lindberg</strong></p>
                <p>Blodprov - Prenumeration</p>
              </div>
            </div>
          </div>
          
          <div class="testimonial">
            <div class="testimonial-content">
              <p>"Enkelt att boka och betala. Fick mina resultat snabbt och med tydliga förklaringar."</p>
            </div>
            <div class="testimonial-author">
              <div class="author-avatar">
                <img src="@/assets/images/testimonial-3.jpg" alt="Maria Svensson" />
              </div>
              <div class="author-info">
                <p><strong>Maria Svensson</strong></p>
                <p>Blodprov - Premium</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    
    <section class="section faq-preview-section">
      <div class="container">
        <div class="section-title">
          <h2>Vanliga frågor</h2>
          <p>Här hittar du svar på de vanligaste frågorna</p>
        </div>
        
        <div class="faq-preview">
          <div class="faq-item">
            <div class="faq-question">
              <h3>Vad är Svensk Hälsovård?</h3>
              <i class="ri-arrow-right-s-line"></i>
            </div>
            <div class="faq-answer">
              <p>Svensk Hälsovård är en ledande leverantör av hälsokontroller och blodprovsanalyser i Sverige. Vi erbjuder både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa.</p>
            </div>
          </div>
          
          <div class="faq-item">
            <div class="faq-question">
              <h3>Hur fungerar processen?</h3>
              <i class="ri-arrow-right-s-line"></i>
            </div>
            <div class="faq-answer">
              <p>Processen är enkel! Välj den tjänst du är intresserad av och genomför köpet. Du kommer sedan att få en bekräftelse via e-post med information om din bokning. Efter genomförd undersökning kommer du att få tillgång till dina resultat digitalt.</p>
            </div>
          </div>
          
          <div class="faq-item">
            <div class="faq-question">
              <h3>Vad kostar era hälsokontroller?</h3>
              <i class="ri-arrow-right-s-line"></i>
            </div>
            <div class="faq-answer">
              <p>Priserna för våra hälsokontroller varierar beroende på omfattning och innehåll. Vår baskontroll börjar från 995 kr och våra mer omfattande hälsokontroller kostar från 2995 kr.</p>
            </div>
          </div>
        </div>
        
        <div class="faq-cta">
          <router-link to="/contact#faq" class="btn btn-outline-primary">Se alla frågor</router-link>
        </div>
      </div>
    </section>
    
    <section class="section cta-section bg-primary">
      <div class="container">
        <div class="cta-content">
          <h2>Ta kontroll över din hälsa idag</h2>
          <p>Boka en hälsokontroll eller blodprov och investera i din framtida hälsa.</p>
          <a href="#services" class="btn btn-accent">Se våra tjänster</a>
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
    padding: 7rem 0;
    margin-bottom: 0;
    position: relative;
    overflow: hidden;
    
    .hero-content {
      position: relative;
      z-index: 10;
      
      h1 {
        font-size: 3.5rem;
        font-weight: 800;
        margin-bottom: 1.5rem;
        line-height: 1.2;
        
        span {
          position: relative;
          display: inline-block;
          
          &::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            height: 8px;
            background-color: rgba($accent, 0.5);
            z-index: -1;
          }
        }
        
        @media (max-width: $breakpoint-md) {
          font-size: 2.5rem;
        }
      }
      
      p {
        font-size: 1.25rem;
        margin-bottom: 2rem;
        color: #4a5568;
        max-width: 600px;
      }
      
      .hero-cta {
        display: flex;
        gap: 1rem;
        margin-bottom: 3rem;
        
        @media (max-width: $breakpoint-sm) {
          flex-direction: column;
          
          .btn {
            width: 100%;
          }
        }
        
        .btn {
          padding: 1rem 2rem;
          font-weight: 600;
        }
      }
      
      .hero-features {
        display: flex;
        gap: 2rem;
        
        @media (max-width: $breakpoint-sm) {
          flex-direction: column;
          gap: 1rem;
        }
        
        .feature {
          display: flex;
          align-items: center;
          
          .feature-icon {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            background-color: rgba($primary, 0.1);
            display: flex;
            align-items: center;
            justify-content: center;
            margin-right: 0.75rem;
            color: $primary;
            font-size: 1.25rem;
          }
          
          .feature-text {
            font-weight: 500;
            color: $dark;
          }
        }
      }
    }
    
    .hero-decoration {
      position: absolute;
      top: 0;
      right: 0;
      width: 100%;
      height: 100%;
      z-index: 1;
      
      .blob {
        position: absolute;
        border-radius: 50%;
        
        &.blob-1 {
          top: -150px;
          right: -100px;
          width: 500px;
          height: 500px;
          background-color: rgba($primary-lighter, 0.7);
          animation: float 8s ease-in-out infinite;
        }
        
        &.blob-2 {
          bottom: -200px;
          right: 250px;
          width: 400px;
          height: 400px;
          background-color: rgba($primary-lighter, 0.5);
          animation: float 10s ease-in-out infinite 1s;
        }
      }
      
      @keyframes float {
        0% {
          transform: translateY(0) scale(1);
        }
        50% {
          transform: translateY(-20px) scale(1.05);
        }
        100% {
          transform: translateY(0) scale(1);
        }
      }
    }
  }
  
  .about-section {
    .about-content {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      gap: 3rem;
      
      .about-image {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-md) {
          flex: 0 0 45%;
          max-width: 45%;
        }
        
        img {
          width: 100%;
          border-radius: $border-radius-lg;
          box-shadow: $box-shadow;
        }
      }
      
      .about-text {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-md) {
          flex: 0 0 calc(55% - 3rem);
          max-width: calc(55% - 3rem);
        }
        
        p {
          font-size: 1.1rem;
          margin-bottom: 1.5rem;
          color: #4a5568;
          line-height: 1.7;
        }
        
        .about-stats {
          display: flex;
          flex-wrap: wrap;
          gap: 2rem;
          margin: 2.5rem 0;
          
          .stat-item {
            flex: 1;
            min-width: 100px;
            text-align: center;
            padding: 1.5rem;
            background-color: $light;
            border-radius: $border-radius-lg;
            box-shadow: $box-shadow;
            transition: $transition-base;
            
            &:hover {
              transform: translateY(-5px);
              box-shadow: $box-shadow-hover;
            }
            
            .stat-value {
              font-size: 2.5rem;
              font-weight: 800;
              color: $primary;
              margin-bottom: 0.5rem;
              line-height: 1;
              
              span {
                font-size: 1.5rem;
                font-weight: 700;
              }
            }
            
            .stat-label {
              font-size: 1rem;
              color: #64748b;
              font-weight: 500;
            }
          }
        }
        
        .about-features {
          display: flex;
          flex-wrap: wrap;
          gap: 1.5rem;
          
          .feature-item {
            flex: 0 0 100%;
            display: flex;
            align-items: flex-start;
            padding: 1.5rem;
            background-color: $light;
            border-radius: $border-radius-lg;
            box-shadow: $box-shadow;
            transition: $transition-base;
            
            @media (min-width: $breakpoint-sm) {
              flex: 0 0 calc(50% - 0.75rem);
            }
            
            &:hover {
              transform: translateY(-5px);
              box-shadow: $box-shadow-hover;
            }
            
            .feature-icon {
              flex-shrink: 0;
              width: 50px;
              height: 50px;
              border-radius: 50%;
              background-color: rgba($primary, 0.1);
              display: flex;
              align-items: center;
              justify-content: center;
              margin-right: 1rem;
              color: $primary;
              font-size: 1.5rem;
            }
            
            .feature-content {
              flex: 1;
              
              h3 {
                margin-bottom: 0.5rem;
                font-size: 1.2rem;
                font-weight: 700;
              }
              
              p {
                margin-bottom: 0;
                font-size: 1rem;
              }
            }
          }
        }
      }
    }
  }
  
  .how-it-works-section {
    .steps {
      display: grid;
      grid-template-columns: repeat(1, 1fr);
      gap: 2rem;
      
      @media (min-width: $breakpoint-sm) {
        grid-template-columns: repeat(2, 1fr);
      }
      
      @media (min-width: $breakpoint-lg) {
        grid-template-columns: repeat(4, 1fr);
      }
      
      .step {
        position: relative;
        text-align: center;
        padding: 3rem 1.5rem 2rem;
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-8px);
          box-shadow: $box-shadow-hover;
          
          .step-icon {
            background-color: $primary;
            color: $light;
          }
        }
        
        .step-number {
          position: absolute;
          top: -20px;
          left: 50%;
          transform: translateX(-50%);
          width: 40px;
          height: 40px;
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
          width: 80px;
          height: 80px;
          margin: 0 auto 1.5rem;
          background-color: rgba($primary, 0.1);
          color: $primary;
          font-size: 2rem;
          border-radius: 50%;
          display: flex;
          align-items: center;
          justify-content: center;
          transition: $transition-base;
        }
        
        h3 {
          margin-bottom: 1rem;
          font-size: 1.25rem;
          font-weight: 700;
        }
        
        p {
          color: #64748b;
          margin-bottom: 0;
        }
      }
    }
  }
  
  .faq-preview-section {
    .faq-preview {
      max-width: 800px;
      margin: 0 auto;
      
      .faq-item {
        background-color: $light;
        border-radius: $border-radius-lg;
        box-shadow: $box-shadow;
        margin-bottom: 1.5rem;
        overflow: hidden;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-5px);
          box-shadow: $box-shadow-hover;
        }
        
        .faq-question {
          padding: 1.5rem;
          display: flex;
          justify-content: space-between;
          align-items: center;
          cursor: pointer;
          
          h3 {
            margin: 0;
            font-size: 1.2rem;
            font-weight: 600;
          }
          
          i {
            color: $primary;
            font-size: 1.5rem;
            transition: $transition-base;
          }
          
          &:hover i {
            transform: rotate(90deg);
          }
        }
        
        .faq-answer {
          padding: 0 1.5rem 1.5rem;
          
          p {
            margin: 0;
            color: #64748b;
          }
        }
      }
    }
    
    .faq-cta {
      text-align: center;
      margin-top: 2.5rem;
      
      .btn {
        padding: 0.875rem 2rem;
        font-weight: 600;
      }
    }
  }
  
  .cta-section {
    position: relative;
    overflow: hidden;
    
    &::before {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-image: linear-gradient(135deg, rgba($primary, 0.9) 0%, rgba($primary-light, 0.9) 100%);
      z-index: 1;
    }
    
    .cta-content {
      position: relative;
      z-index: 2;
      text-align: center;
      max-width: 800px;
      margin: 0 auto;
      
      h2 {
        color: $light;
        margin-bottom: 1rem;
        font-weight: 700;
        font-size: 2.5rem;
      }
      
      p {
        color: rgba($light, 0.9);
        margin-bottom: 2rem;
        font-size: 1.25rem;
      }
      
      .btn {
        padding: 1rem 2.5rem;
        font-weight: 600;
        font-size: 1.1rem;
      }
    }
  }
}
</style>