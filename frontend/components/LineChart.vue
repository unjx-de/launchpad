<script setup lang="ts">
import { PropType } from "@vue/runtime-core";
import { scaleLinear, select, curveBasis } from "d3";
import { line } from "d3-shape";
import { PercentageArraySize } from "~/types";

const props = defineProps({
  data: { type: Array as PropType<number[]>, default: [] },
  name: { type: String, default: "" },
  color: { type: String, default: "white" },
});
const myChart = ref();
let curve: any;
let x: any;
let y: any;

watch(props, () => {
  update();
});

const myLine = line()
  .x((_, i) => x(i))
  .y((d) => y(d))
  .curve(curveBasis);

function update() {
  x = scaleLinear()
    .domain([0, PercentageArraySize - 2])
    .range([0, myChart.value.clientWidth]);
  y = scaleLinear().domain([0, 100]).range([myChart.value.clientHeight, 0]);
  curve
    .datum(props.data)
    .attr("d", myLine)
    .attr("transform", null)
    .transition()
    .attr("transform", "translate(" + x(-1) + ")");
}

onMounted(() => {
  const width = myChart.value.clientWidth;
  const height = myChart.value.clientHeight;
  x = scaleLinear()
    .domain([0, PercentageArraySize - 2])
    .range([0, width]);
  y = scaleLinear().domain([0, 100]).range([height, 0]);

  const svg = select("#myChart" + props.name)
    .append("svg")
    .attr("width", width)
    .attr("height", height);

  curve = svg.append("path").datum(props.data).attr("fill", "none").attr("stroke", props.color).attr("stroke-width", 2).attr("d", myLine);
});
</script>

<template>
  <div :id="'myChart' + name" ref="myChart" class="w-full h-14 rounded-md overflow-hidden"></div>
</template>
