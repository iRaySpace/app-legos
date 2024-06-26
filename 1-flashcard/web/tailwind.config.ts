/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './pages/**/*.{vue,js,ts}',
    './components/**/*.{vue,js,ts}',
  ],
  plugins: [
    require('daisyui'),
  ],
  theme: {
    extend: {
      fontFamily: {
        'poppins': ['poppins', 'sans-serif'],
      },
    },
  },
  daisyui: {
    themes: ["light"],
  },
}

