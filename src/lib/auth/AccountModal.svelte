<script lang="ts">
  import {authSettings, client} from "../pocketbase/index.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";

  const {parent} = $props();

  async function logout() {
    //@ts-ignore
    parent.onClose();
    client.authStore.clear();
  }

  async function closeModal() {
    // @ts-ignore => this is why this is wrapped in another function
    parent.onClose();
  }

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

<div class="w-modal-slim">
  <div class="card p-4">
    <div class="flex mt-1 mb-5 justify-between items-center">
      <p>Logged in as:</p>

      <a
              href="/account"
              class="badge variant-filled-primary"
              onclick={closeModal}
      >
        <samp
        >{`${authRecord?.first_name ?? ""} ${authRecord?.last_name ?? ""}`}</samp
        >
      </a>
    </div>

    <div class="flex mt-1 mb-5 justify-between items-start">
      <p>Club:</p>

      <div class="flex flex-wrap gap-2 justify-end">
        {#each authRecord?.expand?.club as club}
          <a href="/account/clubs/{club.id}" class="badge variant-filled-tertiary" onclick={closeModal}>
            {club.name}
          </a>
        {/each}
      </div>
    </div>

    <button class="btn variant-ghost-primary mb-1" onclick={logout}
    >Log out
    </button
    >
  </div>
</div>
