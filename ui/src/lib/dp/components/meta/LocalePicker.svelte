<script lang="ts">
  import { goto } from "$app/navigation";
  import { page } from "$app/state";
  import { appLocale } from "$lib/dp/locale.svelte.ts";

  async function reload() {
    await goto(page.url, {
      noScroll: true,
      keepFocus: true,
      invalidateAll: true,
    });
    const htmlElement = document.querySelector("html");

    if (htmlElement) {
      htmlElement.lang = appLocale.current;
    }
  }

  const uid = $props.id();
</script>

<label for={uid} class="sr-only">App Language: </label>
<select
  id={uid}
  bind:value={appLocale.current}
  class="select"
  onchange={reload}
>
  <option value="en" aria-label="English">🇺🇸</option>
  <option value="de" aria-label="German">🇩🇪</option>
  <option value="fr" aria-label="French">🇫🇷</option>
  <option value="es" aria-label="Spanish">🇪🇸</option>
  <option value="pl" aria-label="Polish">🇵🇱</option>
</select>

<style>
  select {
    font-size: 1.5rem;
  }

  option {
    font-size: 2rem;
  }
</style>
