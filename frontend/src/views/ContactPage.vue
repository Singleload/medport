<template>
  <div class="contact-page">
    <SEOHead
      title="Kontakta Oss - Svensk Hälsovård"
      description="Kontakta oss för frågor om våra tjänster eller boka tid för blodprov och hälsokontroller."
    />
    
    <section class="page-header bg-primary-lighter">
      <div class="container">
        <div class="page-header-content">
          <h1>Kontakta Oss</h1>
          <p>Vi finns här för att hjälpa dig med alla dina frågor</p>
        </div>
      </div>
    </section>
    
    <section class="section contact-section">
      <div class="container">
        <div class="contact-grid">
          <div class="contact-info">
            <div class="info-card">
              <h2>Kontaktinformation</h2>
              <p>Tveka inte att kontakta oss om du har några frågor eller behöver hjälp. Vi svarar normalt inom 24 timmar.</p>
              
              <div class="info-items">
                <div class="info-item">
                  <div class="icon">
                    <i class="ri-map-pin-line"></i>
                  </div>
                  <div class="content">
                    <h3>Adress</h3>
                    <p>Hälsogatan 123</p>
                    <p>123 45 Stockholm</p>
                  </div>
                </div>
                
                <div class="info-item">
                  <div class="icon">
                    <i class="ri-phone-line"></i>
                  </div>
                  <div class="content">
                    <h3>Telefon</h3>
                    <p><a href="tel:+46801234567">08-123 45 67</a></p>
                    <p class="info-note">Vardagar 08:00-17:00</p>
                  </div>
                </div>
                
                <div class="info-item">
                  <div class="icon">
                    <i class="ri-mail-line"></i>
                  </div>
                  <div class="content">
                    <h3>E-post</h3>
                    <p><a href="mailto:info@svenskhalsovard.se">info@svenskhalsovard.se</a></p>
                    <p class="info-note">Vi svarar normalt inom 24 timmar</p>
                  </div>
                </div>
                
                <div class="info-item">
                  <div class="icon">
                    <i class="ri-time-line"></i>
                  </div>
                  <div class="content">
                    <h3>Öppettider</h3>
                    <p>Måndag - Fredag: 08:00 - 17:00</p>
                    <p>Lördag: Stängt</p>
                    <p>Söndag: Stängt</p>
                  </div>
                </div>
              </div>
              
              <div class="social-links">
                <h3>Följ oss</h3>
                <div class="links">
                  <a href="#" aria-label="Facebook">
                    <i class="ri-facebook-fill"></i>
                  </a>
                  <a href="#" aria-label="Instagram">
                    <i class="ri-instagram-line"></i>
                  </a>
                  <a href="#" aria-label="LinkedIn">
                    <i class="ri-linkedin-fill"></i>
                  </a>
                </div>
              </div>
            </div>
          </div>
          
          <div class="contact-form-container">
            <div class="form-card">
              <h2>Skicka meddelande</h2>
              <p>Fyll i formuläret nedan så återkommer vi till dig så snart som möjligt</p>
              
              <form @submit.prevent="submitContactForm" class="contact-form">
                <div class="form-group">
                  <label for="name">Namn <span class="required">*</span></label>
                  <input 
                    type="text" 
                    id="name" 
                    v-model="contactForm.name" 
                    class="form-control" 
                    required
                  >
                </div>
                
                <div class="form-group">
                  <label for="email">E-post <span class="required">*</span></label>
                  <input 
                    type="email" 
                    id="email" 
                    v-model="contactForm.email" 
                    class="form-control" 
                    required
                  >
                </div>
                
                <div class="form-group">
                  <label for="subject">Ämne <span class="required">*</span></label>
                  <input 
                    type="text" 
                    id="subject" 
                    v-model="contactForm.subject" 
                    class="form-control" 
                    required
                  >
                </div>
                
                <div class="form-group">
                  <label for="message">Meddelande <span class="required">*</span></label>
                  <textarea 
                    id="message" 
                    v-model="contactForm.message" 
                    class="form-control" 
                    rows="5" 
                    required
                  ></textarea>
                </div>
                
                <div class="form-group">
                  <div class="form-check">
                    <input 
                      type="checkbox" 
                      id="consent" 
                      v-model="contactForm.consent" 
                      class="form-check-input" 
                      required
                    >
                    <label for="consent" class="form-check-label">
                      Jag godkänner att Svensk Hälsovård behandlar mina personuppgifter enligt deras integritetspolicy.
                    </label>
                  </div>
                </div>
                
                <div v-if="formSubmitted" class="alert" :class="formSubmitSuccess ? 'alert-success' : 'alert-danger'">
                  <div class="alert-icon">
                    <i v-if="formSubmitSuccess" class="ri-check-line"></i>
                    <i v-else class="ri-error-warning-line"></i>
                  </div>
                  <div class="alert-content">
                    {{ formSubmitMessage }}
                  </div>
                </div>
                
                <div class="form-actions">
                  <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                    <span v-if="isSubmitting">
                      <i class="ri-loader-4-line ri-spin"></i> Skickar...
                    </span>
                    <span v-else>
                      <i class="ri-send-plane-line"></i> Skicka meddelande
                    </span>
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
    
    <section class="section map-section">
      <div class="container">
        <div class="map-container">
          <iframe 
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2034.7806141120752!2d18.054708516169036!3d59.33483218166178!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x465f9d5cbe7d9a1f%3A0x266f5c00bb93edd1!2sStockholm%2C%20Sweden!5e0!3m2!1sen!2sus!4v1631007425457!5m2!1sen!2sus" 
            width="100%" 
            height="450" 
            style="border:0;" 
            allowfullscreen="" 
            loading="lazy"
            title="Karta till Svensk Hälsovård"
          ></iframe>
        </div>
      </div>
    </section>
    
    <section id="faq" class="section faq-section bg-primary-lighter">
      <div class="container">
        <div class="section-title">
          <h2>Vanliga frågor</h2>
          <p>Här hittar du svar på de vanligaste frågorna. Hittar du inte svaret på din fråga? Kontakta oss!</p>
        </div>
        
        <FAQSection />
      </div>
    </section>
  </div>
