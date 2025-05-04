<template>
  <div id="contact-form" class="contact-form-component">
    <form @submit.prevent="submitForm" class="contact-form">
      <div class="form-group">
        <label for="name">Namn <span class="required">*</span></label>
        <input 
          type="text" 
          id="name" 
          v-model="form.name" 
          class="form-control" 
          placeholder="Ditt namn"
          required
        >
      </div>
      
      <div class="form-group">
        <label for="email">E-post <span class="required">*</span></label>
        <input 
          type="email" 
          id="email" 
          v-model="form.email" 
          class="form-control"
          placeholder="Din e-postadress" 
          required
        >
      </div>
      
      <div class="form-group">
        <label for="subject">Ämne <span class="required">*</span></label>
        <input 
          type="text" 
          id="subject" 
          v-model="form.subject" 
          class="form-control"
          placeholder="Vad gäller ditt meddelande?" 
          required
        >
      </div>
      
      <div class="form-group">
        <label for="message">Meddelande <span class="required">*</span></label>
        <textarea 
          id="message" 
          v-model="form.message" 
          class="form-control" 
          rows="5"
          placeholder="Skriv ditt meddelande här..." 
          required
        ></textarea>
      </div>
      
      <div class="form-group">
        <div class="form-check">
          <input 
            type="checkbox" 
            id="consent" 
            v-model="form.consent" 
            class="form-check-input" 
            required
          >
          <label for="consent" class="form-check-label">
            Jag godkänner att Svensk Hälsovård behandlar mina personuppgifter enligt deras <a href="#" class="privacy-link">integritetspolicy</a>.
          </label>
        </div>
      </div>
      
      <div v-if="formSubmitted" class="alert" :class="formSuccess ? 'alert-success' : 'alert-danger'">
        <div class="alert-icon">
          <i v-if="formSuccess" class="ri-check-line"></i>
          <i v-else class="ri-error-warning-line"></i>
        </div>
        <div class="alert-content">
          {{ formMessage }}
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
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'ContactForm',
  emits: ['form-submitted'],
  setup(props, { emit }) {
    const form = ref({
      name: '',
      email: '',
      subject: '',
      message: '',
      consent: false
    });
    
    const isSubmitting = ref(false);
    const formSubmitted = ref(false);
    const formSuccess = ref(false);
    const formMessage = ref('');
    
    const submitForm = async () => {
      try {
        isSubmitting.value = true;
        formSubmitted.value = false;
        
        // In a real application, you would send this to the backend
        // For this implementation, we'll simulate a successful submission
        await new Promise(resolve => setTimeout(resolve, 1500));
        
        // Emit event with form data
        emit('form-submitted', {
          success: true,
          data: { ...form.value }
        });
        
        // Reset form on success
        form.value = {
          name: '',
          email: '',
          subject: '',
          message: '',
          consent: false
        };
        
        formSuccess.value = true;
        formMessage.value = 'Tack för ditt meddelande! Vi återkommer till dig så snart som möjligt.';
        formSubmitted.value = true;
      } catch (error) {
        console.error('Contact form submission error:', error);
        
        // Emit event with error
        emit('form-submitted', {
          success: false,
          error: error.message
        });
        
        formSuccess.value = false;
        formMessage.value = 'Ett fel uppstod när meddelandet skulle skickas. Försök igen senare.';
        formSubmitted.value = true;
      } finally {
        isSubmitting.value = false;
      }
    };
    
    return {
      form,
      isSubmitting,
      formSubmitted,
      formSuccess,
      formMessage,
      submitForm
    };
  }
};
</script>

<style lang="scss" scoped>
.contact-form-component {
  .contact-form {
    background-color: $light;
    border-radius: $border-radius-lg;
    box-shadow: $box-shadow;
    padding: 2rem;
    
    .form-group {
      margin-bottom: 1.5rem;
      
      label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: $font-weight-medium;
        color: $dark;
        
        .required {
          color: $danger;
          margin-left: 2px;
        }
      }
      
      .form-control {
        width: 100%;
        padding: 0.75rem 1rem;
        border: 1px solid #e2e8f0;
        border-radius: $border-radius;
        transition: $transition-base;
        
        &:focus {
          border-color: $primary;
          box-shadow: 0 0 0 3px rgba($primary, 0.15);
          outline: none;
        }
        
        &::placeholder {
          color: #a0aec0;
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
          cursor: pointer;
        }
        
        .form-check-label {
          font-weight: normal;
          font-size: 0.95rem;
          color: #4a5568;
          
          .privacy-link {
            color: $primary;
            text-decoration: underline;
            transition: $transition-base;
            
            &:hover {
              color: darken($primary, 10%);
            }
          }
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
        
        .alert-content {
          color: darken($success, 20%);
        }
      }
      
      &.alert-danger {
        background-color: lighten($danger, 45%);
        border-left: 4px solid $danger;
        
        .alert-icon {
          color: $danger;
        }
        
        .alert-content {
          color: darken($danger, 20%);
        }
      }
      
      animation: fadeIn 0.3s ease-in-out;
      
      @keyframes fadeIn {
        from {
          opacity: 0;
          transform: translateY(-10px);
        }
        to {
          opacity: 1;
          transform: translateY(0);
        }
      }
    }
    
    .form-actions {
      .btn {
        padding: 0.875rem 2rem;
        font-weight: $font-weight-medium;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.5rem;
        box-shadow: $box-shadow;
        
        i {
          font-size: 1.25rem;
        }
        
        &:hover:not(:disabled) {
          transform: translateY(-3px);
          box-shadow: $box-shadow-hover;
        }
        
        &:disabled {
          opacity: 0.7;
          cursor: not-allowed;
        }
      }
    }
  }
}
</style>