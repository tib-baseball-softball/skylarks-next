<script>
    import Svelte from "../icons/Svelte.svelte";
    import {Heart} from "lucide-svelte";
    import {env} from "$env/dynamic/public";
</script>

<section class="flex justify-center items-center gap-8 mt-5">
  <a aria-label="to Skylarks Facebook page" href="https://www.facebook.com/TiBBaseball/" rel="noreferrer"
     target="_blank">
    <img alt="Facebook brand logo" src="/Facebook_Logo_Primary.svg" width="50">
  </a>

  <a aria-label="to Skylarks Instagram profile" href="https://www.instagram.com/berlinskylarks/" rel="noreferrer"
     target="_blank">
    <img alt="Instagram brand logo" class="m-1 max-w-[40px]" src="/Instagram_Glyph_Gradient.png">
  </a>

  <a aria-label="to Turngemeinde in Berlin main website" href="https://tib1848ev.de/" target="_blank">
    <img alt="TiB Logo" class="min-w-8 max-w-14" src="/tib_logo.svg" width="38">
  </a>

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

  <!-- @wc-include -->
  <span class="flex justify-end items-center gap-2">
    <span>Built with</span>
    <Heart class="text-primary-500"/>
    <span>and</span>
    <a href="https://svelte.dev/" target="_blank" title="Svelte JavaScript Framework website">
      <Svelte classes="w-6" fillColor={"#ff3e00"}/>
    </a>
  </span>
</div>

<style>
    .list-nav {
        a:hover {
            text-decoration: underline;
        }
    }
</style>