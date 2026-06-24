<script lang="ts">
  import type { AnnouncementsResponse } from "$lib/dp/types/pb-types.ts";

  type Priority = AnnouncementsResponse["priority"];

  interface Props {
    priority: Priority;
  }

  let { priority }: Props = $props();

  const displayString = $derived.by(() => {
    switch (priority) {
      case "info":
        return "Info";
      case "warning":
        return "Important";
      case "danger":
        return "Critical";
    }
  });
</script>

<div class="root">
  <span class="badge" data-priority={priority}>
    {displayString}
  </span>
</div>

<style>
  .root {
    justify-self: flex-end;
    grid-column: span 1 / span 1;
  }

  .badge {
    background: var(--badge-bg);
    color: contrast-color(var(--badge-bg));

    &[data-priority="info"] {
      --badge-bg: light-dark(
        var(--color-secondary-50),
        var(--color-secondary-400)
      );
    }

    &[data-priority="warning"] {
      --badge-bg: var(--color-warning-50-950);
    }

    &[data-priority="danger"] {
      --badge-bg: var(--color-error-50-950);
    }
  }
</style>
