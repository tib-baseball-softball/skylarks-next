<script lang="ts">
  import {invalidate} from "$app/navigation";
  import type {ExpandedParticipation} from "$lib/dp/types/ExpandedResponse.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {sendParticipationData} from "$lib/dp/utility/sendParticipationData.ts";

  interface Props {
    participation: ExpandedParticipation;
  }

  let {participation}: Props = $props();

  let form = $derived.by(() => {
    const formData = $state(participation);
    return formData;
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    await sendParticipationData(form).then(() => closeModal());
    await invalidate("event:list");
  }

  const createdDate = $derived(
    participation.created !== "" ? new Date(participation.created).toLocaleString() : "---"
  );
  const updatedDate = $derived(
    participation.updated !== "" ? new Date(participation.updated).toLocaleString() : "---"
  );
</script>

<form class="participation-form" onsubmit={submitForm}>
  <div class="form-grid">

    <dl class="info-item">
      <dt>Created</dt>
      <dd>{createdDate}</dd>
    </dl>

    <dl class="info-item">
      <dt>Last Updated</dt>
      <dd>{updatedDate}</dd>
    </dl>

    <input
      autocomplete="off"
      bind:value={form.id}
      class="input"
      name="id"
      readonly
      type="hidden"
    />

    <label class="label label-full">
      <span class="label-text">Comment</span>
      <input
        autocomplete="off"
        bind:value={form.comment}
        class="input"
        name="id"
        placeholder="background info about your attendance"
        type="text"
      />
    </label>

    <div class="label label-full">
      <span>State</span>

      <span class="btn-group state-selector">
        <button
          class={["btn btn-in", form.state === "in" && "preset-filled-success-500 text-black"]}
          onclick={() => form.state = "in"}
          type="button"
        >
          <span>In</span>
        </button>
        <button
          class={["btn btn-maybe", form.state === "maybe" && "preset-filled-warning-500 text-black"]}
          onclick={() => form.state = "maybe"}
          type="button"
        >
          <span>Maybe</span>
        </button>
        <button
          class={["btn btn-out", form.state === "out" && "preset-filled-error-500 text-white"]}
          onclick={() => form.state = "out"}
          type="button"
        >
          <span>Out</span>
        </button>
      </span>
    </div>

    <div class="submit-container">
      <button class="btn preset-filled-primary-500" type="submit">
        Submit
      </button>
    </div>
  </div>
</form>

<style>
  .participation-form {
    margin-top: calc(var(--spacing) * 4);
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 3);
  }

  .form-grid {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: calc(var(--spacing) * 2);

    @media (min-width: 48rem) {
      gap: calc(var(--spacing) * 3);
    }

    @media (min-width: 64rem) {
      gap: calc(var(--spacing) * 4);
    }
  }

  .info-item {
    display: flex;
    flex-direction: column;
    gap: calc(var(--spacing) * 1);

    dt {
      font-size: var(--text-sm);
      font-weight: 300;
    }
  }

  .label-full {
    grid-column: span 2 / span 2;
  }

  .label-text {
    display: block;
  }

  .state-selector {
    display: flex;
    justify-items: stretch;
    border: 1px solid;

    button {
      flex: 1 1 0%;
      
      &.btn-in:hover {
          background-color: var(--color-success-50-950);
          color: var(--color-success-950-50);
      }
      
      &.btn-maybe:hover {
          background-color: var(--color-warning-50-950);
          color: var(--color-warning-950-50);
      }
      
      &.btn-out:hover {
          background-color: var(--color-error-50-950);
          color: var(--color-error-950-50);
      }
    }
  }

  .submit-container {
    display: flex;
    justify-content: center;
    grid-column: span 2 / span 2;

    button {
      margin-top: calc(var(--spacing) * 2);
    }
  }
</style>
