import adapter from '@sveltejs/adapter-node';
import {vitePreprocess} from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    extensions: ['.svelte'],
    // Consult https://kit.svelte.dev/docs/integrations#preprocessors
    // for more information about preprocessors
    preprocess: [vitePreprocess()],

    kit: {
        adapter: adapter(),
        csp: {
            directives: {
                'script-src': ['self'],
                'script-src-elem': ['self'],
                'style-src': ['self', "unsafe-inline"],
                'frame-ancestors': ['none'],
            },
        }
    }
};
export default config;