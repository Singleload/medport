<template>
    <div id="app">
      <AppNavbar />
      <main>
        <router-view />
      </main>
      <AppFooter />
    </div>
  </template>
  
  <script>
  import AppNavbar from '@/components/common/AppNavbar.vue';
  import AppFooter from '@/components/common/AppFooter.vue';
  
  export default {
    name: 'App',
    components: {
      AppNavbar,
      AppFooter
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