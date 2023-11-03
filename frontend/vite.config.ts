import react from '@vitejs/plugin-react'
import * as path from 'path'
import { defineConfig } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'
// https://vitejs.dev/config/
export default defineConfig({
  build: {
    assetsDir: "assets",
    base: "./",
  },
  plugins: [
    react({
      jsxImportSource: '@emotion/react',
      babel: {
        plugins: ['@emotion/babel-plugin'],
      },
    }),
    tsconfigPaths(),
  ],
  envDir: path.resolve(__dirname, 'config/env'),
})
