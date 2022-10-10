<script setup lang="ts">
import { PropType } from "@vue/runtime-core";
import { bookmark_Bookmark } from "~/services/openapi";
import { emptyBookmark } from "~/types";

const config = useRuntimeConfig();
const props = defineProps({ bookmark: { type: Object as PropType<bookmark_Bookmark>, default: emptyBookmark } });
const icon = computed(() => {
  if (props.bookmark.icon.startsWith("http")) {
    return props.bookmark.icon;
  } else {
    return config.public.baseUrl + "static/" + props.bookmark.icon;
  }
});
</script>

<template>
  <a
    :href="bookmark.url"
    class="link flex items-center rounded-md hover:underline underline-offset-2 decoration-gray-500 text-sm text-slate-700 dark:text-slate-300 hover:text-slate-900 dark:hover:text-slate-50 ease-in duration-200"
  >
    <div class="img rounded-md w-8 h-8 bg-cover bg-center opacity-90" :style="{ backgroundImage: `url(${icon})` }"></div>
    <div class="uppercase truncate ml-2">
      {{ bookmark.name }}
    </div>
  </a>
</template>

<style scoped>
.link:hover .img {
  opacity: 1;
  transition: opacity ease-in 0.2s;
}
</style>
