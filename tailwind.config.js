/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./ui/html/**/*.html",
  ],
  theme: {
    fontSize: {
      '2xs': '0.6rem',
    },
    extend: {
      backgroundImage: {
        'orange-linear': 'linear-gradient(-7deg, #EC2C5A, #FA7C54)',
      },
      maxWidth: {
        '3xl': '1800px',
      },
      colors: {
        'gray-icon': '#8A8A8A'
      }
    },
  },
  plugins: [],
}

