<script lang="ts">
  import type {ExpandedParticipation} from "$lib/model/ExpandedResponse";
  import {RadioGroup, RadioItem} from "@skeletonlabs/skeleton";
  import {sendParticipationData} from "$lib/functions/sendParticipationData";
  import {invalidate} from "$app/navigation";
  import {X} from "lucide-svelte";

  interface Props {
    participation: ExpandedParticipation,
    parent: any,
  }

  let {participation, parent}: Props = $props();

  const form = $state(participation);

  async function closeModal() {
    // @ts-ignore
    parent.onClose();
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    await sendParticipationData(form).then(() => closeModal());
    await invalidate("event:list");
  }
</script>

<div class="w-modal-slim">
  <article class="card p-4">

    <div class="flex items-center gap-5">
      <button
              aria-label="cancel and close"
              class="btn variant-ghost-surface"
              onclick={() => closeModal()}
      >
        <X/>
      </button>
      <header class="text-xl font-semibold">
        <h2 class="h3">Edit participation data
          for {participation?.expand?.user?.first_name} {participation?.expand?.user?.last_name}</h2>
      </header>
    </div>

    <form onsubmit={submitForm} class="mt-4 space-y-3">
      <div class="grid grid-cols-2 gap-2 md:gap-3 lg:gap-4">

        <div class="space-y-1">
          <p class="text-sm font-light">Created</p>
          <p>{new Date(participation.created).toLocaleString()}</p>
        </div>

        <div class="space-y-1">
          <p class="text-sm font-light">Last Updated</p>
          <p>{new Date(participation.updated).toLocaleString()}</p>
        </div>

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

          <RadioGroup display="flex">
            <RadioItem fill="!bg-success-500" hover="hover:variant-soft-success" bind:group={form.state}
                       name="state" value="in">
              In
            </RadioItem>
            <RadioItem fill="!bg-warning-500" hover="hover:variant-soft-warning" bind:group={form.state}
                       name="state"
                       value="maybe">
              Maybe
            </RadioItem>
            <RadioItem color="!text-white" fill="!bg-error-500" hover="hover:variant-soft-error"
                       bind:group={form.state}
                       name="state" value="out">
              Out
            </RadioItem>
          </RadioGroup>
        </label>

        <div class="flex justify-center col-span-2">
          <button type="submit" class="mt-2 btn variant-ghost-primary">
            Submit
          </button>
        </div>
      </div>
    </form>

  </article>
</div>