<script setup lang="ts">
import {
  bookmark_Bookmark,
  BookmarksService,
  OpenAPI,
  system_LiveInformation,
  system_StaticInformation,
  SystemService,
  weather_OpenWeatherApiResponse,
  WeatherService,
} from "~/services/openapi";
import AuthLogin from "~/components/AuthLogin.vue";
import { ref } from "@vue/reactivity";

const config = useRuntimeConfig();
OpenAPI.BASE = config.public.baseUrl + "api";
OpenAPI.WITH_CREDENTIALS = true;
const loading = ref(true);
const fetchError = ref(false);
const bookmarks = ref<bookmark_Bookmark[]>([]);
const weather = ref<weather_OpenWeatherApiResponse>();
const liveSystem = ref<system_LiveInformation>();
const staticSystem = ref<system_StaticInformation>();

function getData() {
  loading.value = true;
  fetchError.value = false;
  BookmarksService.getBookmarks()
    .then(async (res) => {
      bookmarks.value = res;
      weather.value = await WeatherService.getWeather();
      liveSystem.value = await SystemService.getSystemLive();
      staticSystem.value = await SystemService.getSystemStatic();
      loading.value = false;
    })
    .catch(() => {
      loading.value = false;
      fetchError.value = true;
    });
}
getData();
</script>

<template>
  <div class="mx-auto max-w-7xl px-5 lg:px-8 my-3 md:my-10 lg:my-18 xl:my-28">
    <Transition mode="out-in">
      <div v-if="!loading && !fetchError">
        <WeatherCard v-if="weather" :weather="weather" />
        <SystemView v-if="liveSystem && staticSystem" :liveSystem="liveSystem" :static-system="staticSystem" />
        <div v-if="bookmarks && bookmarks.length > 0">
          <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-5">
            <BookmarkCard v-for="bookmark in bookmarks" :key="bookmark.name" :bookmark="bookmark" />
          </div>
        </div>
        <EmptyList v-else />
      </div>
      <AuthLogin v-else-if="!loading" @try="getData" />
      <WaveSpinner v-else />
    </Transition>
  </div>
</template>

<style>
body {
  margin: 0;
  padding: 0;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.2s ease;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
