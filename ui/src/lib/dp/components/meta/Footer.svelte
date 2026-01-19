<script lang="ts">
  import {env} from "$env/dynamic/public";
  import type {Snippet} from "svelte";

  interface Props {
    staticLinks?: Snippet;
  }

  let {staticLinks}: Props = $props();
</script>

<section class="flex justify-center items-center gap-8 mt-5">
  {@render staticLinks?.()}
</section>

<nav class="list-nav pt-1 pb-3 px-6">

  <ul class="flex mb-3 flex-col md:flex-row items-center justify-center md:gap-4">
    <li><a href="/terms">Terms & Conditions</a></li>
    <li><a href="/legalnotice">Impressum</a></li>
    <li><a href="/privacy">Datenschutzerklärung</a></li>
  </ul>

</nav>

<div class="flex justify-end items-center gap-4 pb-3 mx-4">
  <!-- @wc-ignore -->
  {#if env.PUBLIC_APPLICATION_CONTEXT !== "Production"}
    <div class="preset-tonal-warning p-1.5 rounded-container border border-warning-500 text-sm">
      {env.PUBLIC_APPLICATION_CONTEXT}
    </div>
  {/if}

  <span class="flex justify-end items-center gap-2">
    Built with <span class="love">❤️</span> in <span class="flag">🇪🇺</span>
  </span>
</div>

<style>
  .list-nav {
    a:hover {
      text-decoration: underline;
    }
  }

  .love {
    font-size: var(--text-lg);
  }

  .flag {
    font-size: var(--text-3xl);
  }
</style>
