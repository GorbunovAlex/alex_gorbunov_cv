<script setup lang="ts">
import { nextTick, onMounted, ref } from 'vue';
import dayjs from 'dayjs';

import useGsapAnim from '@/composables/useGsapAnim';
import { useRouter } from 'vue-router';

const { createTypewriter } = useGsapAnim();

const cursor = ref(false);
const prog = ref(false);

const router = useRouter();

onMounted(async () => {
  createTypewriter('anim-typewriter', 'gcc -o init_alex_dev init_alex_dev.c').play();
  const promise = new Promise((resolve) => {
    setTimeout(() => {
      cursor.value = true;
      nextTick(() => createTypewriter('anim-typewriter_second', './init_alex_dev').play());
      setTimeout(() => resolve(null), 2500);
    }, 2000);
  });
  await promise.then(() => {
    prog.value = true;
    const innerPromise = new Promise((resolve) => {
      nextTick(() =>
        createTypewriter(
          'loading-progress',
          '##################################################'
        ).play()
      );
      setTimeout(() => resolve(null), 2000);
    });
    innerPromise.then(() => {
      router.push({ name: 'Main' });
    });
  });
});
</script>

<template>
  <div class="w-full">
    <div class="w-full">
      <div>Last login: {{ dayjs().format('ddd MMM DD HH:mm:ss') }} on console</div>
      <div class="flex gap-4">
        alex_gorbunov@somewhere_in_net ~ %
        <div class="anim-typewriter"></div>
        <span v-if="!cursor" class="cursor">|</span>
      </div>
      <div v-if="cursor" class="flex gap-4">
        alex_gorbunov@somewhere_in_net ~ %
        <div class="anim-typewriter_second"></div>
        <span v-if="cursor" class="cursor">|</span>
      </div>
      <div class="column items-center justify-center" v-if="prog">
        <pre class="landing__pre">
          <div>**************************************************</div>
          <div>*                                                *</div>
          <div>*      Welcome to Alex Gorbunov's dev site!      *</div>
          <div>*                                                *</div>
          <div>**************************************************</div>
          <div>*                                                *</div>
          <div>*         Stay awesome and keep coding!          *</div>
          <div>*                                                *</div>
          <div>**************************************************</div>
          <div>*                                                *</div>
          <div>*                    Loading...                  *</div>
          <div class="loading-progress"></div>
        </pre>
      </div>
    </div>
  </div>
</template>

<style scoped>
.anim-typewriter {
  -webkit-box-sizing: content-box;
  box-sizing: content-box;
  white-space: nowrap;
  overflow: hidden;
}

.landing__pre {
  line-height: 10px;
}
</style>
