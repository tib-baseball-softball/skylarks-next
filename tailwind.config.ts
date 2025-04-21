import {join} from 'path';
import type {Config} from 'tailwindcss';

export default {
  darkMode: 'media',
  content: [
    './src/**/*.{html,js,svelte,ts}',
    join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte,ts}')
  ],
  theme: {
    extend: {},
  },
  plugins: [
    // forms,
    // typography,
    // tailwindcssAnimate,
    // skeleton({
    //   themes: {
    //     custom: [
    //       skylarksTheme
    //     ]
    //   },
    // }),
  ],
} satisfies Config;
