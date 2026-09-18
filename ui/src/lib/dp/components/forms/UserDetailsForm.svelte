<script lang="ts">
  import { authSettings, client } from "$lib/dp/client.svelte.js";
  import type { CustomAuthModel } from "$lib/dp/types/ExpandedResponse.ts";
  import { Collection } from "$lib/dp/enum/Collection.ts";
  import { ClientResponseError } from "pocketbase";
  import { toastController } from "$lib/dp/service/ToastController.svelte";

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let files: FileList | null = $state(null);

  function formFromProps(record: CustomAuthModel) {
    return {
      id: record.id,
      firstName: record.first_name,
      lastName: record.last_name,
      displayName: record.display_name ?? "",
      emailVisibility: record.emailVisibility,
    };
  }

  let form = $derived.by(() => {
    const formData = $state(formFromProps(authRecord));
    return formData;
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

    if (form.displayName) {
      formData.append("display_name", form.displayName);
    }

    formData.append(
      "emailVisibility",
      form.emailVisibility === true ? "true" : "false",
    );

    if (files) {
      for (const file of files) {
        formData.append("avatar", file);
      }
    }

    try {
      client
        .collection(Collection.Users)
        .update(form.id, formData, { expand: "club" });
    } catch (error) {
      if (error instanceof ClientResponseError) {
        toastController.trigger({
          message: error.message,
          background: "preset-filled-error-500",
        });
      } else {
        toastController.triggerGenericErrorMessage();
      }
    }
  }
</script>

<p class="info-text">
  This data is only used for internal purposes and never shown anywhere
  publicly. (Your team and club administrators have access).
</p>

<form class="user-details-form" onsubmit={submitForm}>
  <input
    autocomplete="off"
    bind:value={form.id}
    class="input"
    name="id"
    readonly
    type="hidden"
  />

  <label class="label">
    <span>First Name(s)</span>
    <input
      autocomplete="given-name"
      bind:value={form.firstName}
      class="input"
      name="firstname"
      required
    />
  </label>

  <label class="label">
    <span>Last Name</span>
    <input
      autocomplete="family-name"
      bind:value={form.lastName}
      class="input"
      name="lastname"
      required
    />
  </label>

  <label class="label">
    <span>Display Name</span>

    <span class="info-text"> Will be set automatically if left blank. </span>

    <input bind:value={form.displayName} class="input" name="display_name" />
  </label>

  <div class="check-outer">
    <div class="check-inner">
      <input
        id="checkbox-emailVisibility"
        type="checkbox"
        class="checkbox"
        bind:checked={form.emailVisibility}
      />
      <label class="label" for="checkbox-emailVisibility"
        >Email Visibility</label
      >
    </div>
    <p class="info-text">
      Controls whether other app users can see your email address.
    </p>
  </div>

  <label class="label">
    <span>Profile Image</span>

    <span class="info-text">
      This is your app profile picture. It will not be used for your player
      profile image.
    </span>

    <input
      accept="image/*"
      bind:files
      class="input"
      name="avatar"
      type="file"
    />
  </label>

  {#if files}
    <h2>Selected file:</h2>
    {#each Array.from(files) as file}
      <p>{file.name} ({file.size} bytes)</p>
    {/each}
  {/if}

  <div class="actions">
    <button data-dialog-close class="btn preset-tonal-primary" type="submit"
      >Confirm</button
    >
  </div>
</form>

<style>
  .check-inner {
    display: flex;
    gap: calc(var(--spacing) * 3);
    align-items: center;
    margin-block-start: calc(var(--spacing) * 4);
  }

  .check-outer {
    .info-text {
      margin-block-start: var(--spacing);
    }
  }

  .info-text {
    margin-block: calc(var(--spacing) * 2);
    font-weight: var(--font-weight-light);
    font-size: var(--text-sm);
    display: block;
  }

  .user-details-form {
    margin-top: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .actions {
    display: flex;
    justify-content: flex-end;
    gap: calc(var(--spacing) * 3);
    margin-top: calc(var(--spacing) * 3);

    button {
      margin-top: calc(var(--spacing) * 2);
      border: 1px solid var(--color-primary-500);
    }
  }
</style>
