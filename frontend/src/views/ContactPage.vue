<template>
    <div class="contact-page">
      <SEOHead
        title="Kontakta Oss - Svensk Hälsovård"
        description="Kontakta oss för frågor om våra tjänster eller boka tid för blodprov och hälsokontroller."
      />
      
      <div class="container">
        <h1>Kontakta Oss</h1>
        
        <div class="contact-section">
          <div class="contact-info">
            <h2>Kontaktinformation</h2>
            
            <div class="info-item">
              <div class="icon">
                <i class="fas fa-map-marker-alt"></i>
              </div>
              <div class="content">
                <h3>Adress</h3>
                <p>Hälsogatan 123</p>
                <p>123 45 Stockholm</p>
              </div>
            </div>
            
            <div class="info-item">
              <div class="icon">
                <i class="fas fa-phone-alt"></i>
              </div>
              <div class="content">
                <h3>Telefon</h3>
                <p><a href="tel:+46801234567">08-123 45 67</a></p>
                <p class="info-note">Vardagar 08:00-17:00</p>
              </div>
            </div>
            
            <div class="info-item">
              <div class="icon">
                <i class="fas fa-envelope"></i>
              </div>
              <div class="content">
                <h3>E-post</h3>
                <p><a href="mailto:info@svenskhalsovard.se">info@svenskhalsovard.se</a></p>
                <p class="info-note">Vi svarar normalt inom 24 timmar</p>
              </div>
            </div>
            
            <div class="info-item">
              <div class="icon">
                <i class="fas fa-clock"></i>
              </div>
              <div class="content">
                <h3>Öppettider</h3>
                <p>Måndag - Fredag: 08:00 - 17:00</p>
                <p>Lördag: Stängt</p>
                <p>Söndag: Stängt</p>
              </div>
            </div>
          </div>
          
          <div class="contact-form-container">
            <h2>Skicka meddelande</h2>
            
            <form @submit.prevent="submitContactForm" class="contact-form">
              <div class="form-group">
                <label for="name">Namn *</label>
                <input 
                  type="text" 
                  id="name" 
                  v-model="contactForm.name" 
                  class="form-control" 
                  required
                >
              </div>
              
              <div class="form-group">
                <label for="email">E-post *</label>
                <input 
                  type="email" 
                  id="email" 
                  v-model="contactForm.email" 
                  class="form-control" 
                  required
                >
              </div>
              
              <div class="form-group">
                <label for="subject">Ämne *</label>
                <input 
                  type="text" 
                  id="subject" 
                  v-model="contactForm.subject" 
                  class="form-control" 
                  required
                >
              </div>
              
              <div class="form-group">
                <label for="message">Meddelande *</label>
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
                {{ formSubmitMessage }}
              </div>
              
              <div class="form-actions">
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting">
                    <i class="fas fa-spinner fa-spin"></i> Skickar...
                  </span>
                  <span v-else>Skicka meddelande</span>
                </button>
              </div>
            </form>
          </div>
        </div>
        
        <div id="faq" class="faq-section">
          <h2>Vanliga frågor</h2>
          
          <FAQSection />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import SEOHead from '@/components/common/SEOHead.vue';
  import FAQSection from '@/components/contact/FAQSection.vue';
  import api from '@/services/api';
  
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
          await new Promise(resolve => setTimeout(resolve, 1000));
          
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
    padding: 2rem 0;
    
    h1 {
      margin-bottom: 2rem;
      color: $primary;
    }
    
    .contact-section {
      display: flex;
      flex-wrap: wrap;
      gap: 2rem;
      margin-bottom: 4rem;
      
      .contact-info {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-lg) {
          flex: 0 0 calc(40% - 1rem);
          max-width: calc(40% - 1rem);
        }
        
        h2 {
          margin-bottom: 2rem;
          color: $primary;
          font-size: 1.5rem;
        }
        
        .info-item {
          display: flex;
          margin-bottom: 2rem;
          
          .icon {
            flex: 0 0 40px;
            width: 40px;
            height: 40px;
            background-color: rgba($primary, 0.1);
            border-radius: 50%;
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
            }
            
            p {
              margin: 0 0 0.25rem;
              color: #495057;
              
              &.info-note {
                font-size: 0.9rem;
                color: #6c757d;
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
      
      .contact-form-container {
        flex: 0 0 100%;
        max-width: 100%;
        
        @media (min-width: $breakpoint-lg) {
          flex: 0 0 calc(60% - 1rem);
          max-width: calc(60% - 1rem);
        }
        
        h2 {
          margin-bottom: 2rem;
          color: $primary;
          font-size: 1.5rem;
        }
        
        .contact-form {
          background-color: white;
          border-radius: $border-radius;
          box-shadow: $box-shadow-sm;
          padding: 2rem;
          
          .form-group {
            margin-bottom: 1.5rem;
            
            label {
              display: block;
              margin-bottom: 0.5rem;
              font-weight: 500;
            }
            
            .form-control {
              width: 100%;
              padding: 0.75rem;
              border: 1px solid #dee2e6;
              border-radius: $border-radius-sm;
              transition: $transition-base;
              
              &:focus {
                border-color: $primary;
                box-shadow: 0 0 0 0.2rem rgba($primary, 0.25);
                outline: none;
              }
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
            padding: 1rem;
            border-radius: $border-radius-sm;
            margin-bottom: 1.5rem;
            
            &.alert-success {
              background-color: lighten($success, 45%);
              border: 1px solid lighten($success, 35%);
              color: darken($success, 10%);
            }
            
            &.alert-danger {
              background-color: lighten($danger, 45%);
              border: 1px solid lighten($danger, 35%);
              color: darken($danger, 10%);
            }
          }
          
          .form-actions {
            .btn {
              padding: 0.75rem 2rem;
            }
          }
        }
      }
    }
    
    .faq-section {
      margin-bottom: 4rem;
      
      h2 {
        margin-bottom: 2rem;
        color: $primary;
        font-size: 1.5rem;
      }
    }
  }
  </style>