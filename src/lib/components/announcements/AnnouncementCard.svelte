<script lang="ts">
  import type { ExpandedAnnouncement } from "$lib/model/ExpandedResponse";
  import { Pencil } from "lucide-svelte";
  import PriorityBadge from "./PriorityBadge.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let { announcement }: Props = $props();
</script>

<article class="card p-4 preset-tonal-surface shadow-xl">
  <header class="grid grid-cols-6">
    <h3 class="col-span-5 h4">{announcement.title}</h3>

    <PriorityBadge priority={announcement.priority} />
  </header>

  <p class="my-3 flex gap-3">
    <Pencil />
    <span>
      {announcement.expand?.author.first_name}
      {announcement.expand?.author.last_name}
    </span>
  </p>

  <p class="text-sm">
    {@html announcement.bodytext}
  </p>
</article>
