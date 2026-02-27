<script lang="ts">
  import {invalidate} from "$app/navigation";
  import {save} from "$lib/dp/records/RecordOperations.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {UniformsetsCreate, UniformsetsResponse} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";

  interface Props {
    uniformSet?: UniformsetsResponse | null;
    clubID: string;
  }

  const {uniformSet, clubID}: Props = $props();

  function formFromProps(data: UniformsetsResponse | null | undefined): UniformsetsCreate {
    return data ?? {
      id: "",
      name: "",
      cap: "",
      jersey: "",
      pants: "",
      club: "",
    };
  }

  let form: UniformsetsCreate = $derived.by(() => {
    const formData = $state(formFromProps(uniformSet));
    return formData;
  });

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

<form onsubmit={submitForm}>
  <input
    autocomplete="off"
    bind:value={form.id}
    class="input"
    name="id"
    readonly
    type="hidden"
  />

  <label class="label">
    <span>Name</span>
    <input
      bind:value={form.name}
      class="input"
      name="name"
      required
      type="text"
    />
  </label>

  <label class="label">
    <span>Cap Color</span>
    <span class="color-inputs">
                <input bind:value={form.cap} class="input" type="color"/>
                <input bind:value={form.cap} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <div class="hint">
    You can use the special identifiers <code class="code">B</code> and <code class="code">Bird</code> to get
    the Skylarks Home and Road Cap instead of a
    generic colored one.
  </div>

  <label class="label">
    <span>Jersey Color</span>
    <span class="color-inputs">
                <input bind:value={form.jersey} class="input" type="color"/>
                <input bind:value={form.jersey} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <label class="label">
    <span>Pants Color</span>
    <span class="color-inputs">
                <input bind:value={form.pants} class="input" type="color"/>
                <input bind:value={form.pants} class="input" required tabindex="-1" type="text"/>
            </span>
  </label>

  <div class="hint">
    Both hexadecimal color values (6 characters with leading '#') and
    <a class="anchor" href="https://developer.mozilla.org/en-US/docs/Web/CSS/named-color" target="_blank">
      named CSS colors
    </a>
    can be used.
  </div>

  <div class="submit-wrapper">
    <button class="btn preset-tonal-primary border-primary" type="submit">Submit</button>
  </div>
</form>

<style>
  form {
    margin-block-start: calc(var(--spacing) * 4);
  }

  .color-inputs {
    display: grid;
    grid-template-columns: auto 1fr;
    gap: calc(var(--spacing) * 2);
  }

  input[type=color] {
    border: 1px var(--tw-border-style);
  }

  .hint {
    font-size: var(--text-sm);
    margin-block-start: calc(var(--spacing) * 4);
  }

  .submit-wrapper {
    display: flex;
    justify-content: end;
    gap: calc(var(--spacing) * 3);
    margin-block-start: calc(var(--spacing) * 3);
  }

  .label, .hint {
    margin-block-start: calc(var(--spacing) * 3);
  }
</style>
