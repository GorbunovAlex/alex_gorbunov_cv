import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { IArticle } from '@/helpers/types';

export const useArticlesStore = defineStore('articles', () => {
  const articles = ref<IArticle[]>([]);

  async function fetchArticles() {
  }

  return { articles, fetchArticles };
})