</template>

<script>
import { ref } from 'vue';
import SEOHead from '@/components/common/SEOHead.vue';
import FAQSection from '@/components/contact/FAQSection.vue';

export default {
  name: 'ContactPage',
  components: {
    SEOHead,
    FAQSection
  },
  setup() {
    const contactForm = ref({
      name: '',
      email: '',
      subject: '',
      message: '',
      consent: false
    });
    
    const isSubmitting = ref(false);
    const formSubmitted = ref(false);
    const formSubmitSuccess = ref(false);
    const formSubmitMessage = ref('');
    
    const submitContactForm = async () => {
      try {
        isSubmitting.value = true;
        formSubmitted.value = false;
        
        // In a real application, you would send this to the backend
        // For this implementation, we'll simulate a successful submission
        // await api.post('/api/contact', contactForm.value);
        
        // Simulate API delay
        await new Promise(resolve => setTimeout(resolve, 1500));
        
        // Reset form on success
        contactForm.value = {
          name: '',
          email: '',
          subject: '',
          message: '',
          consent: false
        };
        
        formSubmitSuccess.value = true;
        formSubmitMessage.value = 'Tack för ditt meddelande! Vi återkommer till dig så snart som möjligt.';
        formSubmitted.value = true;
        
        // Scroll to show the success message
        setTimeout(() => {
          const alertElement = document.querySelector('.alert');
          if (alertElement) {
            alertElement.scrollIntoView({ behavior: 'smooth', block: 'center' });
          }
        }, 100);
      } catch (error) {
        console.error('Contact form submission error:', error);
        formSubmitSuccess.value = false;
        formSubmitMessage.value = 'Ett fel uppstod när meddelandet skulle skickas. Försök igen senare.';
        formSubmitted.value = true;
      } finally {
        isSubmitting.value = false;
      }
    };
    
    return {
      contactForm,
      isSubmitting,
      formSubmitted,
      formSubmitSuccess,
      formSubmitMessage,
      submitContactForm
    };
  }
};
</script>

