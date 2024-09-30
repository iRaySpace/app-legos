/** @type {import('tailwindcss').Config} */
module.exports = {
    plugins: [
        require('daisyui'),
    ],
    content: [
        './pages/**/*.{vue,js,ts}',
        './components/**/*.{vue,js,ts}',
    ],
    daisyui: {
        themes: [
            'system',
            'light',
            'dark',
        ],
    },
}