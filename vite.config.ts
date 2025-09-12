import {sveltekit} from '@sveltejs/kit/vite';
import {defineConfig} from 'vite';
import tailwindcss from "@tailwindcss/vite";
import {wuchale} from "@wuchale/vite-plugin";

export default defineConfig({
  server: {
    host: 'localhost'
  },
  plugins: [
    tailwindcss(),
    wuchale(),
    sveltekit(),
  ],
});
