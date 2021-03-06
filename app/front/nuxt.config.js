const pkg = require("./package");

const environment = process.env.NODE_ENV || 'development';
const envSet = {
    production: {
        SSSIGNAL_API_DOMAIN: "https://api.sssignal.com"
    },
    development: {
        SSSIGNAL_API_DOMAIN: "http://host.docker.internal:1323"
    }
}

module.exports = {
  ssr: true,

  /*
   ** Headers of the page
   */
  head: {
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },

  /*
   ** Customize the progress-bar color
   */
  loading: { color: "#fff" },

  /*
   ** Global CSS
   */
  css: [],

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],

  /*
   ** Nuxt.js modules
   */
  modules: [
    [
      "@nuxtjs/google-analytics",
      {
        id: "UA-79013420-3"
      }
    ]
  ],

  /*
   ** Build configuration
   */
  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  },

  env: envSet[environment]
};
