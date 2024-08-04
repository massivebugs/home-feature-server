<template>
  <WindowComponent
    :size="new RelativeSize(70, 80)"
    :title="t('portfolio.toolbarTitle')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    :toolbar="[
      {
        isMenu: true,
        items: [
          {
            label: t('portfolio.help'),
            contextMenuOptions: {
              itemGroups: [
                [
                  {
                    label: t('portfolio.about'),
                    isDisabled: false,
                    onClick: () => {
                      console.log('Clicked Foo')
                    },
                  },
                ],
              ],
            },
          },
        ],
      },
    ]"
    :statusBarInfo="['Something goes here...', 'Something else here']"
    :isResizable="true"
    @click-close="emit('clickClose')"
  >
    <div class="portfolio-view">
      <h1>Portfolio</h1>
      <h2>About me</h2>
      <section>
        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iure perspiciatis obcaecati,
        reprehenderit non hic consequatur corrupti illum molestiae illo ipsum, tenetur libero facere
        recusandae. Quisquam molestias nulla earum cum at!
      </section>
      <h2>Professional Experience</h2>
      <section>
        Lorem ipsum, dolor sit amet consectetur adipisicing elit. Iure perspiciatis obcaecati,
        reprehenderit non hic consequatur corrupti illum molestiae illo ipsum, tenetur libero facere
        recusandae. Quisquam molestias nulla earum cum at!
      </section>
      <h2>Personal Projects</h2>
      <section>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Pariatur modi reiciendis voluptatum
        ipsum ab nostrum voluptatem, beatae eligendi deserunt provident rerum sunt. Corrupti quos
        praesentium corporis ducimus quae autem vitae!
      </section>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativeSize } from '@/core/models/relative_size'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
</script>

<style scoped lang="scss">
.portfolio-view {
  width: 100%;
  height: 100%;
  padding: 5px;
}
</style>
