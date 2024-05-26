<script setup lang="ts">
import { useRoute } from 'vue-router';

import Linkedin from '@/assets/icons/Linkedin.vue';
import Telegram from '@/assets/icons/Telegram.vue';
import Github from '@/assets/icons/Github.vue';

const route = useRoute();

function navigateToSite(link: string) {
  window.open(link, '_blank');
}
</script>

<template>
  <q-layout view="hHh lpR fFf">
    <q-header class="header">
      <q-toolbar class="header__toolbar row items-center justify-between">
        <div class="row items-center gap-8">
          <q-btn flat label="Main" to="/main" :class="{ header__active: route.path === '/main' }" />
          <q-btn
            flat
            label="About"
            to="/main/about"
            :class="{ header__active: route.path === '/main/about' }"
          />
          <q-btn
            flat
            label="Blog"
            to="/main/blog"
            :class="{ header__active: route.path === '/main/blog' }"
          />
        </div>
        <div class="row item-center gap-8">
          <Linkedin
            class="icon"
            @click="navigateToSite('https://www.linkedin.com/in/gorbunovalex/')"
          />
          <Github class="icon" @click="navigateToSite('https://github.com/GorbunovAlex')" />
          <Telegram class="icon" />
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container class="q-pa-md">
      <router-view />
    </q-page-container>

    <q-footer class="bg-transparent text-primary">
      <q-toolbar>
        <q-toolbar-title>
          <p>&copy; 2024 Alex Gorbunov. All rights coded.</p>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<style lang="scss">
body {
  color: white;
}

.header {
  display: flex;
  justify-content: flex-end;
  background: transparent;
  padding: 10px;

  &__toolbar {
    width: 400px;
    background: $primary;
    border-radius: 5px;
  }

  &__active {
    text-shadow:
      0 0 5px #fff,
      0 0 10px #fff,
      0 0 15px #fff,
      0 0 20px #0fa,
      0 0 30px #0fa,
      0 0 40px #0fa,
      0 0 55px #0fa,
      0 0 75px #0fa;
  }
}
</style>
