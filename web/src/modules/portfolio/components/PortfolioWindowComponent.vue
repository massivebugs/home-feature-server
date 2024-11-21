<template>
  <WindowComponent
    :size="new RelativeSize(70, 80)"
    :title="t('portfolio.title')"
    :controls="{
      minimize: true,
      maximize: true,
      close: true,
    }"
    title-bar-icon="/images/portfolio_icon_small.png"
    :resizable="true"
    @click-close="emit('clickClose')"
    @window-resize="onWindowResize"
    v-slot="windowProps"
  >
    <div
      ref="container"
      :class="['portfolio__container', `portfolio__container-${windowProps.windowSize}`]"
      @scrollend="onContainerScrollEnd"
    >
      <section id="portfolio__intro" class="portfolio__section">
        <p class="portfolio__intro__title sandman-text">Hi, I'm Da-Hyun.</p>
        <p class="portfolio__intro__description sandman-text">I'm a full-stack developer.</p>
        <p class="portfolio__intro__links">
          <ButtonComponent type="primary" class="sandman-text">View Resume</ButtonComponent>
          <a href="#portfolio__contact">
            <ButtonComponent href="#portfolio__contact" class="sandman-text">
              Contact Me
            </ButtonComponent>
          </a>
        </p>
        <p>
          <button @click="sandman?.init()">Init</button>
          <button @click="sandman?.drop(0.1)">Drop</button>
          <button @click="sandman?.collect(1)">Collect</button>
          <button
            @click="
              async () => {
                await sandman?.spiralFill(0, 0, 0)
                await onClickRestoreText()
              }
            "
          >
            Surround
          </button>
          <button @click="sandman?.drop(0.1).then((s) => s.collect(1))">Connect</button>
          <button @click="onClickRestoreText">SetText</button>
          <button @click="sandman?.reset()">Reset</button>
        </p>
      </section>
      <section id="portfolio__experience" class="portfolio__section">
        <h1 class="sandman-text">Experience</h1>
        <div class="portfolio__experience-list">
          <div class="portfolio__experience-item">
            <p class="portfolio__experience-name sandman-text">LEAN BODY Inc.</p>
            <p class="portfolio__experience-title sandman-text">
              Web Engineer | Nov 2023 - Oct 2024 (Full-time)
            </p>
            <ul>
              <li class="sandman-text">
                Full-stack engineer for one of Japan’s largest online fitness platforms, developing
                new features and fixes for hundreds of thousands of users.
              </li>
              <li class="sandman-text">
                Collaborated and brainstormed closely with cross-functional teams (product, design,
                analytics, support, engineering) to enhance user experience, contributing to 7+
                major feature releases in under a year with minimal bugs.
              </li>
              <li class="sandman-text">
                Implemented test code generation, optimized test execution, and managed a major
                database upgrade, reducing test writing time by 3 minutes per test and cutting
                integration test time fivefold.
              </li>
            </ul>
          </div>

          <div class="portfolio__experience-item">
            <p class="portfolio__experience-name sandman-text">TERADOGA Co., Ltd</p>
            <p class="portfolio__experience-title sandman-text">
              Software Engineer | Jul 2020 - Oct 2023 (Full-time) | Nov 2023 - Oct 2024 (Contract)
            </p>
            <ul>
              <li class="sandman-text">
                Led the development of TERADOGA, the company’s flagship product, adjusting goals to
                align with a new business model that secured long-term agreements with three new
                business clients within the first year of development.
              </li>
              <li class="sandman-text">
                Developed and deployed over five full-stack Laravel and Vue.js applications from the
                ground up for various clients, each with unique business requirements, over a span
                of 2.5 years.
              </li>
              <li class="sandman-text">
                Acted as a bridge software engineer, effectively collaborating across three teams
                from separate companies in both English and Japanese to successfully meet feature,
                schedule, and deployment requirements.
              </li>
            </ul>
          </div>
        </div>
      </section>
      <section id="portfolio__skills" class="portfolio__section">
        <h1 class="sandman-text">Technical Skills</h1>
        <div class="portfolio__skill-list">
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Programming Languages</div>
            <p class="sandman-text">Go, PHP, Javascript, Typescript</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Back-end Frameworks</div>
            <p class="sandman-text">Echo(PHP), Laravel(PHP),</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Front-end Frameworks</div>
            <p class="sandman-text">Vue.js, React.js, Flutter, Bootstrap, Tailwind CSS</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Cloud Platforms</div>
            <p class="sandman-text">Amazon Web Services (AWS), Google Cloud Platform (GCP)</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Databases</div>
            <p class="sandman-text">MySQL, MariaDB</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Containerization</div>
            <p class="sandman-text">Docker</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type sandman-text">Other</div>
            <p class="sandman-text">
              Attention to detail, strong work ethic, flexibility and adaptability
            </p>
          </div>
        </div>
      </section>
      <section id="portfolio__projects" class="portfolio__section">
        <h1 class="sandman-text">Projects</h1>
        <div class="portfolio__project-list">
          <div class="portfolio__project-item">
            <div class="sandman-text">teradoga.jp</div>
            <div class="sandman-text">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam quasi saepe, a minima
              velit possimus necessitatibus, at facilis molestias dignissimos eaque consectetur
              optio ipsam ab ipsum molestiae dicta inventore minus.
            </div>
          </div>
          <div class="portfolio__project-item">
            <div class="sandman-text">lean-body.jp</div>
            <div class="sandman-text">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam quasi saepe, a minima
              velit possimus necessitatibus, at facilis molestias dignissimos eaque consectetur
              optio ipsam ab ipsum molestiae dicta inventore minus.
            </div>
          </div>
          <div class="portfolio__project-item">
            <div class="sandman-text">Home feature server (Portfolio site)</div>
            <div class="sandman-text">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam quasi saepe, a minima
              velit possimus necessitatibus, at facilis molestias dignissimos eaque consectetur
              optio ipsam ab ipsum molestiae dicta inventore minus.
            </div>
          </div>
          <div class="portfolio__project-item">
            <div class="sandman-text">TELEBYTE</div>
            <div class="sandman-text">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quam quasi saepe, a minima
              velit possimus necessitatibus, at facilis molestias dignissimos eaque consectetur
              optio ipsam ab ipsum molestiae dicta inventore minus.
            </div>
          </div>
        </div>
      </section>
      <section id="portfolio__contact" class="portfolio__section">
        <h1 class="sandman-text">Contact Me</h1>
      </section>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ButtonComponent from '@/core/components/ButtonComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativeSize } from '@/core/models/relativeSize'
