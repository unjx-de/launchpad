/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ["nuxt.config.ts"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography")],
};
