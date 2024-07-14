<template>
  <main ref="desktopViewEl">
    <template v-for="process in store.processes.values()" :key="process.id">
      <WindowsXPWindowComponent :options="process.windowOptions">
        <component :is="process.program.component" v-bind="process.program.props"></component>
      </WindowsXPWindowComponent>
    </template>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WindowsXPWindowComponent from '../components/WindowsXPWindowComponent.vue'
import { useContext } from '../composables/useContext'
import { useStore } from '../stores'

const store = useStore()
const desktopViewEl = ref()
useContext(desktopViewEl, {})
</script>

<style scoped lang="scss">
main {
  width: 100%;
  height: 100%;
}
</style>
