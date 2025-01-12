<script setup lang="ts">
import { useQuasar } from 'quasar';

import { useArticlesStore } from '@/stores/articles_store';
import { useEventsStore } from '@/stores/events_store';

const articlesStore = useArticlesStore()
const eventsStore = useEventsStore()

const $q = useQuasar()

</script>

<template>
  <div class="flex items-start justify-center">
    <div class="main flex column justify-center items-center q-mt-lg">
      <q-avatar :size="!$q.screen.lt.lg ? '150px' : '100px'">
        <img src="me.webp" alt="My image">
      </q-avatar>
      <h4 class="main-color-text q-mt-lg q-mb-none">X_23KSZZS_MJ</h4>
      <span class="main__greetings-caption text-caption"> – as one droid said, or greetings to my space!</span>
      <p class="main__alex-text main-color-text q-mt-lg">My name is Alex, and I’m a software Jedi striving to reach a CTO position. Most of my magic happens with <span class="text-white">JavaScript</span> and <span class="text-white">TypeScript</span>, but I also wield the Force with other technologies and languages.</p>
      <div>
        <div v-if="eventsStore.events.length !== 0">
          <h4 class="text-center">Latest events</h4>
          <div>

          </div>
        </div>
        <div v-if="articlesStore.articles.length !== 0">
          <h4 class="text-center">Latest posts</h4>
          <div></div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.main {
  width: 60%;

  @media (max-width: 430px) {
    width: 100%;
  }

  &__greetings-caption {
    transform: translate(100px, 0);

    @media (max-width: 430px) {
      transform: translate(0, 0);
    }
  }

  &__alex-text {
    width: 50%;
    text-align: center;

    @media (max-width: 430px) {
      width: 100%;
    }
  }

}
</style>
