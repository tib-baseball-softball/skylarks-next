<script lang="ts">
  import type {
    ExpandedAnnouncement,
    ExpandedClub,
  } from "$lib/model/ExpandedResponse";
  import Dialog from "../utility/Dialog.svelte";
  import AnnouncementCoreContent from "./AnnouncementCoreContent.svelte";
  import AnnouncementDialog from "../forms/AnnouncementForm.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let { announcement }: Props = $props();
</script>

<article class="card p-4 preset-tonal-surface shadow-xl">
  <AnnouncementCoreContent {announcement} textClasses="line-clamp-5" />

  <hr class="mt-6 mb-2" />

  <footer class="mt-3 flex gap-2 justify-end">
    <a
      href="/account/announcements/{announcement.id}"
      class="btn btn-sm preset-filled-primary-500">Read more</a
    >
  </footer>
</article>
