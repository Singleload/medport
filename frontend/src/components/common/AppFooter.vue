<template>
  <footer class="footer">
    <div class="footer-wave">
      <svg viewBox="0 0 1200 120" preserveAspectRatio="none">
        <path d="M321.39,56.44c58-10.79,114.16-30.13,172-41.86,82.39-16.72,168.19-17.73,250.45-.39C823.78,31,906.67,72,985.66,92.83c70.05,18.48,146.53,26.09,214.34,3V0H0V27.35A600.21,600.21,0,0,0,321.39,56.44Z" fill="#f9fbfa"></path>
      </svg>
    </div>
    
    <div class="container">
      <div class="footer-content">
        <div class="footer-logo">
          <img src="@/assets/images/logo.png" alt="Svensk Hälsovård" />
          <p>Din väg till bättre hälsa</p>
          <div class="social-links">
            <a href="#" aria-label="Facebook">
              <i class="ri-facebook-fill"></i>
            </a>
            <a href="#" aria-label="Instagram">
              <i class="ri-instagram-line"></i>
            </a>
            <a href="#" aria-label="LinkedIn">
              <i class="ri-linkedin-fill"></i>
            </a>
            <a href="#" aria-label="Twitter">
              <i class="ri-twitter-fill"></i>
            </a>
          </div>
        </div>
        
        <div class="footer-links">
          <div class="link-column">
            <h4>Tjänster</h4>
            <ul>
              <li><router-link to="/service/1">Hälsokontroll - Kvinna</router-link></li>
              <li><router-link to="/service/2">Hälsokontroll - Man</router-link></li>
              <li><router-link to="/service/3">Blodprov - Bas</router-link></li>
              <li><router-link to="/service/4">Blodprov - Premium</router-link></li>
              <li><router-link to="/service/5">Blodprov - Prenumeration</router-link></li>
            </ul>
          </div>
          
          <div class="link-column">
            <h4>Information</h4>
            <ul>
              <li><router-link to="/contact">Kontakta oss</router-link></li>
              <li><router-link to="/contact#faq">Vanliga frågor</router-link></li>
              <li><a href="#">Integritetspolicy</a></li>
              <li><a href="#">Villkor</a></li>
              <li><a href="#">Leveransvillkor</a></li>
            </ul>
          </div>
          
          <div class="link-column">
            <h4>Kontakt</h4>
            <address>
              <p><i class="ri-map-pin-line"></i> Hälsogatan 123, 123 45 Stockholm</p>
              <p><i class="ri-phone-line"></i> <a href="tel:+46801234567">08-123 45 67</a></p>
              <p><i class="ri-mail-line"></i> <a href="mailto:info@svenskhalsovard.se">info@svenskhalsovard.se</a></p>
              <p><i class="ri-time-line"></i> Mån-Fre: 08:00-17:00</p>
            </address>
          </div>
        </div>
      </div>
      
      <div class="footer-newsletter">
        <div class="newsletter-content">
          <h3>Prenumerera på vårt nyhetsbrev</h3>
          <p>Få de senaste uppdateringarna om våra tjänster och hälsotips direkt i din inkorg.</p>
          <form class="newsletter-form" @submit.prevent="subscribeNewsletter">
            <div class="form-group">
              <input 
                type="email" 
                placeholder="Din e-postadress" 
                class="form-control"
                v-model="newsletterEmail"
                required
              />
              <button type="submit" class="btn btn-accent">
                <i class="ri-send-plane-line"></i>
              </button>
            </div>
          </form>
          <p v-if="newsletterStatus" class="newsletter-status" :class="{'success': newsletterSuccess}">
            {{ newsletterMessage }}
          </p>
        </div>
      </div>
      
      <div class="footer-bottom">
        <p>&copy; {{ currentYear }} Svensk Hälsovård. Alla rättigheter förbehållna.</p>
        <div class="payment-methods">
          <span>Betalningsmetoder:</span>
          <i class="ri-bank-card-line"></i>
          <i class="ri-mastercard-line"></i>
          <i class="ri-visa-line"></i>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'AppFooter',
  setup() {
    const currentYear = ref(new Date().getFullYear());
    const newsletterEmail = ref('');
    const newsletterStatus = ref(false);
    const newsletterSuccess = ref(false);
    const newsletterMessage = ref('');
    
    const subscribeNewsletter = () => {
      // In a real application, you would call an API here
      // For this implementation, we'll simulate a successful subscription
      
      // Show success message
      newsletterSuccess.value = true;
      newsletterMessage.value = 'Tack för din anmälan! Vi har skickat en bekräftelse till din e-post.';
      newsletterStatus.value = true;
      
      // Reset form
      newsletterEmail.value = '';
      
      // Hide message after 5 seconds
      setTimeout(() => {
        newsletterStatus.value = false;
      }, 5000);
    };
    
    return {
      currentYear,
      newsletterEmail,
      newsletterStatus,
      newsletterSuccess,
      newsletterMessage,
      subscribeNewsletter
    };
  }
};
</script>

