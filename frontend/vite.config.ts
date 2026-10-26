import { defineConfig } from "vite";

// Proxy /api to the Go server so the browser sees one origin (no CORS setup).
export const apiTarget = "http://localhost:8080";

export default defineConfig({
  server: {
    proxy: {
      "/api": apiTarget,
    },
  },
});
