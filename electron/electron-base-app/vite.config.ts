import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  //base:path.resolve(__dirname,'./dist/'),
  base:"./",
  plugins: [vue()]
})
