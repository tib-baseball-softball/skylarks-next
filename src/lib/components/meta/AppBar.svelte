<script lang="ts">
  import type {Snippet} from "svelte";

  interface Props {
    lead?: Snippet;
    children?: Snippet;
    trail?: Snippet;
  }

  let {lead, children, trail}: Props = $props();
</script>

<header
        class="w-full flex flex-col bg-surface-500/5 dark:bg-surface-800 space-y-4  p-3"
        data-testid="app-bar"
>
  <section
          class="flex justify-between grid-cols-[auto_1fr_auto] gap-4 "
          data-testid="app-bar-toolbar"
  >

    <div class="flex space-x-4 rtl:space-x-reverse  ">
      {@render lead?.()}
    </div>

    <div class="grow place-self-center">
      {@render children?.()}
    </div>

    <div class="flex space-x-4 rtl:space-x-reverse">
      {@render trail?.()}
    </div>

  </section>
</header>