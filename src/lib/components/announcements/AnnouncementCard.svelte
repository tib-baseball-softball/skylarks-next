<script lang="ts">
  import type {ExpandedAnnouncement,} from "$lib/model/ExpandedResponse";
  import AnnouncementCoreContent from "./AnnouncementCoreContent.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let {announcement}: Props = $props();
</script>

<article class="card p-4 preset-tonal-surface shadow-xl h-full">
  <div class="flex flex-col justify-between h-full">
    <div>
      <AnnouncementCoreContent {announcement} textClasses="line-clamp-5"/>
    </div>

    <div>
      <hr class="mt-6 mb-2"/>

      <footer class="mt-3 flex gap-2 justify-end">
        <a
                href="/account/announcements/{announcement.id}"
                class="btn btn-sm preset-filled-primary-500">Read more</a
        >
      </footer>
    </div>
  </div>
</article>
