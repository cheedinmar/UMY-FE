/** @type {import('tailwindcss').Config} */
export default {
   content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    
    extend: {
      colors:{
        primary:'#10093D',
        secondary:'#FFA500',
        cardColor1 :'#0047AB',
        cardColor2 :'#0E3770'
      },
      backgroundImage: {
        'background-pattern': "url('/src/assets/background.svg')",
      }
    },
  },
  plugins: [],
}

