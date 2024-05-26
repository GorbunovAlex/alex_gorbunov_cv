import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Landing',
      component: () => import('@/views/Landing.vue')
    },
    {
      path: '/main',
      component: () => import('@/layouts/Default.vue'),
      children: [
        {
          path: '',
          name: 'Main',
          component: () => import('@/views/Main.vue')
        },
        {
          path: '/main/about',
          name: 'About',
          component: () => import('@/views/About.vue')
        },
        {
          path: '/main/blog',
          name: 'Blog',
          component: () => import('@/views/Blog.vue')
        }
      ]
    }
  ]
});

export default router;
