import type { AxiosError } from 'axios';
import Toastify from 'toastify-js';

export class ApiError {
  private error;

  constructor(error: AxiosError) {
    this.error = error;
  }

  handle() {
    console.error('Error occurred: ', this.error.message);
    Toastify({
      text: this.error.message,
      duration: 2000,
      style: {
        background: 'rgb(var(--v-theme-error)',
      },
    }).showToast();
  }
}