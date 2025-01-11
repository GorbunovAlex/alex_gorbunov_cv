<script setup lang="ts">
</script>

<template>
  <div class="flex justify-center">
    <div class="main flex column justify-center items-center q-mt-lg">
      <q-avatar size="150px">
        <img src="me.webp" alt="My image">
      </q-avatar>
      <p class="main-color-text q-mt-lg q-mb-none">X_23KSZZS_MJ</p>
      <span class="main__greetings-caption text-caption"> – as one droid said, or greetings to my space!</span>
      <p class="main__alex-text main-color-text text-body1 q-mt-lg">My name is Alex, and I’m a software Jedi striving to reach a CTO position. Most of my magic happens with <span class="text-white">JavaScript</span> and <span class="text-white">TypeScript</span>, but I also wield the Force with other technologies and languages.</p>
    </div>
  </div>
</template>

<style scoped lang="scss">
.main {
  max-width: 60%;

  &__greetings-caption {
    transform: translate(100px, 0);
  }

  &__alex-text {
    width: 50%;
    text-align: center;
  }

}
</style>
