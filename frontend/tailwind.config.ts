import type { Config } from "tailwindcss";
import typography from '@tailwindcss/typography';

export default {
  content: ["./src/**/*.{html,js,svelte,ts}"],

  theme: {
    extend: {
      colors: {
        primary: '#681726',
        bg: '#ebdcc5',
        secondary: '#64748b',
        dark: '#0f172a' 
      }
    }
  },

  plugins: [typography]
} as Config;
