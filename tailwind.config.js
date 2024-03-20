/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./templates/**/*.templ"],
    theme: {
        extend: {},
    },
    daisyui: {
        themes: ["light"],
    },
    plugins: [require("daisyui")],
}