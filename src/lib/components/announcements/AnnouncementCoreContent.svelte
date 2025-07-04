<script lang="ts">
  import type { ExpandedAnnouncement } from "$lib/model/ExpandedResponse";
  import { User } from "lucide-svelte";
  import PriorityBadge from "./PriorityBadge.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
    textClasses?: string;
  }

  let { announcement, textClasses = "" }: Props = $props();
</script>

<header class="grid grid-cols-6">
  <h3 class="col-span-5 h4">{announcement.title}</h3>

  <PriorityBadge priority={announcement.priority} />
</header>

<p class="my-3 flex gap-3 {textClasses}">
  <User />
  <span>
    {announcement.expand?.author.first_name}
    {announcement.expand?.author.last_name}
  </span>
</p>

<p class={textClasses}>
  {@html announcement.bodytext}
</p>
