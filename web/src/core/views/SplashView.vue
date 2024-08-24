<template>
  <div class="splash">
    <p v-for="message in displayedMessages" :key="message">[ OK ] {{ message }}</p>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const displayedMessages = ref<string[]>([])

const loadingMessages: [string, number][] = [
  ['Initializing boot sequence...', 800],
  ['Loading modules...', 500],
  ['Starting system logger...', 100],
  ['Mounting local filesystems...', 100],
  ['Started login service', 100],
  ['Started network manager', 100],
  ['Loading user settings...', 100],
  ['Starting desktop environment...', 100],
  ['Applying system configurations...', 100],
  ['Checking for updates...', 1500],
  ['System startup complete.', 1500],
]

const displayLoadingMessages = () => {
  const nextMessage = loadingMessages.shift()
  if (nextMessage) {
    displayedMessages.value.push(nextMessage[0])
    setTimeout(displayLoadingMessages, nextMessage[1])
  } else {
    router.push({ name: 'login' })
  }
}

onMounted(() => {
  displayLoadingMessages()
  // TODO: Retrieve supported currencies
})
</script>

<style scoped lang="scss">
.splash {
  width: 100vw;
  height: 100vh;
}
</style>
