<script lang="ts">
  import { MessageCircleMore } from "lucide-svelte";
  import ParticipationForm from "$lib/dp/components/forms/ParticipationForm.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import { authSettings } from "$lib/dp/client.svelte.js";
  import type {
    CustomAuthModel,
    ExpandedParticipation,
  } from "$lib/dp/types/ExpandedResponse.ts";
  import type { HTMLButtonAttributes } from "svelte/elements";

  interface Props {
    participation: ExpandedParticipation;
    isAdmin: boolean;
    classes?: string;
  }

  const { participation, isAdmin, classes = "" }: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  const canEdit = $derived(participation.user === authRecord.id || isAdmin);

  const id = Math.random().toString(36).substring(3, 16);

  const createdDate = $derived(
    participation.created !== ""
      ? new Date(participation.created).toLocaleString()
      : "---",
  );
  const updatedDate = $derived(
    participation.updated !== ""
      ? new Date(participation.updated).toLocaleString()
      : "---",
  );

  const triggerProps: {
    popovertarget: HTMLButtonAttributes["popovertarget"];
    popovertargetaction: HTMLButtonAttributes["popovertargetaction"];
  } = {
    popovertarget: id,
    popovertargetaction: "hide",
  };
</script>

<button class="{classes} trigger-button" popovertarget={id}>
  {#if participation.comment}
    <span class="badge-icon preset-filled-surface-400-600">
      <MessageCircleMore size="13" />
    </span>
  {/if}
  {participation?.expand?.user?.display_name ||
    participation?.expand?.user?.first_name}
</button>

<div class="popover-content card shadow-2xl" popover="auto" {id}>
  <div class="item-container">
    <p class="label-text">Created:</p>
    <p>{createdDate}</p>
  </div>

  <div class="item-container">
    <p class="label-text">Last Updated:</p>
    <p>{updatedDate}</p>
  </div>

  <div class="item-container">
    <p class="label-text">Comment/Reason:</p>
    <p>{participation.comment}</p>
  </div>

  {#if canEdit}
    <Dialog
      triggerClasses="btn btn-sm preset-tonal-primary border border-primary-500 edit-btn"
      {triggerProps}
    >
      {#snippet triggerContent()}
        Edit
      {/snippet}

      {#snippet title()}
        <header>
          <h2 class="h3">
            Edit participation data for {participation?.expand?.user
              ?.first_name}
            {participation?.expand?.user?.last_name}
          </h2>
        </header>
      {/snippet}

      <ParticipationForm {participation} />
    </Dialog>
  {/if}
</div>

<style>
  .trigger-button {
    position: relative;
  }

  .popover-content {
    position: absolute;
    position-area: bottom;
    top: anchor(right, var(--spacing));
    border: 1px solid;
    color: light-dark(black, white);
    background-color: light-dark(
      var(--color-surface-50),
      var(--color-surface-800)
    );
    padding: calc(var(--spacing) * 4);
    font-size: var(--text-sm);
    line-height: var(--tw-leading, var(--text-sm--line-height));
    max-width: calc(var(--spacing) * 80);

    &:popover-open {
      transform: translateY(0) scale(1);
      opacity: 1;

      @starting-style {
        transform: translateY(30px) scale(0);
        opacity: 0;
      }
    }

    transform: translateY(-50px);
    opacity: 0;

    transition:
      transform,
      opacity,
      display allow-discrete,
      overlay allow-discrete;
    transition-duration: 0.2s;

    @media (prefers-reduced-motion: reduce) {
      transition: none;
    }

    :global(.edit-btn) {
      margin-top: calc(var(--spacing) * 2);
    }
  }

  .item-container {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    justify-content: space-between;
    margin-block: 0.3em;
  }

  .label-text {
    font-weight: var(--font-weight-light);
    font-size: var(--text-xs);
    line-height: var(--tw-leading, var(--text-xs--line-height));
  }

  .badge-icon {
    color: white;
    top: -15px;
    right: -11px;
    z-index: 10;
    position: absolute;
  }
</style>
