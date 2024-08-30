<template>
  <WindowComponent
    :size="new RelativeSize(70, 80)"
    :title="t('portfolio.title')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="toolbarOptions"
    title-bar-icon="/images/portfolio_icon_small.png"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :resizable="true"
    @click-close="emit('clickClose')"
  >
    <div class="hfs-portfolio__container">
      <h1>Portfolio</h1>
      <p>A professional looking photo</p>
      <p>Da Hyun Kim</p>
      <p>dhkp443@gmail.com</p>
      <p>https://github.com/massivebugs</p>
      <p>
        Hi, I'm Da Hyun. I'm a software engineer who loves making neat backend systems and unique
        frontend experiences.
      </p>
      <h2>Professional Experiences</h2>
      <div>
        <p>...</p>
        <p>full-stack engineer to develop and maintain features</p>
        <p>Projects</p>
        <ul>
          <li>LEAN BODY</li>
        </ul>
        <p>...</p>
        <p>Did various projects for clients, made company product</p>
        <p>Projects</p>
        <ul>
          <li>...</li>
          <li>internal system cms api</li>
          <li>mobile app for car sharing</li>
        </ul>
      </div>
      <h2>Personal Projects</h2>
      <ul>
        <li>
          Home Feature Server / Portfolio website
          <p>
            Initially created for writing utility apps for home usage, I decided to expand it as a
            personal website as well
          </p>
        </li>
        <li>
          Cashbunny Budget Planner
          <p>
            A simple accounting system with transaction scheduling and budget planning/visualization
          </p>
        </li>
        <li>
          TELEBYTE
          <p>A desktop telepresence robot controlled via WebRTC</p>
          <p>Made using ESP32 WROOM 32, servo controller module, servos, brushed dc motors etc</p>
        </li>
        <li>
          Cirro Virtual assistant
          <p>
            A simple virtual assistant designed to provide support to various features and
            applications for this server
          </p>
          <p>Made using a behavior tree and a knowledge graph</p>
        </li>
        <li>
          Document Editor
          <p>A simple document editor based on the TipTap API</p>
        </li>
        <li>
          Blog writeups
          <p>Check out this directory for writeups</p>
        </li>
        <li>
          Video Conferencing
          <p>So that I can call my parents who live far away</p>
        </li>
      </ul>
      <h2>CTA</h2>
      <div>
        <div>Contact me via Email</div>
        <div>Or Call me now! (available times -)</div>
        <div>In the meantime, check out my GitHub!</div>
      </div>
    </div>
    <ContactFormDialogComponent
      v-if="showContactFormDialog"
      pos="center"
      @submit="onClickSuccessContactFormDialog"
      @click-close="onClickCloseContactFormDialog"
    />
    <AboutDialogComponent
      v-if="showAboutDialog"
      pos="center"
      @click-close="onClickCloseAboutDialog"
    />
  </WindowComponent>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import WindowComponent from '@/core/components/WindowComponent.vue'
import type { WindowToolbarRow } from '@/core/components/WindowToolbarComponent.vue'
import { RelativePosition } from '@/core/models/relativePosition'
import { RelativeSize } from '@/core/models/relativeSize'
import AboutDialogComponent from './AboutDialogComponent.vue'
import ContactFormDialogComponent, {
  type ContactFormSuccessEvent,
} from './ContactFormDialogComponent.vue'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const showContactFormDialog = ref<boolean>(false)
const showAboutDialog = ref<boolean>(false)
const toolbarOptions = computed<WindowToolbarRow[]>(() => [
  {
    isMenu: true,
    items: [
      {
        label: 'File',
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('portfolio.exit'),
                shortcutKey: 'Alt+F4',
                isDisabled: false,
                onClick: () => {
                  emit('clickClose')
                },
              },
            ],
          ],
        },
      },
      {
        label: t('portfolio.help'),
        contextMenuOptions: {
          itemGroups: [
            [
              {
                label: t('portfolio.contact'),
                isDisabled: false,
                onClick: () => {
                  showContactFormDialog.value = true
                },
              },
            ],
            [
              {
                label: t('portfolio.about'),
                isDisabled: false,
                onClick: () => {
                  showAboutDialog.value = true
                },
              },
            ],
          ],
        },
      },
    ],
  },
])

const onClickSuccessContactFormDialog = (payload: ContactFormSuccessEvent) => {
  console.log('contact form success', payload)
}

const onClickCloseContactFormDialog = () => {
  showContactFormDialog.value = false
}

const onClickCloseAboutDialog = () => {
  showAboutDialog.value = false
}
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.hfs-portfolio__container {
  width: 100%;
  height: 100%;
  padding: 5px;
  background-color: colors.$light-grey;
}
</style>
