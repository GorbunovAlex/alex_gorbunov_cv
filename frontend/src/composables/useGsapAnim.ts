import { gsap } from 'gsap';
import { TextPlugin } from 'gsap/TextPlugin';

gsap.registerPlugin(TextPlugin);

export default function useGsapAnim() {
  function createTypewriter(selector: string, text: string) {
    const tl = gsap.timeline({
      paused: true
    });
    tl.to(`.${selector}`, {
      text: text,
      duration: 2,
      ease: 'power1.in'
    });

    return tl;
  }

  return {
    createTypewriter
  };
}
