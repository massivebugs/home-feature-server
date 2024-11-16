<template>
  <SystemSettingsWindowComponent :api="api" />
</template>

<script setup lang="ts">
import SystemSettingsWindowComponent from '../components/SystemSettingsWindowComponent.vue'
import { useAPI } from '../composables/useAPI'
import { API_URL } from '../constants'

const api = useAPI(API_URL)
</script>
