<script lang="ts">
  import type {PageProps} from './$types';
  import LicenseDetailRow from "$lib/components/club/licenses/LicenseDetailRow.svelte";
  import {User} from "lucide-svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">License Details</h1>

<article class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">
  <LicenseDetailRow
          categoryName="License Holder"
          rowValue={`${data.license?.person.first_name} ${data.license?.person.last_name}`}
  >
    {#snippet icon()}
      <User/>
    {/snippet}
  </LicenseDetailRow>

  <hr class="my-2">
</article>

