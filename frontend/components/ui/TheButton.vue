<script setup lang="ts">
interface Props {
  btnStyle?: string;
  color?: string;
  label?: string;
  disabled?: boolean;
}
const props = withDefaults(defineProps<Props>(), {
  btnStyle: "",
  color: "blue",
  label: "",
  disabled: false,
});
const bgColor = {
  blue: "bg-blue-600 hover:bg-blue-700 active:bg-blue-800 disabled:bg-blue-400 disabled:cursor-not-allowed ",
  red: "bg-red-600 hover:bg-red-700 active:bg-red-800 disabled:bg-red-400 disabled:cursor-not-allowed ",
};
defineEmits(["click"]);
</script>

<template>
  <button
    :disabled="disabled"
    @click="$emit('click')"
    class="text-white shadow-md uppercase transition duration-150 ease-in-out px-3 py-2 leading-tight rounded"
    :class="bgColor[color] + btnStyle"
  >
    {{ label }}
  </button>
</template>
