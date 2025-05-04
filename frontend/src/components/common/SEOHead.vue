<template>
    <div style="display: none;">
    <!-- This component doesn't render anything visible in the DOM -->
  </div>
</template>
  
  <script>
  export default {
    name: 'SEOHead',
    props: {
      title: {
        type: String,
        required: true
      },
      description: {
        type: String,
        required: true
      },
      keywords: {
        type: String,
        default: 'hälsovård, blodprov, hälsokontroll, svensk hälsovård'
      },
      author: {
        type: String,
        default: 'Svensk Hälsovård'
      },
      type: {
        type: String,
        default: 'website'
      },
      url: {
        type: String,
        default: ''
      },
      image: {
        type: String,
        default: '/assets/images/og-image.jpg'
      },
      twitterCard: {
        type: String,
        default: 'summary_large_image'
      },
      schemaType: {
        type: String,
        default: 'WebPage'
      },
      schemaData: {
        type: Object,
        default: () => ({})
      }
    },
    created() {
      this.updateMetaTags();
    },
    methods: {
      updateMetaTags() {
        // Update document title
        document.title = this.title;
        
        // Get existing meta tags
        let metaTags = document.querySelectorAll('meta');
        let linkTags = document.querySelectorAll('link[rel="canonical"]');
        
        // Helper function to update or create meta tag
        const updateOrCreateMetaTag = (name, content, property = null) => {
          // Check if meta tag exists
          let metaTag = null;
          if (name) {
            metaTag = document.querySelector(`meta[name="${name}"]`);
          } else if (property) {
            metaTag = document.querySelector(`meta[property="${property}"]`);
          }
          
          if (metaTag) {
            // Update existing tag
            metaTag.setAttribute('content', content);
          } else {
            // Create new tag
            metaTag = document.createElement('meta');
            if (name) {
              metaTag.setAttribute('name', name);
            }
            if (property) {
              metaTag.setAttribute('property', property);
            }
            metaTag.setAttribute('content', content);
            document.head.appendChild(metaTag);
          }
        };
        
        // Update basic meta tags
        updateOrCreateMetaTag('description', this.description);
        updateOrCreateMetaTag('keywords', this.keywords);
        updateOrCreateMetaTag('author', this.author);
        
        // Update Open Graph meta tags
        updateOrCreateMetaTag(null, this.title, 'og:title');
        updateOrCreateMetaTag(null, this.description, 'og:description');
        updateOrCreateMetaTag(null, this.type, 'og:type');
        updateOrCreateMetaTag(null, this.url || window.location.href, 'og:url');
        updateOrCreateMetaTag(null, this.image, 'og:image');
        updateOrCreateMetaTag(null, 'sv_SE', 'og:locale');
        updateOrCreateMetaTag(null, 'Svensk Hälsovård', 'og:site_name');
        
        // Update Twitter Card meta tags
        updateOrCreateMetaTag('twitter:card', this.twitterCard);
        updateOrCreateMetaTag('twitter:title', this.title);
        updateOrCreateMetaTag('twitter:description', this.description);
        updateOrCreateMetaTag('twitter:image', this.image);
        
        // Update canonical link
        let canonicalLink = document.querySelector('link[rel="canonical"]');
        if (canonicalLink) {
          canonicalLink.setAttribute('href', this.url || window.location.href);
        } else {
          canonicalLink = document.createElement('link');
          canonicalLink.setAttribute('rel', 'canonical');
          canonicalLink.setAttribute('href', this.url || window.location.href);
          document.head.appendChild(canonicalLink);
        }
        
        // Add Schema.org JSON-LD
        this.updateSchemaOrgData();
      },
      
      updateSchemaOrgData() {
        // Create schema data object
        const schema = {
          '@context': 'https://schema.org',
          '@type': this.schemaType,
          name: this.title,
          description: this.description,
          url: this.url || window.location.href,
          ...this.schemaData
        };
        
        // Find existing script tag
        let scriptTag = document.querySelector('script[type="application/ld+json"]');
        
        if (scriptTag) {
          // Update existing tag
          scriptTag.textContent = JSON.stringify(schema);
        } else {
          // Create new script tag
          scriptTag = document.createElement('script');
          scriptTag.setAttribute('type', 'application/ld+json');
          scriptTag.textContent = JSON.stringify(schema);
          document.head.appendChild(scriptTag);
        }
      }
    }
  };
  </script>