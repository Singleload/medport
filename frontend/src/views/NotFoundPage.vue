<template>
    <div class="not-found-page">
      <div class="container">
        <div class="not-found-content">
          <div class="error-code">404</div>
          <h1>Sidan hittades inte</h1>
          <p>Tyvärr kunde vi inte hitta sidan du söker. Den kan ha flyttats, tagits bort eller så har du angett fel adress.</p>
          <div class="actions">
            <router-link to="/" class="btn btn-primary">Tillbaka till startsidan</router-link>
            <router-link to="/contact" class="btn btn-secondary">Kontakta oss</router-link>
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
      max-width: 600px;
      margin: 0 auto;
      
      .error-code {
        font-size: 8rem;
        font-weight: bold;
        color: $primary;
        line-height: 1;
        margin-bottom: 1rem;
      }
      
      h1 {
        margin-bottom: 1.5rem;
      }
      
      p {
        font-size: 1.1rem;
        margin-bottom: 2rem;
        color: #6c757d;
      }
      
      .actions {
        display: flex;
        justify-content: center;
        gap: 1rem;
        
        @media (max-width: $breakpoint-sm) {
          flex-direction: column;
          
          .btn {
            width: 100%;
          }
        }
      }
    }
  }
  </style>