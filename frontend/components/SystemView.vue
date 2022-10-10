<script setup lang="ts">
import { PropType } from "@vue/runtime-core";
import { system_LiveInformation, system_StaticInformation } from "~/services/openapi";
import { useWebSocket } from "@vueuse/core";
import { emptyLiveSystem, emptyStaticSystem, PercentageArraySize } from "~/types";
import { computed } from "@vue/reactivity";
import { staticBlock } from "@babel/types";

const config = useRuntimeConfig();
const props = defineProps({
  liveSystem: { type: Object as PropType<system_LiveInformation>, default: emptyLiveSystem },
  staticSystem: { type: Object as PropType<system_StaticInformation>, default: emptyStaticSystem },
});
const { status, data: wsLiveSystem } = useWebSocket<string>(config.public.wsBaseUrl + "api/system/ws", {
  autoClose: true,
  autoReconnect: true,
});
const parsedData = computed<system_LiveInformation>(() => {
  if (wsLiveSystem.value && status.value === "OPEN") {
    return JSON.parse(wsLiveSystem.value);
  } else {
    return props.liveSystem;
  }
});
const cpuData = computed<number[]>(() => parsedData.value.cpu.percentage);
const cpuValue = computed<string>(() => cpuData.value[PercentageArraySize - 1] + " %");
const ramValue = computed<string>(() => parsedData.value.ram.value + "/" + props.staticSystem?.ram.readable);
const diskValue = computed<string>(() => parsedData.value.disk.value + "/" + props.staticSystem?.disk.readable);
</script>

<template>
  <div class="mb-6 md:mb-10">
    <div class="flex items-center bg-slate-200 dark:bg-slate-900 rounded-md relative">
      <SystemCard icon="cpu" color="crimson" :info="cpuValue" />
      <SystemCard icon="ram" color="steelblue" :info="ramValue" cardStyle="right-[7rem] hidden sm:flex" />
      <SystemCard icon="disk" color="forestgreen" :info="diskValue" cardStyle="right-[17rem] hidden sm:flex" />
      <LineChart name="cpu" :data="cpuData" color="crimson" />
    </div>
  </div>
</template>
