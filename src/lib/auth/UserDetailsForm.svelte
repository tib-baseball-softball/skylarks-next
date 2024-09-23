<script lang="ts">
  import type { CustomAuthModel } from "$lib/model/ExpandedResponse";
  import { client } from "$lib/pocketbase";
  import { authModel } from "$lib/pocketbase/Auth";

  interface props {
    parent: any;
  }

  const model = $authModel as CustomAuthModel;

  let { parent }: props = $props();

  let files: FileList | null = $state(null);

  const form = $state({
    id: model.id,
    email: model.email,
    firstName: model.first_name,
    lastName: model.last_name,
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    parent.onClose();

    const formData = new FormData();

    formData.append("id", form.id);
    formData.append("email", form.email);

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

    client.collection("users").update(form.id, formData, { expand: "club" });
  }
</script>

<div class="card p-6">
  <header class="text-xl font-semibold">
    Edit User Data for {model.first_name}
    {model.last_name}
  </header>

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
      E-Mail
      <input
        name="email"
        class="input"
        bind:value={form.email}
        required
        type="email"
      />
    </label>

    <label class="label">
      First Name(s)
      <input name="firstname" class="input" bind:value={form.firstName} />
    </label>

    <label class="label">
      Last Name
      <input name="lastname" class="input" bind:value={form.lastName} />
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

    <div class="flex justify-end gap-3 mt-3">
      <button type="submit" class="mt-2 btn variant-ghost-primary"
        >Confirm</button
      >
    </div>
  </form>
</div>