<style lang="scss" scoped>
.contact-page {
  .page-header {
    padding: 5rem 0;
    background-color: $primary-lighter;
    text-align: center;
    position: relative;
    overflow: hidden;
    
    &::before {
      content: '';
      position: absolute;
      bottom: 0;
      left: 0;
      width: 100%;
      height: 50px;
      background-image: url("data:image/svg+xml,%3Csvg viewBox='0 0 1200 120' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M985.66,92.83C906.67,72,823.78,31,743.84,14.19c-82.26-17.34-168.06-16.33-250.45.39-57.84,11.73-114,31.07-172,41.86A600.21,600.21,0,0,1,0,27.35V120H1200V95.8C1132.19,118.92,1055.71,111.31,985.66,92.83Z' fill='%23ffffff'/%3E%3C/svg%3E");
      background-size: cover;
      background-position: center;
      z-index: 1;
    }
    
    .page-header-content {
      position: relative;
      z-index: 2;
      
      h1 {
        color: $primary;
        font-weight: 800;
        margin-bottom: 1rem;
        font-size: 3rem;
        
        @media (max-width: $breakpoint-md) {
          font-size: 2.5rem;
        }
      }
      
      p {
        font-size: 1.25rem;
        color: #4a5568;
        max-width: 700px;
        margin: 0 auto;
      }
    }
  }
  
  .contact-section {
    padding-top: 3rem;
    
    .contact-grid {
      display: grid;
      grid-template-columns: 1fr;
      gap: 2.5rem;
      
      @media (min-width: $breakpoint-lg) {
        grid-template-columns: 400px 1fr;
      }
      
      .contact-info {
        .info-card {
          background-color: $light;
          border-radius: $border-radius-lg;
          box-shadow: $box-shadow;
          padding: 2rem;
          height: 100%;
          
          h2 {
            color: $primary;
            margin-bottom: 1rem;
            font-weight: 700;
            font-size: 1.75rem;
          }
          
          p {
            color: #4a5568;
            margin-bottom: 2rem;
          }
          
          .info-items {
            margin-bottom: 2rem;
            
            .info-item {
              display: flex;
              align-items: flex-start;
              margin-bottom: 1.5rem;
              
              .icon {
                flex-shrink: 0;
                width: 45px;
                height: 45px;
                border-radius: 50%;
                background-color: rgba($primary, 0.1);
                display: flex;
                align-items: center;
                justify-content: center;
                margin-right: 1rem;
                color: $primary;
                font-size: 1.25rem;
              }
              
              .content {
                flex: 1;
                
                h3 {
                  margin-bottom: 0.5rem;
                  font-size: 1.1rem;
                  font-weight: 700;
                  color: $dark;
                }
                
                p {
                  margin: 0 0 0.25rem;
                  color: #4a5568;
                  
                  &.info-note {
                    font-size: 0.9rem;
                    color: #64748b;
                  }
                }
                
                a {
                  color: $primary;
                  
                  &:hover {
                    text-decoration: underline;
                  }
                }
              }
            }
          }
          
          .social-links {
            h3 {
              margin-bottom: 1rem;
              font-size: 1.1rem;
              font-weight: 700;
              color: $dark;
            }
            
            .links {
              display: flex;
              gap: 1rem;
              
              a {
                width: 40px;
                height: 40px;
                border-radius: 50%;
                background-color: rgba($primary, 0.1);
                color: $primary;
                display: flex;
                align-items: center;
                justify-content: center;
                font-size: 1.25rem;
                transition: $transition-base;
                
                &:hover {
                  background-color: $primary;
                  color: $light;
                  transform: translateY(-3px);
                }
              }
            }
          }
        }
      }
      
      .contact-form-container {
        .form-card {
          background-color: $light;
          border-radius: $border-radius-lg;
          box-shadow: $box-shadow;
          padding: 2rem;
          
          h2 {
            color: $primary;
            margin-bottom: 0.5rem;
            font-weight: 700;
            font-size: 1.75rem;
          }
          
          p {
            color: #4a5568;
            margin-bottom: 2rem;
          }
          
          .contact-form {
            .form-group {
              margin-bottom: 1.5rem;
              
              label {
                display: block;
                margin-bottom: 0.5rem;
                font-weight: 600;
                color: $dark;
                
                .required {
                  color: $danger;
                }
              }
              
              .form-control {
                width: 100%;
                padding: 0.875rem 1rem;
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
                min-height: 120px;
              }
              
              .form-check {
                display: flex;
                align-items: flex-start;
                
                .form-check-input {
                  margin-top: 0.3rem;
                  margin-right: 0.75rem;
                }
                
                .form-check-label {
                  font-weight: normal;
                  font-size: 0.9rem;
                }
              }
            }
            
            .alert {
              display: flex;
              align-items: flex-start;
              padding: 1.25rem;
              border-radius: $border-radius;
              margin-bottom: 1.5rem;
              
              .alert-icon {
                flex-shrink: 0;
                margin-right: 1rem;
                font-size: 1.5rem;
              }
              
              &.alert-success {
                background-color: lighten($success, 45%);
                border-left: 4px solid $success;
                
                .alert-icon {
                  color: $success;
                }
              }
              
              &.alert-danger {
                background-color: lighten($danger, 45%);
                border-left: 4px solid $danger;
                
                .alert-icon {
                  color: $danger;
                }
              }
            }
            
            .form-actions {
              .btn {
                padding: 0.875rem 2rem;
                font-weight: 600;
                display: flex;
                align-items: center;
                justify-content: center;
                gap: 0.5rem;
                
                i {
                  font-size: 1.25rem;
                }
              }
            }
          }
        }
      }
    }
  }
  
  .map-section {
    .map-container {
      border-radius: $border-radius-lg;
      overflow: hidden;
      box-shadow: $box-shadow;
    }
  }
}
</style>