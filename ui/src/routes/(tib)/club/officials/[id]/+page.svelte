<script lang="ts">
  import type {PageProps} from "../../../../../../.svelte-kit/types/src/routes";
  import ClubOfficialDetailCard from "$lib/tib/components/club/officials/ClubOfficialDetailCard.svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Official Details</h1>

<ClubOfficialDetailCard clubOfficial={data.clubOfficial}/>