<template>
    <div id="app">
      <Navbar />
      <main>
        <router-view />
      </main>
      <Footer />
    </div>
  </template>
  
  <script>
  import Navbar from '@/components/common/Navbar.vue';
  import Footer from '@/components/common/Footer.vue';
  
  export default {
    name: 'App',
    components: {
      Navbar,
      Footer
    }
  };
  </script>
  
  <style lang="scss">
  #app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    
    main {
      flex: 1;
      padding: 2rem 0;
    }
  }
  </style>