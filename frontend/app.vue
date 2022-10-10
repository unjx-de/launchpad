<script setup lang="ts">
import { bookmark_Bookmark, weather_OpenWeatherApiResponse } from "~/services/openapi";

const config = useRuntimeConfig();
const {
  pending: pendingBookmarks,
  data: bookmarks,
  error: bookmarksError,
} = useLazyFetch<bookmark_Bookmark[]>(config.public.baseUrl + "api/bookmarks", { default: () => [] });
const { pending: pendingWeather, data: weather, error: weatherError } = useLazyFetch<weather_OpenWeatherApiResponse>(config.public.baseUrl + "api/weather");
const {
  pending: pendingLiveSystem,
  data: liveSystem,
  error: liveError,
} = useLazyFetch<weather_OpenWeatherApiResponse>(config.public.baseUrl + "api/system/live");
const {
  pending: pendingStaticSystem,
  data: staticSystem,
  error: staticError,
} = useLazyFetch<weather_OpenWeatherApiResponse>(config.public.baseUrl + "api/system/static");

const doneLoading = computed(() => !pendingBookmarks.value && !pendingWeather.value && !pendingLiveSystem.value && !pendingStaticSystem.value);
const noError = computed(() => !bookmarksError.value && !weatherError.value && !liveError.value && !staticError.value);
</script>

<template>
  <div class="mx-auto max-w-7xl px-5 lg:px-8 my-3 md:my-10 lg:my-18 xl:my-28">
    <Transition mode="out-in">
      <div v-if="doneLoading && noError">
        <WeatherCard v-if="weather" :weather="weather" />
        <SystemView v-if="liveSystem && staticSystem" :liveSystem="liveSystem" :static-system="staticSystem" />
        <div v-if="bookmarks && bookmarks.length > 0">
          <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 gap-5">
            <BookmarkCard v-for="bookmark in bookmarks" :key="bookmark.id" :bookmark="bookmark" />
          </div>
        </div>
        <EmptyList v-else />
      </div>
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
