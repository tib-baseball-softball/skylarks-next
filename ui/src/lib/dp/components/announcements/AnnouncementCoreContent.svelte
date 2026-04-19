<script lang="ts">
  import {User} from "lucide-svelte";
  import type {ExpandedAnnouncement} from "$lib/dp/types/ExpandedResponse.ts";
  import PriorityBadge from "./PriorityBadge.svelte";
  import LocalDate from "$lib/dp/components/utils/LocalDate.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let {announcement}: Props = $props();
</script>

<header class="header">
  <h3 class="h5 clamped title">{announcement.title}</h3>

  <PriorityBadge priority={announcement.priority}/>
</header>

<time datetime={announcement.created}>
  <LocalDate date={announcement.created}/>
</time>

<p class="author-info">
  <User/>
  <span>
    {announcement.expand?.author?.first_name}
    {announcement.expand?.author?.last_name}
  </span>
</p>

<p class="body-text clamped">
  {@html announcement.bodytext}
</p>

<style>
  time {
    font-weight: var(--font-weight-light);
  }

  .header {
    display: flex;
    justify-content: space-between;
    hyphens: auto;
  }

  .author-info {
    display: flex;
    gap: calc(var(--spacing) * 3);
  }

  .body-text {
    line-clamp: 5;
    -webkit-line-clamp: 5;
  }

  .title {
    line-clamp: 2;
    -webkit-line-clamp: 2;
  }
</style>
