<script lang="ts">
  import {goto} from "$app/navigation";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {authSettings, client} from "../client.svelte.js";

  async function logout() {
    closeModal();
    authSettings.record = null;
    client.authStore.clear();
    await goto("/login");
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

<article class="root">
  <div class="row header-row">
    <p>Logged in as:</p>

    <a class="badge preset-filled-primary-500" href="/account" onclick={closeModal}>
      <samp>{`${authRecord?.first_name ?? ""} ${authRecord?.last_name ?? ""}`}</samp>
    </a>
  </div>

  <div class="row clubs-row">
    <p>Club:</p>

    <div class="chips">
      {#each authRecord?.expand?.club as club}
        <a href="/account/clubs/{club.id}" class="badge preset-filled-tertiary-500" onclick={closeModal}>
          {club.name}
        </a>
      {/each}
    </div>
  </div>

  <button class="btn preset-tonal-primary logout-btn" onclick={logout}>Log out</button>
</article>

<style>
  .root {
    padding-block: calc(var(--spacing) * 1);
  }

  .row {
    display: flex;
    justify-content: space-between;
    margin-top: calc(var(--spacing) * 1);
    margin-bottom: calc(var(--spacing) * 5);
  }

  .header-row {
    align-items: center;
  }

  .clubs-row {
    align-items: flex-start;
  }

  .chips {
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 2);
    justify-content: flex-end;
  }

  .logout-btn {
    border: 1px solid var(--color-primary-500);
    margin-bottom: calc(var(--spacing) * 1);
  }
</style>
