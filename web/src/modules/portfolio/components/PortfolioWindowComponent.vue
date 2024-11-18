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
    v-slot="windowProps"
  >
    <div
      ref="container"
      :class="['portfolio__container', `portfolio__container-${windowProps.windowSize}`]"
    >
      <section ref="sectionIntro" id="portfolio__intro" class="portfolio__section">
        <p class="portfolio__intro__title">Hi, I'm Da-Hyun.</p>
        <p class="portfolio__intro__description">I'm a full-stack developer.</p>
        <p class="portfolio__intro__links">
          <ButtonComponent> Download Resume </ButtonComponent>
          <a href="#portfolio__contact">
            <ButtonComponent> Contact Me </ButtonComponent>
          </a>
        </p>
        <p>or</p>
        <p>
          <ButtonComponent> Protect this Portfolio </ButtonComponent>
        </p>
      </section>
      <section ref="sectionExperience" id="portfolio__experience" class="portfolio__section">
        <h1>Experience</h1>
        <div class="portfolio__company">
          <p class="portfolio__company-name">LEAN BODY Inc.</p>
          <p class="portfolio__company-title">Web Engineer | Nov 2023 - Oct 2024</p>
          <ul>
            <li>
              Full-stack engineer for one of Japan’s largest online fitness platforms, developing
              new features and fixes for hundreds of thousands of users.
            </li>
            <li>
              Collaborated and brainstormed closely with cross-functional teams (product, design,
              analytics, support, engineering) to enhance user experience, contributing to 7+ major
              feature releases in under a year with minimal bugs.
            </li>
            <li>
              Implemented test code generation, optimized test execution, and managed a major
              database upgrade, reducing test writing time by 3 minutes per test and cutting
              integration test time fivefold.
            </li>
          </ul>
        </div>

        <div class="portfolio__company">
          <p class="portfolio__company-name">TERADOGA Co., Ltd</p>
          <p class="portfolio__company-title">Software Engineer | Jul 2020 - Oct 2024</p>
          <ul>
            <li>
              Led the development of TERADOGA, the company’s flagship product, adjusting goals to
              align with a new business model that secured long-term agreements with three new
              business clients within the first year of development.
            </li>
            <li>
              Developed and deployed over five full-stack Laravel and Vue.js applications from the
              ground up for various clients, each with unique business requirements, over a span of
              2.5 years.
            </li>
            <li>
              Acted as a bridge software engineer, effectively collaborating across three teams from
              separate companies in both English and Japanese to successfully meet feature,
              schedule, and deployment requirements.
            </li>
          </ul>
        </div>
      </section>
      <section ref="sectionSkills" id="portfolio__skills" class="portfolio__section">
        <h1>Technical Skills</h1>
        <div class="portfolio__skill-list">
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Programming Languages</div>
            <p>Go, PHP, Javascript, Typescript</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Back-end Frameworks</div>
            <p>Echo(PHP), Laravel(PHP),</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Front-end Frameworks</div>
            <p>Vue.js, React.js, Flutter, Bootstrap, Tailwind CSS</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Cloud Platforms</div>
            <p>Amazon Web Services (AWS), Google Cloud Platform (GCP)</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Databases</div>
            <p>MySQL, MariaDB</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Containerization</div>
            <p>Docker</p>
          </div>
          <div class="portfolio__skill">
            <div class="portfolio__skill-type">Other</div>
            <p>Attention to detail, strong work ethic, flexibility and adaptability</p>
          </div>
        </div>
      </section>
      <section ref="sectionProjects" id="portfolio__projects" class="portfolio__section">
        <h1>Projects</h1>
        <div class="portfolio__project-list">
          <div class="portfolio__project-item">
            <div>teradoga.jp</div>
          </div>
          <div class="portfolio__project-item">
            <div>lean-body.jp</div>
          </div>
          <div class="portfolio__project-item">
            <div>Home feature server (Portfolio site)</div>
          </div>
          <div class="portfolio__project-item">
            <div>TELEBYTE</div>
          </div>
        </div>
      </section>
      <section ref="sectionContact" id="portfolio__contact" class="portfolio__section">
        Contact Me
      </section>
    </div>
  </WindowComponent>
</template>

<script setup lang="ts">
import { scroll } from 'motion'
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import ButtonComponent from '@/core/components/ButtonComponent.vue'
import WindowComponent from '@/core/components/WindowComponent.vue'
import { RelativeSize } from '@/core/models/relativeSize'

const emit = defineEmits<{
  (e: 'clickClose'): void
}>()

const { t } = useI18n()
const container = ref<HTMLElement>()
const sectionIntro = ref<HTMLElement>()
const sectionExperience = ref<HTMLElement>()
const sectionSkills = ref<HTMLElement>()
const sectionProjects = ref<HTMLElement>()
const sectionContact = ref<HTMLElement>()

onMounted(() => {
  console.log(container.value?.textContent)
  scroll(
    (progress: number) => {
      console.log(progress)
    },
    { container: container.value },
  )

  // document.querySelectorAll('section').forEach((section) => {
  //   const header = section.querySelector('h2')
  //   scroll(animat(header, { y: [-400, 400] }, { ease: 'linear' }), {
  //     target: header,
  //   })
  // })
})
</script>

<style scoped lang="scss">
@use '@/assets/colors';

.portfolio__container {
  width: 100%;
  height: 100%;
  max-width: 100%;
  max-height: 100%;
  overflow-y: scroll;
  padding: 1em;
  scroll-snap-type: y mandatory;
}

.portfolio__section {
  scroll-snap-align: start;

  height: 100%;
  display: flex;
  flex-direction: column;

  > h1 {
    margin-top: 0;
  }
}

#portfolio__intro {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.portfolio__intro__title {
  margin-bottom: 0;
}

.portfolio__intro__links {
  display: flex;
  gap: 0.7em;
  a {
    text-decoration: none;
  }
}

.portfolio__container-xl #portfolio__experience {
  padding: 0 30%;
}

.portfolio__container-lg #portfolio__experience {
  padding: 0 15%;
}

.portfolio__container-md #portfolio__experience {
  padding: 0 5%;
}

.portfolio__company {
  margin-bottom: 2em;
}

.portfolio__company-name {
  font-weight: 700;
  margin-bottom: 0.3em;
}

.portfolio__company-title {
  margin-top: 0;
  font-style: italic;
}

.portfolio__company > ul > li {
  margin-bottom: 0.7em;
}

.portfolio__container-xl #portfolio__skills {
  padding: 0 30%;
}

.portfolio__container-lg #portfolio__skills {
  padding: 0 15%;
}

.portfolio__container-md #portfolio__skills {
  padding: 0 5%;
}

.portfolio__skill-list {
  flex: 1;
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
  flex: 1;
  min-width: 50%;
  height: 50%;
}
</style>
