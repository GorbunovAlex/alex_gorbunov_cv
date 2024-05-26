<script setup lang="ts">
import Langs from '@/assets/icons/Langs.vue';
</script>

<template>
  <div class="main relative">
    <h2 class="">Namaste</h2>
    <p>Welcome to my personal space where I share my development thoughts and showcase my work.</p>
    <Langs class="main__langs" />
    <p>
      I'm <span class="text-primary">Alex Gorbunov </span>, a software developer with 5 years of
      experience in <span class="text-primary">JavaScript and TypeScript</span>, focusing on the Vue
      ecosystem.
    </p>
    <p>
      I have worked with Vue2, <span class="text-primary">Vue3, Nuxt 3</span>, Webpack, Vite,
      Vitest, Cypress, SASS, and Stylus. I also develop REST APIs using
      <span class="text-primary">Go</span>, Beego, <span class="text-primary">Python</span>, FastAPI
      and PostgreSQL, and have experience with Nginx, Docker, and Linux.
    </p>
    <p>
      Fluent in English (C1), I enjoy communication, learning, and attending tech conferences. In my
      spare time, I enjoy jogging, hiking, learning languages, and spending time with my family.
    </p>
  </div>
</template>

<style scoped lang="scss">
.main {
  display: grid;
  max-width: 60%;

  &__langs {
    position: absolute;
    bottom: 0;
    right: 0;
  }
}
</style>
