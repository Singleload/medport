<template>
  <div class="company-intro">
    <div class="intro-image img-decoration decoration-1">
      <img src="@/assets/images/company-intro.jpg" alt="Svensk Hälsovård Team" />
    </div>
    <div class="intro-content">
      <h2>Välkommen till Svensk Hälsovård</h2>
      <p>Vi är dedikerade till att hjälpa dig uppnå optimal hälsa genom våra professionella hälsokontroller och blodprovsanalyser.</p>
      <p>Med över 10 års erfarenhet i branschen erbjuder vi både engångsundersökningar och prenumerationstjänster för kontinuerlig uppföljning av din hälsa.</p>
      
      <div class="intro-features">
        <div class="feature">
          <div class="feature-icon">
            <i class="ri-shield-check-line"></i>
          </div>
          <div class="feature-content">
            <h3>Kvalitetssäkrad</h3>
            <p>Våra laboratoriepartners är ackrediterade enligt svenska kvalitetsstandarder.</p>
          </div>
        </div>
        
        <div class="feature">
          <div class="feature-icon">
            <i class="ri-user-heart-line"></i>
          </div>
          <div class="feature-content">
            <h3>Expertis</h3>
            <p>Vårt team består av erfarna läkare och sjuksköterskor med specialistkompetens.</p>
          </div>
        </div>
      </div>
      
      <div class="intro-stats">
        <div class="stat-item">
          <div class="stat-value">10<span>+</span></div>
          <div class="stat-label">Års erfarenhet</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">15 000<span>+</span></div>
          <div class="stat-label">Nöjda kunder</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">99<span>%</span></div>
          <div class="stat-label">Kundnöjdhet</div>
        </div>
      </div>
      
      <div class="intro-cta">
        <a href="#services" class="btn btn-primary">
          <i class="ri-heart-pulse-line"></i> Utforska våra tjänster
        </a>
        <a href="#contact" class="btn btn-outline-primary">
          <i class="ri-customer-service-line"></i> Kontakta oss
        </a>
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
  gap: 3rem;
  margin: 3rem 0;
  
  .intro-image {
    flex: 0 0 100%;
    max-width: 100%;
    position: relative;
    
    @media (min-width: $breakpoint-md) {
      flex: 0 0 calc(40% - 1.5rem);
      max-width: calc(40% - 1.5rem);
    }
    
    img {
      width: 100%;
      border-radius: $border-radius-lg;
      box-shadow: $box-shadow;
      transition: $transition-base;
    }
    
    &.decoration-1::before {
      content: '';
      position: absolute;
      width: 100px;
      height: 100px;
      background-color: rgba($primary, 0.1);
      border-radius: 50%;
      z-index: -1;
      top: -20px;
      left: -20px;
    }
    
    &::after {
      content: '';
      position: absolute;
      width: 70px;
      height: 70px;
      background-color: rgba($accent, 0.15);
      border-radius: 50%;
      z-index: -1;
      bottom: -15px;
      right: -15px;
    }
  }
  
  .intro-content {
    flex: 0 0 100%;
    max-width: 100%;
    
    @media (min-width: $breakpoint-md) {
      flex: 0 0 calc(60% - 1.5rem);
      max-width: calc(60% - 1.5rem);
    }
    
    h2 {
      color: $primary;
      margin-bottom: 1.5rem;
      font-weight: 700;
      font-size: 2.25rem;
      
      &::after {
        content: '';
        display: block;
        width: 80px;
        height: 3px;
        background-color: $primary-light;
        margin-top: 1rem;
      }
      
      @media (max-width: $breakpoint-md) {
        font-size: 1.75rem;
      }
    }
    
    p {
      margin-bottom: 1.5rem;
      font-size: 1.1rem;
      line-height: 1.7;
      color: #4a5568;
    }
    
    .intro-features {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      margin-bottom: 2rem;
      
      .feature {
        display: flex;
        align-items: flex-start;
        gap: 1.25rem;
        background-color: $light;
        border-radius: $border-radius-lg;
        padding: 1.5rem;
        box-shadow: $box-shadow-sm;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-3px);
          box-shadow: $box-shadow;
        }
        
        .feature-icon {
          width: 50px;
          height: 50px;
          background-color: $primary-lighter;
          color: $primary;
          border-radius: 50%;
          display: flex;
          align-items: center;
          justify-content: center;
          font-size: 1.5rem;
          flex-shrink: 0;
        }
        
        .feature-content {
          flex: 1;
          
          h3 {
            margin-bottom: 0.5rem;
            font-size: 1.2rem;
            font-weight: 700;
            color: $primary;
          }
          
          p {
            margin-bottom: 0;
            font-size: 0.95rem;
          }
        }
      }
    }
    
    .intro-stats {
      display: flex;
      justify-content: space-between;
      flex-wrap: wrap;
      gap: 1rem;
      margin-bottom: 2rem;
      
      .stat-item {
        flex: 1;
        min-width: 120px;
        background-color: $light;
        border-radius: $border-radius;
        padding: 1.5rem;
        text-align: center;
        box-shadow: $box-shadow-sm;
        transition: $transition-base;
        
        &:hover {
          transform: translateY(-3px);
          box-shadow: $box-shadow;
        }
        
        .stat-value {
          font-size: 2.25rem;
          font-weight: 800;
          color: $primary;
          margin-bottom: 0.5rem;
          
          span {
            font-size: 1.5rem;
            color: $accent;
          }
        }
        
        .stat-label {
          font-size: 0.95rem;
          color: #64748b;
          font-weight: 500;
        }
      }
    }
    
    .intro-cta {
      display: flex;
      flex-wrap: wrap;
      gap: 1rem;
      
      .btn {
        padding: 0.875rem 1.5rem;
        font-weight: 600;
        display: flex;
        align-items: center;
        gap: 0.5rem;
        box-shadow: $box-shadow-sm;
        
        i {
          font-size: 1.2rem;
        }
        
        &:hover {
          transform: translateY(-3px);
          box-shadow: $box-shadow-hover;
        }
      }
    }
  }
}
</style>