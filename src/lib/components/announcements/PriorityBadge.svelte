<script lang="ts">
  import type { AnnouncementsResponse } from "$lib/model/pb-types";
  type Priority = AnnouncementsResponse["priority"];

  interface Props {
    priority: Priority;
  }

  let { priority }: Props = $props();

  const badgeClass = $derived.by(() => {
    switch (priority) {
      case "info":
        return "preset-outlined-secondary-500";
      case "warning":
        return "preset-outlined-warning-500";
      case "danger":
        return "preset-outlined-error-500";
    }
  });
</script>

<div class="col-span-1 justify-self-end">
  <span class="badge {badgeClass}">MÖP</span>
</div>
