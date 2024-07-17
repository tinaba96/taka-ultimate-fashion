/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/*.{js,jsx,ts,tsx}',
    './src/**/*.{js,jsx,ts,tsx}',
    './public/index.html',
    './node_modules/react-tailwindcss-select/dist/index.esm.js'
  ],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/aspect-ratio')],
}

