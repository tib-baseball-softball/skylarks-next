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
                'default-src': ['self', '*.baseball-softball.de', '*.tib-baseball.de', '*.berlinskylarks.de'],
                'connect-src': ['self', 'https://*.baseball-softball.de', 'https://*.tib-baseball.de', 'https://*.berlinskylarks.de', 'http://127.0.0.1:8090'],
                'img-src': ['self', 'data: w3.org/svg/2000', 'https://*.baseball-softball.de', 'https://*.tib-baseball.de', 'https://*.berlinskylarks.de', 'http://127.0.0.1:8090', 'https://*.tile.openstreetmap.de'],
                'script-src': ['self'],
                'script-src-elem': ['self'],
                'style-src': ['self', "unsafe-inline"],
                'frame-ancestors': ['none'],
            },
        }
    }
};
export default config;