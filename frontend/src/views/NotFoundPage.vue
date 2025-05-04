<template>
  <div class="not-found-page">
    <div class="container">
      <div class="not-found-content">
        <div class="error-image">
          <i class="ri-error-warning-line"></i>
        </div>
        <div class="error-code">404</div>
        <h1>Sidan hittades inte</h1>
        <p>Tyvärr kunde vi inte hitta sidan du söker. Den kan ha flyttats, tagits bort eller så har du angett fel adress.</p>
        <div class="actions">
          <router-link to="/" class="btn btn-primary">
            <i class="ri-home-line"></i> Tillbaka till startsidan
          </router-link>
          <router-link to="/contact" class="btn btn-outline-primary">
            <i class="ri-customer-service-line"></i> Kontakta oss
          </router-link>
        </div>
        
        <div class="suggestions">
          <h2>Du kan även prova dessa sidor:</h2>
          <ul>
            <li>
              <router-link to="/">
                <i class="ri-arrow-right-s-line"></i> Hem
              </router-link>
            </li>
            <li>
              <router-link to="/contact">
                <i class="ri-arrow-right-s-line"></i> Kontakt
              </router-link>
            </li>
            <li>
              <router-link to="/service/1">
                <i class="ri-arrow-right-s-line"></i> Hälsokontroll - Kvinna
              </router-link>
            </li>
            <li>
              <router-link to="/service/2">
                <i class="ri-arrow-right-s-line"></i> Hälsokontroll - Man
              </router-link>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFoundPage',
  created() {
    // Set page title
    document.title = 'Sidan hittades inte - Svensk Hälsovård';
  }
};
</script>

<style lang="scss" scoped>
.not-found-page {
  padding: 5rem 0;
  
  .not-found-content {
    text-align: center;
    max-width: 700px;
    margin: 0 auto;
    background-color: $light;
    padding: 3rem 2rem;
    border-radius: $border-radius-lg;
    box-shadow: $box-shadow;
    
    .error-image {
      font-size: 6rem;
      color: $accent;
      margin-bottom: 1rem;
      animation: pulse 2s infinite;
      
      @keyframes pulse {
        0% {
          transform: scale(1);
        }
        50% {
          transform: scale(1.05);
        }
        100% {
          transform: scale(1);
        }
      }
    }
    
    .error-code {
      font-size: 8rem;
      font-weight: 800;
      color: $primary;
      line-height: 1;
      margin-bottom: 1rem;
      background: linear-gradient(135deg, $primary, $primary-light);
      -webkit-background-clip: text;
      background-clip: text;
      -webkit-text-fill-color: transparent;
    }
    
    h1 {
      margin-bottom: 1.5rem;
      font-weight: 700;
      color: $dark;
    }
    
    p {
      font-size: 1.1rem;
      margin-bottom: 2rem;
      color: #4a5568;
    }
    
    .actions {
      display: flex;
      justify-content: center;
      gap: 1rem;
      margin-bottom: 2.5rem;
      
      @media (max-width: $breakpoint-sm) {
        flex-direction: column;
        
        .btn {
          width: 100%;
        }
      }
      
      .btn {
        padding: 0.75rem 2rem;
        font-weight: 600;
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 0.5rem;
        
        i {
          font-size: 1.2rem;
        }
      }
    }
    
    .suggestions {
      margin-top: 2rem;
      text-align: left;
      background-color: $primary-lighter;
      padding: 1.5rem;
      border-radius: $border-radius;
      
      h2 {
        font-size: 1.2rem;
        color: $primary;
        margin-bottom: 1rem;
        text-align: center;
      }
      
      ul {
        list-style: none;
        padding: 0;
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 0.5rem 2rem;
        
        @media (max-width: $breakpoint-sm) {
          grid-template-columns: 1fr;
        }
        
        li {
          a {
            display: flex;
            align-items: center;
            padding: 0.5rem;
            color: $primary;
            font-weight: 500;
            transition: $transition-base;
            border-radius: $border-radius-sm;
            
            &:hover {
              background-color: rgba($primary, 0.1);
              transform: translateX(5px);
            }
            
            i {
              margin-right: 0.5rem;
            }
          }
        }
      }
    }
  }
}
</style>