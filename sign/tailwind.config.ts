import type { Config } from "tailwindcss";

export default {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {}
  },
  darkMode: "class", // 需要开启 class 模式
  plugins: []
} satisfies Config;
