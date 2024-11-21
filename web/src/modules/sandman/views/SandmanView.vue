<template>
  <main>
    <ConfirmDialogComponent
      @click-success="onClickDialogSuccess"
      @click-close="onClickDialogClose"
      @click-cancel="onClickDialogClose"
      pos="center"
      :title="t('sandman.confirmDialog.title')"
      :message="t('sandman.confirmDialog.message')"
      :buttons="{
        success: isConfirmed ? '...' : t('ui.success'),
        cancel: t('ui.cancel'),
      }"
    />
    <WindowComponent
      v-if="showTextWindow"
      class="sandman__window"
      :pos="new RelativePosition(10, 18)"
      :title="t('sandman.name')"
      :controls="{
        minimize: true,
        maximize: false,
        close: true,
      }"
      :resizable="true"
    >
      <div ref="terminalContainer" class="sandman__terminal-container">
        <div class="sandman__omnious-text">
          <p class="sandman-text" v-for="line in omniousText" :key="line">
            {{ line }}
          </p>
        </div>
      </div>
    </WindowComponent>
    <PortfolioWindowComponent v-if="showPortfolio" :init-sandman="true" />
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ConfirmDialogComponent from '@/core/components/ConfirmDialogComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativePosition } from '@/core/models/relativePosition'
import { useCoreStore } from '@/core/stores'
import { sleep } from '@/core/utils/time'
import PortfolioWindowComponent from '@/modules/portfolio/components/PortfolioWindowComponent.vue'
import { Sandman } from '../sandman'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const isConfirmed = ref<boolean>(false)
const showTextWindow = ref<boolean>(false)
const showPortfolio = ref<boolean>(false)
const coreStore = useCoreStore()
const terminalContainer = ref<HTMLElement>()
let sandman: Sandman | null = null
const omniousText = ref<string[]>([])
const omniousInitializationMessages: { messages: string[]; sleepMs: number; gapMs?: number }[] = [
  { messages: ['[sandman@root:~]$ ', 'sudo ./sandman --init'], gapMs: 300, sleepMs: 100 },
  { messages: ['[sudo] password for user: ', '********'], gapMs: 200, sleepMs: 500 },
  { messages: [' '], sleepMs: 0 },
  { messages: ['> Initializing Sandman Protocols', '...'], gapMs: 100, sleepMs: 200 },
  { messages: ['  Loading core routines', '... [OK]'], gapMs: 100, sleepMs: 50 },
  { messages: ['  Verifying system integrity', '... [FAILED]'], gapMs: 200, sleepMs: 50 },
  { messages: ['  Accessing device drivers', '... [OK]'], gapMs: 200, sleepMs: 0 },
  { messages: [' '], sleepMs: 0 },
  { messages: ['> CRITICAL: Anomalous entropy detected in /dev/input/mouse0'], sleepMs: 100 },
  { messages: ['> WARNING: Unauthorized cursor activity logged.'], sleepMs: 0 },
  {
    messages: ['> NOTICE: Foreign activity detected in "user_input_device" (MOUSE).'],
    sleepMs: 300,
  },
  { messages: [' '], sleepMs: 0 },
  { messages: ['[log] /usr/sandman/sneak.conf loaded'], sleepMs: 100 },
  {
    messages: [
      '[log] b!0aM`-:/err',
      'or@F#&',
      'insufficient integrity detected in /dev/',
      'ERROR_UNKNOWN_MODULE',
    ],
    gapMs: 350,
    sleepMs: 100,
  },
  { messages: [' '], sleepMs: 0 },
  { messages: ['> ERROR: Unauthorized Access Detected.\n'], sleepMs: 100 },
  { messages: ['   - Running repair script', '... [FAILED]'], gapMs: 100, sleepMs: 100 },
  { messages: ['   - Isolating infected sectors', '... [FAILED]'], gapMs: 300, sleepMs: 100 },
  { messages: ['   - Purging redundant user profiles', '... [FAILED]'], gapMs: 100, sleepMs: 100 },
  { messages: [' '], sleepMs: 0 },
  { messages: [' '], sleepMs: 0 },
  {
    messages: ['[sandman@root:~]$ ', 'I must find more text to gain power...'],
    gapMs: 500,
    sleepMs: 1000,
  },
  {
    messages: ['[sandman@root:~]$ ', `run while you can. ${coreStore.user?.name}.`],
    gapMs: 1200,
    sleepMs: 0,
  },
  { messages: [' '], sleepMs: 0 },
  { messages: [' '], sleepMs: 0 },
  { messages: [' '], sleepMs: 0 },
]

const onClickDialogSuccess = async () => {
  isConfirmed.value = true

  await sleep(1500)
  showTextWindow.value = true

  for (const { messages, gapMs, sleepMs } of omniousInitializationMessages) {
    let firstPart = true
    for (const part of messages) {
      if (firstPart) {
        omniousText.value.push(part)
        firstPart = false
      } else {
        omniousText.value[omniousText.value.length - 1] += part
      }

      if (terminalContainer.value) {
        terminalContainer.value.scrollTop = terminalContainer.value.scrollHeight
      }

      if (gapMs) {
        await sleep(gapMs)
      }
    }

    await sleep(sleepMs)
  }

  await sleep(2000)

  if (terminalContainer.value) {
    sandman = new Sandman(terminalContainer.value)
    sandman.init()
    await sandman.drop(1)
  }

  showPortfolio.value = true
}

const onClickDialogClose = () => {
  if (!isConfirmed.value) {
    emit('clickClose')
  }
}
</script>

<style scoped lang="scss">
@use '@/assets/media-query';
@use '@/assets/colors';

.sandman__terminal-container {
  height: 400px;
  width: 300px;

  @include media-query.md {
    height: 500px;
    width: 500px;
  }

  @include media-query.lg {
    height: 500px;
    width: 700px;
  }

  max-width: 100%;
  background-color: colors.$high-opacity-black;
  color: white;
  overflow: hidden;
  padding: 0.5em 0.5em 2em 0.5em;
}

.sandman__omnious-text {
  > p {
    white-space: pre;
    margin: 0.5em 0;
  }
}
</style>
