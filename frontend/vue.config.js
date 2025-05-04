module.exports = {
    publicPath: '/',
    productionSourceMap: false,
    css: {
      loaderOptions: {
        scss: {
          additionalData: `@import "@/assets/styles/_variables.scss";`
        }
      }
    },
    chainWebpack: config => {
      config.plugin('html').tap(args => {
        args[0].title = 'Svensk Hälsovård - Din väg till bättre hälsa';
        args[0].meta = {
          description: 'Boka blodprov och hälsokontroller enkelt och snabbt. Vi erbjuder professionella hälsotjänster för alla.',
          viewport: 'width=device-width,initial-scale=1,minimum-scale=1',
          robots: 'index, follow',
          'og:title': 'Svensk Hälsovård - Din väg till bättre hälsa',
          'og:description': 'Boka blodprov och hälsokontroller enkelt och snabbt. Vi erbjuder professionella hälsotjänster för alla.',
          'og:type': 'website',
          'og:url': 'https://svenskhalsovard.se',
          'og:image': 'https://svenskhalsovard.se/img/og-image.jpg'
        };
        return args;
      });
    },
    devServer: {
      proxy: {
        '/api': {
          target: 'http://localhost:8080',
          changeOrigin: true
        }
      }
    }
  };