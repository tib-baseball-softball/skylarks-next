<script lang="ts">
  import type {UniformsetsCreate, UniformsetsResponse} from "$lib/model/pb-types";
  import {save} from "$lib/pocketbase/RecordOperations";
  import {invalidate} from "$app/navigation";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import {closeModal} from "$lib/functions/closeModal.ts";

  interface Props {
    uniformSet?: UniformsetsResponse | null;
    clubID: string;
  }


  let {uniformSet, clubID}: Props = $props();

  const form: UniformsetsCreate = $state(
      uniformSet ?? {
        id: "",
        name: "",
        cap: "",
        jersey: "",
        pants: "",
        club: "",
      }
  );

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();
    form.club = clubID;

    let result: UniformsetsResponse | null = null;

    try {
      result = await save<UniformsetsResponse>("uniformsets", form);
    } catch {
      toastController.triggerGenericFormErrorMessage("Uniform Set");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Uniform Set");
      await invalidate("club:single");
      closeModal();
    }
  }
</script>

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
    <span>Name</span>
    <input
            name="name"
            class="input"
            bind:value={form.name}
            required
            type="text"
    />
  </label>

  <label class="label">
    <span>Cap Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input class="input" type="color" bind:value={form.cap}/>
                <input class="input" type="text" bind:value={form.cap} tabindex="-1" required/>
            </span>
  </label>

  <div class="text-sm mt-4!">
    You can use the special identifiers <code class="code">B</code> and <code class="code">Bird</code> to get
    the Skylarks Home and Road Cap instead of a
    generic colored one.
  </div>

  <label class="label">
    <span>Jersey Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input class="input" type="color" bind:value={form.jersey}/>
                <input class="input" type="text" bind:value={form.jersey} tabindex="-1" required/>
            </span>
  </label>

  <label class="label">
    <span>Pants Color</span>
    <span class="grid grid-cols-[auto_1fr] gap-2">
                <input class="input" type="color" bind:value={form.pants}/>
                <input class="input" type="text" bind:value={form.pants} tabindex="-1" required/>
            </span>
  </label>

  <div class="text-sm mt-4!">
    Both hexadecimal color values (6 characters with leading '#') and
    <a class="anchor" href="https://developer.mozilla.org/en-US/docs/Web/CSS/named-color" target="_blank">
      named CSS colors
    </a>
    can be used.
  </div>

  <div class="flex justify-end gap-3 mt-3">
    <button type="submit" class="mt-2 btn preset-tonal-primary border border-primary-500">Submit</button>
  </div>
</form>
