<script lang="ts">
  import type {
    CustomAuthModel,
    ExpandedAnnouncement,
    ExpandedClub,
  } from "$lib/model/ExpandedResponse";
  import { authSettings, client } from "$lib/pocketbase/index.svelte";
  import { User } from "lucide-svelte";
  import PriorityBadge from "../announcements/PriorityBadge.svelte";
  import Dialog from "../utility/Dialog.svelte";
  import DeleteButton from "../utility/DeleteButton.svelte";
  import AnnouncementCoreContent from "../announcements/AnnouncementCoreContent.svelte";

  interface Props {
    announcement: ExpandedAnnouncement;
  }

  let { announcement }: Props = $props();
</script>
