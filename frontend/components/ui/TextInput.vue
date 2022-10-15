<script setup lang="ts">
interface Props {
  inputClass?: string;
  modelValue: string;
  placeholder?: string;
  type?: string;
  wrong?: boolean;
  disabled?: boolean;
}
withDefaults(defineProps<Props>(), {
  inputClass: "",
  placeholder: "",
  type: "text",
  wrong: false,
  disabled: false,
});
const emit = defineEmits(["update:modelValue", "enter", "esc"]);
function updateValue(e: any) {
  emit("update:modelValue", e.target.value);
}
</script>

<template>
  <input
    :disabled="disabled"
    autofocus
    @input="updateValue"
    @keyup.enter="$emit('enter')"
    @keyup.esc="$emit('esc')"
    :type="type"
    :value="modelValue"
    :class="`${inputClass} ${wrong ? 'border-red-500' : 'focus:border-blue-500'}`"
    class="px-3 py-1.5 rounded border border-2 border-gray-300 disabled:bg-gray-300 outline-none text-gray-800"
    :id="placeholder"
    :placeholder="placeholder"
  />
</template>
