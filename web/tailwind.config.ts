// tailwindcss ts配置
// https://tailwindcss.com/docs/guides/nextjs
import { Config } from "tailwindcss"
import { nextui } from "@nextui-org/theme"
const config: Config = {
    content: [
        "./index.html",
        "./src/**/*.{js,ts,jsx,tsx}",
        "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
    ],
    theme: {
        extend: {},
    },
    darkMode: "class",
    plugins: [nextui()],
}

export default config