<script lang="ts">
  import type {ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {Segment } from "@skeletonlabs/skeleton-svelte";
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

    <label class="label col-span-2">
      <span class="block">State</span>

      <Segment display="flex">
        <Segment.Item fill="bg-success-500!" hover="hover:preset-tonal-success" bind:group={form.state}
                   name="state" value="in">
          In
        </Segment.Item>
        <Segment.Item fill="bg-warning-500!" hover="hover:preset-tonal-warning" bind:group={form.state}
                   name="state"
                   value="maybe">
          Maybe
        </Segment.Item>
        <Segment.Item color="text-white!" fill="bg-error-500!" hover="hover:preset-tonal-error"
                   bind:group={form.state}
                   name="state" value="out">
          Out
        </Segment.Item>
      </Segment>
    </label>

    <div class="flex justify-center col-span-2">
      <button type="submit" class="mt-2 btn preset-tonal-primary border border-primary-500">
        Submit
      </button>
    </div>
  </div>
</form>