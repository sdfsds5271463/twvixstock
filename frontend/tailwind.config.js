module.exports = {
    content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    darkMode: 'class', //'media'(隨系統設定) or 'class'(隨html的class="dark") 
    theme: {
        extend: {
            colors: {
                transparent: 'transparent',
                current: 'currentColor',
                primary: '#8b5cf6',
                success: '#10b981',
                danger: '#ef4444',
                warning: '#f59e0b',
                info: '#3b82f6',
                light: '#6b7280',
                dark: '#111827',
            },
        },
    },
    variants: {
        extend: {},
    },
    plugins: [],
}
