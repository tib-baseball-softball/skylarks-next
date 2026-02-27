<script lang="ts">
  import {env} from "$env/dynamic/public";
  import type {Snippet} from "svelte";

  interface Props {
    staticLinks?: Snippet;
  }

  let {staticLinks}: Props = $props();
</script>

<section class="links-wrapper">
  {@render staticLinks?.()}
</section>

<nav class="list-nav">

  <ul class="nav-list">
    <li><a href="/terms">Terms & Conditions</a></li>
    <li><a href="/legalnotice">Impressum</a></li>
    <li><a href="/privacy">Datenschutzerklärung</a></li>
  </ul>

</nav>

<div class="footer-bottom">
  <!-- @wc-ignore -->
  {#if env.PUBLIC_APPLICATION_CONTEXT !== "Production"}
    <div class="context-badge preset-tonal-warning rounded-container">
      {env.PUBLIC_APPLICATION_CONTEXT}
    </div>
  {/if}

  <span class="attribution">
    Built with <span class="love">❤️</span> in <span class="flag">🇪🇺</span>
  </span>
</div>

<style>
  .links-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: calc(var(--spacing) * 8);
    margin-top: calc(var(--spacing) * 5);
  }

  .list-nav {
    padding-top: var(--spacing);
    padding-bottom: calc(var(--spacing) * 3);
    padding-inline: calc(var(--spacing) * 6);

    a:hover {
      text-decoration: underline;
    }
  }

  .nav-list {
    display: flex;
    margin-bottom: calc(var(--spacing) * 3);
    flex-direction: column;
    align-items: center;
    justify-content: center;

    @media (min-width: 64rem) {
      flex-direction: row;
      gap: calc(var(--spacing) * 4);
    }
  }

  .footer-bottom {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: calc(var(--spacing) * 4);
    padding-bottom: calc(var(--spacing) * 3);
    margin-inline: calc(var(--spacing) * 4);
  }

  .context-badge {
    padding: calc(var(--spacing) * 1.5);
    border: 1px solid var(--color-warning-500);
    font-size: var(--text-sm);
  }

  .attribution {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: calc(var(--spacing) * 2);
  }

  .love {
    font-size: var(--text-lg);
  }

  .flag {
    font-size: var(--text-3xl);
  }
</style>
