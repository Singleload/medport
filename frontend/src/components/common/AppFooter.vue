<template>
  <footer class="footer">
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
      
      <div class="footer-bottom">
        <p>&copy; {{ currentYear }} Svensk Hälsovård. Alla rättigheter förbehållna.</p>
        <div class="payment-methods">
          <span>Betalningsmetoder:</span>
          <span>Svea Ekonomi</span>
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
    
    return {
      currentYear
    };
  }
};
</script>