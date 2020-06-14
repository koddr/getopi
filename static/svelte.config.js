const sveltePreprocess = require("svelte-preprocess");
const autoprefixer = require("autoprefixer");
const pimport = require("postcss-import");
const minmax = require("postcss-media-minmax");
const csso = require("postcss-csso");

module.exports = {
  preprocess: sveltePreprocess({
    postcss: {
      plugins: [pimport, minmax, autoprefixer, csso],
    },
  }),
};
