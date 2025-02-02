<script lang="ts">
  import type {License} from "bsm.js";
  import LicenseLevelIndicator from "$lib/components/club/licenses/LicenseLevelIndicator.svelte";
  import LicenseDetailRow from "$lib/components/club/licenses/LicenseDetailRow.svelte";
  import {Shirt} from "lucide-svelte";

  interface Props {
    license: License,
  }

  let {license}: Props = $props();
</script>

<article class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">
  <LicenseDetailRow
          categoryName="Level"
          rowValue={license?.level ?? ""}
  >
    {#snippet icon()}
      <LicenseLevelIndicator level={license?.level ?? ""}/>
    {/snippet}
  </LicenseDetailRow>

  {#if license?.baseball}
    <hr class="my-2">
    <LicenseDetailRow
            categoryName="Baseball"
            rowValue={"yes"}
    >
      {#snippet icon()}
        <div class="ms-1">⚾️</div>
      {/snippet}
    </LicenseDetailRow>

  {/if}

  {#if license?.softball}
    <hr class="my-2">
    <LicenseDetailRow
            categoryName="Softball"
            rowValue={"yes"}
    >
      {#snippet icon()}
        <div class="ms-1">🥎</div>
      {/snippet}
    </LicenseDetailRow>
  {/if}

  {#if license?.sleeve_number}
    <hr class="my-2">
    <LicenseDetailRow
            categoryName="Sleeve Number"
            rowValue={license?.sleeve_number ?? ""}
    >
      {#snippet icon()}
        <Shirt/>
      {/snippet}
    </LicenseDetailRow>
  {/if}

</article>