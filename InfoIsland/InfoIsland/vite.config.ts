import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': '/src',
        }
    },
    plugins: [react({})],
    server: {
        watch: {
            usePolling: true
        },
        proxy: {
            '/api/v1': {
                target: 'http://192.168.100.5:9000/api/v1',
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api\/v1/, '')
            },
        }
    }
})
