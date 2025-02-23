<script lang="ts">
  import type {Media} from "$lib/model/MediaData.ts";

  interface Props {
    media: Media;
  }

  let {media}: Props = $props();
</script>

<img class="h-auto max-w-full rounded-lg"
     loading="lazy"
     src="{media.url}"
     alt="{media?.alt}"
     title="{media?.title}"
>