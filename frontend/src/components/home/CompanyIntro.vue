<template>
    <div class="company-intro">
      <div class="intro-image">
        <img src="@/assets/images/company-intro.jpg" alt="Svensk Hälsovård Team" />
      </div>
      <div class="intro-content">
        <h2>Välkommen till Svensk Hälsovård</h2>
        <p>Vi är dedikerade till att hjälpa dig uppnå optimal hälsa genom våra professionella hälsokontroller och blodprovsanalyser.</p>
        <p>Med över 10 års erfarenhet i branschen erbjuder vi både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa.</p>
        <div class="intro-stats">
          <div class="stat-item">
            <div class="stat-value">10+</div>
            <div class="stat-label">Års erfarenhet</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">15 000+</div>
            <div class="stat-label">Nöjda kunder</div>
          </div>
          <div class="stat-item">
            <div class="stat-value">99%</div>
            <div class="stat-label">Kundnöjdhet</div>
          </div>
        </div>
        <div class="intro-cta">
          <router-link to="#services" class="btn btn-primary">Utforska våra tjänster</router-link>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CompanyIntro'
  };
  </script>
  
  <style lang="scss" scoped>
  .company-intro {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 2rem;
    
    .intro-image {
      flex: 0 0 100%;
      max-width: 100%;
      
      @media (min-width: $breakpoint-md) {
        flex: 0 0 40%;
        max-width: 40%;
      }
      
      img {
        width: 100%;
        border-radius: $border-radius;
        box-shadow: $box-shadow;
      }
    }
    
    .intro-content {
      flex: 0 0 100%;
      max-width: 100%;
      
      @media (min-width: $breakpoint-md) {
        flex: 0 0 calc(60% - 2rem);
        max-width: calc(60% - 2rem);
      }
      
      h2 {
        color: $primary;
        margin-bottom: 1.5rem;
      }
      
      p {
        margin-bottom: 1.5rem;
        font-size: 1.1rem;
        line-height: 1.6;
      }
      
      .intro-stats {
        display: flex;
        flex-wrap: wrap;
        gap: 1.5rem;
        margin-bottom: 2rem;
        
        .stat-item {
          flex: 0 0 100%;
          max-width: 100%;
          text-align: center;
          
          @media (min-width: $breakpoint-sm) {
            flex: 0 0 calc(33.333333% - 1rem);
            max-width: calc(33.333333% - 1rem);
          }
          
          .stat-value {
            font-size: 2.5rem;
            font-weight: bold;
            color: $primary;
            margin-bottom: 0.5rem;
          }
          
          .stat-label {
            font-size: 1rem;
            color: #6c757d;
          }
        }
      }
      
      .intro-cta {
        .btn {
          padding: 0.75rem 2rem;
        }
      }
    }
  }
  </style>