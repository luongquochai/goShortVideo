module.exports = {
  content: ["./public/**/*.html", "./src/**/*.{html,js,templ}"],
  theme: {
    extend: {
      colors: {
        gray: {
          800: '#333333',
          300: '#666666',
        },
        white: '#FFFFFF',
      },
      fontFamily: {
        sans: ['Inter', 'system-ui', 'system'],
      },
      fontSize: {
        xl: '18px',
        lg: '16px',
        md: '14px',
        sm: '12px',
      },
    },
  },
  plugins: [],
};
