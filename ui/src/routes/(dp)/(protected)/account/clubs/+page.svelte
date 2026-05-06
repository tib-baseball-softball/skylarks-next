<script lang="ts">
  import type { PageProps } from "./$types";

  const {data}: PageProps = $props();
</script>

{#each data.clubs as club}
	<p>{club.name}</p>
{/each}