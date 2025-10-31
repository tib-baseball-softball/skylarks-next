<script lang="ts">
  import {preferences} from "$lib/globals.svelte.ts";
  import {goto} from "$app/navigation";
  import {page} from "$app/state";

  async function reload() {
    await goto(page.url, {
      noScroll: true,
      keepFocus: true,
      invalidateAll: true,
    });
    const htmlElement = document.querySelector("html");

    if (htmlElement) {
      htmlElement.lang = preferences.current.locale;
    }
  }
</script>

<select bind:value={preferences.current.locale} class="select" onchange={reload}>
  <option value="en">🇺🇸</option>
  <option value="de">🇩🇪</option>
  <option value="fr">🇫🇷</option>
  <option value="es">🇪🇸</option>
  <option value="pl">🇵🇱</option>
  <option value="ru">🇷🇺</option>
</select>


<style>
    select {
        font-size: 1.5rem;
    }

    option {
        font-size: 2rem;
    }
</style>