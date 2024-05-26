<script setup lang="ts">
import { useRoute } from 'vue-router';

import Linkedin from '@/assets/icons/Linkedin.vue';
import Github from '@/assets/icons/Github.vue';

const route = useRoute();
</script>

<template>
  <q-layout view="lHh Lpr lFf">
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
          <Linkedin class="icon" />
          <Github class="icon" />
        </div>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
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
    background: $normal-header;
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
