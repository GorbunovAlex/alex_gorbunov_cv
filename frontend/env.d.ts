/// <reference types="vite/client" />
declare module 'toastify-js' {
  export default function Toastify(options: {
    text: string;
    duration: number;
    style: Record<string, string>;
  }): { showToast: () => void };
}