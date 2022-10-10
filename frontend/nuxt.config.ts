// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  typescript: { strict: true },
  modules: ["@nuxtjs/tailwindcss", "@vueuse/nuxt"],
  runtimeConfig: {
    public: {
      baseUrl: process.env.BASE_URL || "http://localhost:4000/",
      wsBaseUrl: (process.env.BASE_URL && process.env.BASE_URL.replace("http", "ws")) || "ws://localhost:4000/",
    },
  },
  app: {
    head: {
      bodyAttrs: {
        class: "bg-slate-100 dark:bg-slate-800 text-slate-800 dark:text-slate-50",
      },
      htmlAttrs: {
        lang: "en",
      },
      title: "Dashboard",
      meta: [
        { name: "viewport", content: "width=device-width, initial-scale=1" },
        { name: "description", content: "dashboard with bookmarks" },
        { name: "msapplication-TileColor", content: "#2e323e" },
        { name: "theme-color", content: "#2e323e" },
      ],
      link: [
        { rel: "icon", type: "image/x-icon", href: "/favicon.ico" },
        { rel: "apple-touch-icon", sizes: "180x180", href: "/favicon/apple-touch-icon.png" },
        { rel: "icon", type: "image/png", sizes: "32x32", href: "/favicon/favicon-32x32.png" },
        { rel: "icon", type: "image/png", sizes: "16x16", href: "/favicon/favicon-16x16.png" },
        { rel: "manifest", href: "/favicon/site.webmanifest" },
        { rel: "mask-icon", href: "/favicon/safari-pinned-tab.svg", color: "#2e323e" },
      ],
    },
  },
  ssr: false,
});
