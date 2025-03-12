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
    buttonText: string;
    title: Snippet;
    description: Snippet;
    contentProps?: WithoutChild<Dialog.ContentProps>;
    triggerIcon: Snippet;
  };

  let {
    open = $bindable(false),
    children,
    buttonText,
    contentProps,
    title,
    description,
    triggerIcon,
    ...restProps
  }: Props = $props();
</script>

<Dialog.Root bind:open {...restProps}>
  <Dialog.Trigger class="btn variant-ghost-primary flex gap-1">
    {@render triggerIcon()}
    {buttonText}
  </Dialog.Trigger>
  <Dialog.Portal>
    <Dialog.Overlay
            class="fixed inset-0 z-50 bg-surface-backdrop-token"/>
    <Dialog.Content forceMount
                    class="fixed left-[50%] top-[50%] z-50 w-full max-w-[94%] translate-x-[-50%] translate-y-[-50%] card p-5 sm:max-w-[640px] md:w-full"
                    {...contentProps}>

      {#snippet child({props, open})}
        {#if open}
          <div {...props} transition:fly={{y: 150, duration: 100, easing: cubicInOut}}>
            <div class="flex gap-5 items-center mb-2">

              <Dialog.Close class="btn variant-ghost-surface">
                <X/>
              </Dialog.Close>

              <Dialog.Title class="text-xl font-semibold">
                {@render title()}
              </Dialog.Title>

            </div>

            <Dialog.Description>
              {@render description()}
            </Dialog.Description>

            {@render children?.()}
          </div>
        {/if}
      {/snippet}

    </Dialog.Content>
  </Dialog.Portal>
</Dialog.Root>

