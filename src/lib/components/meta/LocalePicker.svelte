<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";
  import {goto} from "$app/navigation";
  import {page} from "$app/state";
  import type {AppLocales} from "$lib/types/AppPreferences.ts";

  function setLocaleToPreferences(locale: AppLocales) {
    preferences.current.locale = locale;
    goto(page.url.toString(), {
      noScroll: true,
      keepFocus: true,
      invalidateAll: true,
    });
  }
</script>

{#if preferences.current.locale === "de"}
  <button onclick={() => setLocaleToPreferences("en")}>
    <span class="preset-outline-surface flag-icon rounded-base">
    🇩🇪
    </span>
  </button>
{:else }
  <button class="" onclick={() => setLocaleToPreferences("de")}>
    <span class="preset-outline-surface flag-icon rounded-base">
    🇬🇧
    </span>
  </button>
{/if}

<style>
    .flag-icon {
        font-size: 2rem;
        padding: 0.5rem;
    }
</style>