<style lang="scss" scoped>
.footer {
  background-color: #f8f9fa;
  padding: 5rem 0 2rem;
  margin-top: 5rem;
  position: relative;
  
  .footer-wave {
    position: absolute;
    top: -120px;
    left: 0;
    width: 100%;
    height: 120px;
    overflow: hidden;
    line-height: 0;
    
    svg {
      position: relative;
      display: block;
      width: 100%;
      height: 120px;
    }
  }
  
  .footer-content {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    margin-bottom: 3rem;
    
    .footer-logo {
      flex: 0 0 100%;
      max-width: 100%;
      margin-bottom: 2rem;
      
      @media (min-width: $breakpoint-md) {
        flex: 0 0 25%;
        max-width: 25%;
        margin-bottom: 0;
      }
      
      img {
        height: 50px;
        margin-bottom: 1rem;
      }
      
      p {
        color: $primary;
        font-weight: 500;
        margin-bottom: 1.5rem;
      }
      
      .social-links {
        display: flex;
        gap: 1rem;
        
        a {
          display: flex;
          align-items: center;
          justify-content: center;
          width: 40px;
          height: 40px;
          border-radius: 50%;
          background-color: rgba($primary, 0.1);
          color: $primary;
          transition: $transition-base;
          
          &:hover {
            background-color: $primary;
            color: $light;
            transform: translateY(-3px);
          }
        }
      }
    }
    
    .footer-links {
      flex: 0 0 100%;
      max-width: 100%;
      display: flex;
      flex-wrap: wrap;
      gap: 2rem;
      
      @media (min-width: $breakpoint-md) {
        flex: 0 0 70%;
        max-width: 70%;
      }
      
      .link-column {
        flex: 1;
        min-width: 200px;
        
        h4 {
          color: $primary;
          margin-bottom: 1.25rem;
          font-size: 1.1rem;
          font-weight: 700;
          position: relative;
          padding-bottom: 0.75rem;
          
          &::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 0;
            width: 30px;
            height: 2px;
            background-color: $primary-light;
          }
        }
        
        ul {
          list-style: none;
          padding: 0;
          margin: 0;
          
          li {
            margin-bottom: 0.75rem;
            
            a {
              color: #4a5568;
              text-decoration: none;
              transition: $transition-base;
              display: inline-block;
              
              &:hover {
                color: $primary;
                transform: translateX(3px);
              }
            }
          }
        }
        
        address {
          font-style: normal;
          color: #4a5568;
          
          p {
            margin-bottom: 0.75rem;
            display: flex;
            align-items: center;
            gap: 0.75rem;
            
            i {
              color: $primary;
              font-size: 1.1rem;
              flex-shrink: 0;
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
    }
  }
  
  .footer-newsletter {
    background-color: $primary-lighter;
    padding: 2.5rem;
    border-radius: $border-radius-lg;
    margin-bottom: 3rem;
    
    .newsletter-content {
      max-width: 600px;
      margin: 0 auto;
      text-align: center;
      
      h3 {
        color: $primary;
        margin-bottom: 0.75rem;
        font-weight: 700;
      }
      
      p {
        color: #4a5568;
        margin-bottom: 1.5rem;
      }
      
      .newsletter-form {
        .form-group {
          display: flex;
          max-width: 500px;
          margin: 0 auto;
          
          .form-control {
            flex: 1;
            border-top-right-radius: 0;
            border-bottom-right-radius: 0;
            border-right: none;
          }
          
          .btn {
            border-top-left-radius: 0;
            border-bottom-left-radius: 0;
            padding: 0.75rem 1.25rem;
            
            i {
              margin-right: 0;
            }
          }
        }
      }
      
      .newsletter-status {
        margin-top: 1rem;
        font-size: 0.9rem;
        
        &.success {
          color: $success;
        }
        
        &:not(.success) {
          color: $danger;
        }
      }
    }
  }
  
  .footer-bottom {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    padding-top: 1.5rem;
    border-top: 1px solid #e2e8f0;
    
    p {
      margin: 0;
      color: #64748b;
      font-size: 0.9rem;
    }
    
    .payment-methods {
      display: flex;
      align-items: center;
      gap: 1rem;
      
      span {
        color: #64748b;
        font-size: 0.9rem;
      }
      
      i {
        font-size: 1.25rem;
        color: #64748b;
      }
    }
    
    @media (max-width: $breakpoint-sm) {
      flex-direction: column;
      gap: 1rem;
      text-align: center;
    }
  }
}
</style>