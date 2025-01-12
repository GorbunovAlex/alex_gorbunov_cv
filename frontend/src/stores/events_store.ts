import { ref } from 'vue';
import { defineStore } from 'pinia';
import type { IEvent } from '@/helpers/types';

export const useEventsStore = defineStore('events', () => {
  const events = ref<IEvent[]>([]);

  async function fetchEvents() {
  }

  return { events, fetchEvents };
})