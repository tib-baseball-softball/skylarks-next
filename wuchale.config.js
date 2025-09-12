// @ts-check
import {adapter as svelte} from "@wuchale/svelte";
import {defineConfig} from "wuchale";

export default defineConfig({
    sourceLocale: "en",
    otherLocales: ['de', 'fr', 'es', 'pl', 'ru'],
    adapters: {
        main: svelte({
            // writeFiles: {
            //     compiled: true,
            //     proxy: true,
            // }
        }),
    },
});