import {optimizeImports} from "carbon-preprocess-svelte";
import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: [optimizeImports()],
    kit: {adapter: adapter(), prerender: {default: true}}
};

export default config;
