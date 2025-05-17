<script lang="ts">
  import type {Snippet} from "svelte";
  import {Dialog, type WithoutChild} from "bits-ui";
  import {fly} from "svelte/transition";
  import {cubicInOut} from "svelte/easing";
  import {X} from "lucide-svelte";

  /**
   * Modal dialog based on headless bits-ui building blocks.
   * Potential replacement for Skeleton modal.
   */

  type Props = Dialog.RootProps & {
    triggerClasses?: string;
    title: Snippet;
    description?: Snippet;
    contentProps?: WithoutChild<Dialog.ContentProps>;
    triggerContent: Snippet;
    closeButtonClasses?: string;
    disabled?: boolean;
  };

  let {
    open = $bindable(false),
    children,
    triggerClasses = "",
    contentProps,
    title,
    description,
    triggerContent,
    closeButtonClasses = "",
    disabled = false,
    ...restProps
  }: Props = $props();
</script>

<Dialog.Root bind:open {...restProps}>
  <Dialog.Trigger class="{triggerClasses} flex gap-1" type="button" disabled={disabled}>
    {@render triggerContent()}
  </Dialog.Trigger>
  <Dialog.Portal>
    <Dialog.Overlay
            class="fixed inset-0 z-50 bg-surface-50/50 dark:bg-surface-950/50"/>
    <Dialog.Content forceMount
                    class="card bg-surface-50 dark:bg-surface-800 border shadow-2xl fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] p-5 sm:max-w-[640px] md:w-full"
                    {...contentProps}>

      {#snippet child({props, open})}
        {#if open}
          <div {...props} transition:fly={{y: 150, duration: 100, easing: cubicInOut}}>
            <div class="flex gap-5 items-center mb-2">

              <Dialog.Close class="btn preset-tonal-surface border border-surface-500 {closeButtonClasses}">
                <X/>
              </Dialog.Close>

              <Dialog.Title class="text-xl font-semibold">
                {@render title()}
              </Dialog.Title>

            </div>

            {#if description}
              <Dialog.Description>
                {@render description?.()}
              </Dialog.Description>
            {/if}

            {@render children?.()}
          </div>
        {/if}
      {/snippet}

    </Dialog.Content>
  </Dialog.Portal>
</Dialog.Root>

