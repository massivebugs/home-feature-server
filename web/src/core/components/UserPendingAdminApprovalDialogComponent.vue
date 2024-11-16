<template>
  <DialogComponent
    class="hfs-user-pending-admin-approval-dialog"
    pos="center"
    :fit-content="true"
    :buttons="{
      success: false,
      cancel: true,
    }"
    :controls="{
      close: false,
      minimize: false,
      maximize: false,
    }"
    :resizable="false"
  >
    <template #title>
      <LockIconComponent />
      {{ t('userPendingAdminApproval.title') }}
    </template>
    <div class="hfs-user-pending-admin-approval-dialog__container">
      <p class="hfs-user-pending-admin-approval-dialog__create-user-disclaimer">
        {{ t('userPendingAdminApproval.message') }}
      </p>
    </div>
  </DialogComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DialogComponent from '@/core/components/DialogComponent.vue'
import type { RelativePosition } from '@/core/models/relativePosition'
import LockIconComponent from './LockIconComponent.vue'

defineProps<{
  pos?: RelativePosition | 'center'
}>()

const { t } = useI18n()
</script>

<style scoped lang="scss">
@use '@/assets/colors';
.hfs-user-pending-admin-approval-dialog {
  transition: box-shadow 0.2s;
}

.hfs-user-pending-admin-approval-dialog:hover {
  box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -webkit-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
  -moz-box-shadow: 1px 1px 25px 2px rgba(0, 0, 0, 0.4) !important;
}

.hfs-user-pending-admin-approval-dialog__container {
  width: 100%;
  margin-bottom: 1em;
  display: flex;
  flex-direction: column;
  gap: 0.5em;
}

.hfs-user-pending-admin-approval-dialog__create-user-disclaimer {
  white-space: pre-line;
  margin: 0;
  font-size: 0.9em;
  text-align: center;
  line-height: 1.6em;
}
</style>
