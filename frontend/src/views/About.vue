<script setup lang="ts">
</script>

<template>
  <div class="w-full flex justify-center">
    <div class="about flex column items-center justify-center">
      <q-list bordered class="w-full">
        <q-expansion-item label="Bio" expand-icon-class="text-white" class="text-center main-color-text">
          <q-card class="bg-transparent">
            <q-card-section>
              <h5 class="text-white">/* Frontend */</h5>
              <p>
                I have strong knowledge and practical experience in <span class="text-white">JavaScript</span>, expanding it with a solid understanding of <span class="text-white">TypeScript</span>.
              </p>

              <p>
                <span class="text-white">Vue.js</span> is my default choice among JS frameworks. I started my journey with Vue2 and Nuxt2 and have since developed extensive experience with <span class="text-white">Vue3</span>, the Composition API, and <span class="text-white">Nuxt3</span>, leveraging them in full-time roles and freelance projects.
              </p>

              <p>
                I used <span class="text-white">Flutter</span> to build applications for Android, iOS, Windows, and macOS of various scales and purposes, including LLM integration, video streaming, and music generation.
              </p>
            </q-card-section>
            <q-card-section>
              <h5 class="text-white">/* Backend */</h5>
              <p>
                I’m experienced in creating RESTful APIs using <span class="text-white">Go</span> and setting up databases with <span>PostgreSQL</span>. This includes working on payments, authorization & authentication, and LLM integration services.
              </p>

              <p>
                Additionally, I have used <span class="text-white">Python</span> and <span class="text-white">C++</span> to develop general business logic in full-time roles.
              </p>

            </q-card-section>
            <q-card-section>
              <h5 class="text-white">/* Other */</h5>
              <p>
                It’s worth mentioning that I have experience <span class="text-white">Dockerizing apps</span>, setting up <span class="text-white">Nginx</span>, and configuring server environments on <span class="text-white">Linux</span> systems.
              </p>

              <p>
                One of my key achievements is having extensive experience in the hiring process as a <span class="text-white">technical interviewer</span>. Over the course of nearly three years, I have interviewed a large number of candidates.
              </p>

              <p>
                I speak English at an advanced level (C1). I enjoy communicating and interacting with people, and I’m easy-going and open to establishing new connections.
              </p>

            </q-card-section>
          </q-card>
        </q-expansion-item>
      </q-list>
    </div>
  </div>
</template>

<style scoped lang="scss">
.about {
  width: 60%;

  @media (max-width: 430px) {
    width: 100%;
  }
}
</style>
