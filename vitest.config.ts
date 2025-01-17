import {coverageConfigDefaults, defineConfig} from 'vitest/config'

export default defineConfig({
  test: {
    globalSetup: './src/test-globals.ts',
    coverage: {
      provider: "v8",
      exclude: [
        'build/**',
        '.svelte-kit/**',
        '*.config.*',
        ...coverageConfigDefaults.exclude
      ]
    }
  },
})