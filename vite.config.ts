import {sveltekit} from "@sveltejs/kit/vite";
import {defineConfig, searchForWorkspaceRoot} from "vite";
import {wuchale} from "@wuchale/vite-plugin";

export default defineConfig({
  server: {
    host: "localhost",
    fs: {
      allow: [
        searchForWorkspaceRoot(process.cwd()),
        'diamond-planner/ui'
      ]
    }
  },
  plugins: [wuchale(), sveltekit()],
});
