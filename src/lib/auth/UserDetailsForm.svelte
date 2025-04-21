<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {closeModal} from "$lib/functions/closeModal.ts";

  const authRecord = authSettings.record as CustomAuthModel;

  let files: FileList | null = $state(null);

  const form = $state({
    id: authRecord.id,
    firstName: authRecord.first_name,
    lastName: authRecord.last_name,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const formData = new FormData();

    formData.append("id", form.id);

    if (form.firstName) {
      formData.append("first_name", form.firstName);
    }

    if (form.lastName) {
      formData.append("last_name", form.lastName);
    }

    if (files) {
      for (let file of files) {
        formData.append("avatar", file);
      }
    }

    client.collection("users").update(form.id, formData, {expand: "club"}).then(closeModal);
  }
</script>

<p class="mt-2 font-light">
  This data is only used for internal purposes and never shown anywhere publicly.
  (Your team and club administrators have access).
</p>

<form onsubmit={submitForm} class="mt-4 space-y-3">
  <input
          name="id"
          autocomplete="off"
          class="input"
          type="hidden"
          readonly
          bind:value={form.id}
  />

  <label class="label">
    First Name(s)
    <input name="firstname" class="input" bind:value={form.firstName} autocomplete="given-name"/>
  </label>

  <label class="label">
    Last Name
    <input name="lastname" class="input" bind:value={form.lastName} autocomplete="family-name"/>
  </label>

  <label class="label">
    Profile Image
    <input
            class="input"
            name="avatar"
            type="file"
            accept="image/*"
            bind:files
    />
  </label>

  {#if files}
    <h2>Selected file:</h2>
    {#each Array.from(files) as file}
      <p>{file.name} ({file.size} bytes)</p>
    {/each}
  {/if}

  <p class="font-light">
    This is your app profile picture. It will not be used for your player profile image.
  </p>

  <div class="flex justify-end gap-3 mt-3">
    <button type="submit" class="mt-2 btn variant-ghost-primary">Confirm</button>
  </div>
</form>
