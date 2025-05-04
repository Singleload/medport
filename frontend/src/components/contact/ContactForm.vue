<template>
    <div id="contact-form" class="contact-form-component">
      <form @submit.prevent="submitForm" class="contact-form">
        <div class="form-group">
          <label for="name">Namn *</label>
          <input 
            type="text" 
            id="name" 
            v-model="form.name" 
            class="form-control" 
            required
          >
        </div>
        
        <div class="form-group">
          <label for="email">E-post *</label>
          <input 
            type="email" 
            id="email" 
            v-model="form.email" 
            class="form-control" 
            required
          >
        </div>
        
        <div class="form-group">
          <label for="subject">Ämne *</label>
          <input 
            type="text" 
            id="subject" 
            v-model="form.subject" 
            class="form-control" 
            required
          >
        </div>
        
        <div class="form-group">
          <label for="message">Meddelande *</label>
          <textarea 
            id="message" 
            v-model="form.message" 
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
              v-model="form.consent" 
              class="form-check-input" 
              required
            >
            <label for="consent" class="form-check-label">
              Jag godkänner att Svensk Hälsovård behandlar mina personuppgifter enligt deras integritetspolicy.
            </label>
          </div>
        </div>
        
        <div v-if="formSubmitted" class="alert" :class="formSuccess ? 'alert-success' : 'alert-danger'">
          {{ formMessage }}
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
          await new Promise(resolve => setTimeout(resolve, 1000));
          
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
  </style>