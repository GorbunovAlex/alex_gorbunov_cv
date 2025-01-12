<script setup lang="ts">
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';

import Linkedin from '@/components/icons/Linkedin.vue';
import Github from '@/components/icons/Github.vue';

const route = useRoute();

function navigateToSite(link: string) {
  window.open(link, '_blank');
}
</script>

<template>
  <q-layout view="hHh lpR fFf" class="app-layout">
    <q-header class="header">
      <q-toolbar class="header__toolbar row items-center justify-between">
        <q-btn flat label="Main" to="/main" :class="{ header__active: route.path === '/main' }" />
        <q-btn
          flat
          label="About"
          to="/main/about"
          :class="{ header__active: route.path === '/main/about' }"
        />
        <q-btn
          flat
          disable
          label="Blog"
          to="/main/blog"
          :class="{ header__active: route.path === '/main/blog' }"
        />

      </q-toolbar>
    </q-header>

    <q-page-container class="q-pa-md">
      <router-view class="page" />
    </q-page-container>

    <q-footer class="bg-transparent">
      <q-toolbar class="flex column gap-4">
        <div class="row item-center gap-8">
          <Linkedin
            class="icon"
            @click="navigateToSite('https://www.linkedin.com/in/gorbunovalex/')"
          />
          <Github class="icon" @click="navigateToSite('https://github.com/GorbunovAlex')" />
        </div>
        <p class="main-color-text text-caption">&copy; {{ dayjs().format('YYYY') }} Alex Gorbunov. All rights coded.</p>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<style lang="scss">
.app-layout {
  color: white;
  max-width: 100dvw;
}

.header {
  display: flex;
  justify-content: center;
  background: transparent;

  &__toolbar {
    width: 300px;
    color: $main-text
  }

  &__active {
    text-shadow:
      0 0 5px #fff,
      0 0 10px #fff,
      0 0 15px rgba(48, 209, 19, 0.29),
      0 0 20px rgba(48, 209, 19, 0.39),
      0 0 30px rgba(48, 209, 19, 0.58),
      0 0 40px rgba(48, 209, 19, 0.79),
      0 0 50px #30d113;
  }
}

.page {
  height: calc(100dvh - 120px);
  overflow: auto;
}
</style>
