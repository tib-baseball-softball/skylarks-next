<script lang="ts">
  import type {ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {sendParticipationData} from "$lib/functions/sendParticipationData";
  import {invalidate} from "$app/navigation";
  import {closeModal} from "$lib/functions/closeModal.ts";

  interface Props {
    participation: ExpandedParticipation,
  }

  let {participation}: Props = $props();

  const form = $state(participation);

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    await sendParticipationData(form).then(() => closeModal());
    await invalidate("event:list");
  }
</script>

<form onsubmit={submitForm} class="mt-4 space-y-3">
  <div class="grid grid-cols-2 gap-2 md:gap-3 lg:gap-4">

    <dl class="space-y-1">
      <dt class="text-sm font-light">Created</dt>
      <dd>{new Date(participation.created).toLocaleString()}</dd>
    </dl>

    <dl class="space-y-1">
      <dt class="text-sm font-light">Last Updated</dt>
      <dd>{new Date(participation.updated).toLocaleString()}</dd>
    </dl>

    <input
            name="id"
            autocomplete="off"
            class="input"
            type="hidden"
            readonly
            bind:value={form.id}
    />

    <label class="label col-span-2">
      <span class="block">Comment</span>
      <input
              name="id"
              autocomplete="off"
              class="input "
              type="text"
              placeholder="background info about your attendance"
              bind:value={form.comment}
      />
    </label>

    <div class="label col-span-2">
      State

      <span class="flex justify-items-stretch btn-group preset-outlined-surface-200-800">
        <button
                type="button"
                class={["btn hover:preset-ghost-success flex-grow", form.state === "in" && "preset-filled-success-500 text-black"]}
                onclick={() => form.state = "in"}
        >
          In
        </button>
        <button
                type="button"
                class={["btn hover:preset-ghost-warning flex-grow", form.state === "maybe" && "preset-filled-warning-500 text-black"]}
                onclick={() => form.state = "maybe"}
        >
          Maybe
        </button>
        <button
                type="button"
                class={["btn hover:preset-ghost-error flex-grow", form.state === "out" && "preset-filled-error-500 text-white"]}
                onclick={() => form.state = "out"}
        >
          Out
        </button>
      </span>
    </div>

    <div class="flex justify-center col-span-2">
      <button type="submit" class="mt-2 btn preset-filled-primary-500">
        Submit
      </button>
    </div>
  </div>
</form>
