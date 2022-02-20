import {optimizeImports} from "carbon-preprocess-svelte";
import adapter from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    // Consult https://github.com/sveltejs/svelte-preprocess
    // for more information about preprocessors
    preprocess: [optimizeImports()],

    kit: {
        // hydrate the <div id="svelte"> element in src/app.html
        target: '#svelte',

        adapter: adapter({
            // default options are shown
            out: 'build',
            precompress: false,
            env: {
                path: 'SOCKET_PATH',
                host: 'HOST',
                port: 'PORT',
                origin: 'ORIGIN',
                headers: {
                    protocol: 'PROTOCOL_HEADER',
                    host: 'HOST_HEADER'
                }
            }
        })
    }
};

export default config;
