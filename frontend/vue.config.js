const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  css: {
    loaderOptions: {
      sass: {
        additionalData: `
          @import "@/sass/main.scss";
        `
      }
    }
  },
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://localhost:3000',
      }
    }
  }
})
