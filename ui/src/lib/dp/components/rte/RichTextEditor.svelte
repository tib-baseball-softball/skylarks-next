<script lang="ts">
  import {Carta, MarkdownEditor} from 'carta-md';
  import 'carta-md/default.css';
  import DOMPurify from 'dompurify';
  import {onMount} from "svelte";

  interface Props {
    value: string;
    required?: boolean;
  }

  let {value = $bindable(), required = false}: Props = $props();

  const carta = new Carta({
    sanitizer: DOMPurify.sanitize,
  });

  onMount(() => {
    // this is a hack, for some reason no classes can be added via the component
    const cartaRenderers = document.querySelectorAll('.carta-renderer.markdown-body');
    if (cartaRenderers.length === 0) {
      console.warn('Carta renderer elements not found');
      return;
    }

    cartaRenderers.forEach((renderer) => {
      renderer.classList.add("prose");
    });
  });
</script>

<MarkdownEditor bind:value mode="auto" {carta} textarea={{required: required}}/>
