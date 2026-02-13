<script lang="ts">
  import type {AnnouncementsResponse} from "$lib/dp/types/pb-types.ts";

  type Priority = AnnouncementsResponse["priority"]

  interface Props {
    priority: Priority;
  }

  let {priority}: Props = $props();

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
    &[data-priority="info"] {
      border: 1px solid white;
      color: var(--color-secondary-500);
    }

    &[data-priority="warning"] {
      border: 1px solid var(--color-warning-500);
      color: var(--color-warning-500);
    }

    &[data-priority="danger"] {
      border: 1px solid var(--color-error-500);
      color: var(--color-error-500);
    }
  }
</style>
