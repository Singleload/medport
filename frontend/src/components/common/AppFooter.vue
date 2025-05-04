<template>
    <footer class="footer">
      <div class="container">
        <div class="footer-content">
          <div class="footer-logo">
            <img src="@/assets/images/logo.png" alt="Svensk Hälsovård" />
            <p>Din väg till bättre hälsa</p>
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
                <p>Hälsogatan 123</p>
                <p>123 45 Stockholm</p>
                <p>Tel: <a href="tel:+46801234567">08-123 45 67</a></p>
                <p>E-post: <a href="mailto:info@svenskhalsovard.se">info@svenskhalsovard.se</a></p>
              </address>
            </div>
          </div>
        </div>
        
        <div class="footer-bottom">
          <p>&copy; {{ currentYear }} Svensk Hälsovård. Alla rättigheter förbehållna.</p>
          <div class="payment-methods">
            <span>Betalningsmetoder:</span>
            <span>Svea Ekonomi</span>
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
  
  <style lang="scss" scoped>
  .footer {
    background-color: #f8f9fa;
    padding: 3rem 0 1rem;
    margin-top: 2rem;
    
    .footer-content {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      margin-bottom: 2rem;
      
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
          height: 40px;
          margin-bottom: 0.5rem;
        }
        
        p {
          color: $primary;
          font-weight: 500;
        }
      }
      
      .footer-links {
        flex: 0 0 100%;
        max-width: 100%;
        display: flex;
        flex-wrap: wrap;
        
        @media (min-width: $breakpoint-md) {
          flex: 0 0 70%;
          max-width: 70%;
        }
        
        .link-column {
          flex: 0 0 100%;
          max-width: 100%;
          margin-bottom: 1.5rem;
          
          @media (min-width: $breakpoint-sm) {
            flex: 0 0 33.333333%;
            max-width: 33.333333%;
            margin-bottom: 0;
          }
          
          h4 {
            color: $primary;
            margin-bottom: 1rem;
            font-size: 1.1rem;
          }
          
          ul {
            list-style: none;
            padding: 0;
            margin: 0;
            
            li {
              margin-bottom: 0.5rem;
              
              a {
                color: $dark;
                text-decoration: none;
                
                &:hover {
                  color: $primary;
                }
              }
            }
          }
          
          address {
            font-style: normal;
            
            p {
              margin-bottom: 0.5rem;
            }
          }
        }
      }
    }
    
    .footer-bottom {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-between;
      align-items: center;
      padding-top: 1rem;
      border-top: 1px solid #dee2e6;
      
      p {
        margin: 0;
      }
      
      .payment-methods {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        
        span {
          &:first-child {
            color: #6c757d;
          }
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