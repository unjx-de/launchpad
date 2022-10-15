<script setup lang="ts">
import TextInput from "~/components/ui/TextInput";
import { AuthService } from "~/services/openapi";
import { ref } from "@vue/reactivity";
import TheButton from "~/components/ui/TheButton.vue";

const emit = defineEmits(["try"]);
const password = ref("");
const wrong = ref(false);
const loading = ref(false);

function handleLogin() {
  if (password.value === "") return;
  loading.value = true;
  AuthService.postAuthLogin(password.value)
    .then(() => {
      emit("try");
    })
    .catch(() => {
      loading.value = false;
      wrong.value = true;
    });
}
</script>

<template>
  <div class="flex flex-col items-center">
    <div>
      <div class="flex w-full sm:w-80">
        <TextInput
          :disabled="loading"
          inputClass="w-full mr-2"
          v-model="password"
          placeholder="Password"
          type="password"
          @enter="handleLogin"
          @esc="password = ''"
          :wrong="wrong"
        />
        <TheButton :disabled="loading" label="login" :color="wrong ? 'red' : 'blue'" @click="handleLogin" />
      </div>
      <div v-if="wrong" class="text-sm text-red-500 ml-2">please try again</div>
    </div>
  </div>
</template>
