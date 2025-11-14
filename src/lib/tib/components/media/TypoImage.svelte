<script lang="ts">
  import type {Media} from "$lib/tib/types/MediaData.ts";

  interface Props {
    media: Media;
  }

  let {media}: Props = $props();
</script>

<img alt="{media?.alt}"
     class="h-auto max-w-full rounded-lg"
     loading="lazy"
     src="{media.url}"
     title="{media?.title}"
>