import { sleep } from '@/core/utils/time'
import { SANDMAN_TEXT_CLASS, Sandman } from '@/modules/sandman/sandman'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const props = defineProps<{ initSandman?: boolean }>()

const { t } = useI18n()
const container = ref<HTMLElement>()
let sandman: Sandman | null = null

const onWindowResize = () => {
  if (!props.initSandman) {
    return
  }

  sandman?.onContainerResize()
  sandman?.drop(0.1)
}

const onContainerScrollEnd = () => {
  if (!props.initSandman) {
    return
  }

  sandman?.drop(0.1)
}

const onClickRestoreText = async () => {
  const texts = document
    .getElementById('portfolio__intro')
    ?.getElementsByClassName(SANDMAN_TEXT_CLASS)

  if (!texts) {
    return
  }

  for (const text of texts) {
    await sandman?.restoreText(text as HTMLElement, 0.05, 0.3, 0.1)
    await sleep(200)
  }
}

onMounted(async () => {
  if (container.value) {
    sandman = new Sandman(container.value)

    await sleep(100)
    sandman?.init()
    sandman?.drop(0.1)
    await sleep(10000)
    container.value.scrollTop = 0
    sandman?.collect()
  }
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.portfolio__container {
  width: 100%;
  height: 100%;
  max-width: 100%;
  max-height: 100%;
  overflow-x: hidden;
  overflow-y: scroll;
  scroll-snap-type: y mandatory;
  background-color: colors.$low-opacity-white;
}

.portfolio__section {
  scroll-snap-align: start;

  min-height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 1em;
}

#portfolio__intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.portfolio__intro__title {
  margin-bottom: 0;
  font-size: 1.3em;
}

.portfolio__intro__links {
  display: flex;
  gap: 0.7em;

  > a {
    text-decoration: none;
    button {
      width: 100%;
    }
  }

  button {
    padding: 0.4em 1em;
  }

  > *,
  > {
    width: 50%;
    white-space: nowrap;
  }
}

.portfolio__container-xl > .portfolio__section {
  padding: 0 23%;
}

.portfolio__container-lg > .portfolio__section {
  padding: 0 15%;
}

.portfolio__container-md > .portfolio__section {
  padding: 0 5%;
}

.portfolio__experience-list {
  display: flex;
  flex-direction: column;
  gap: 2em;
}

.portfolio__experience-name {
  font-weight: 700;
  margin-bottom: 0.3em;
}

.portfolio__experience-title {
  margin-top: 0;
  font-style: italic;
}

.portfolio__experience-item > ul > li {
  margin-bottom: 0.7em;
}

.portfolio__skill-list {
  display: flex;
  flex-direction: column;
}

.portfolio__project-list {
  width: 100%;
  flex: 1;
  display: flex;
  flex-wrap: wrap;
}

.portfolio__project-item {
  min-width: 100%;
  height: 25%;
  flex: 1;
  padding: 0.7em;
}

.portfolio__container-lg,
.portfolio__container-xl {
  .portfolio__project-item {
    min-width: 50%;
    height: 50%;
  }
}
</style>
