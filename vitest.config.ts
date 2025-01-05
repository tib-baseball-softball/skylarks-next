import {coverageConfigDefaults, defineConfig} from 'vitest/config'

export default defineConfig({
  test: {
    coverage: {
      exclude: [
        'build/**',
        '.svelte-kit/**',
        '*.config.*',
        ...coverageConfigDefaults.exclude
      ]
    }
  },
})