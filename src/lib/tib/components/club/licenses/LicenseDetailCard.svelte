<script lang="ts">
  import type {License} from "bsm.js";
  import LicenseLevelIndicator from "$lib/tib/components/club/licenses/LicenseLevelIndicator.svelte";
  import {Shirt} from "lucide-svelte";
  import GenericDetailRow from "$lib/dp/components/utils/GenericDetailRow.svelte";

  interface Props {
    license: License;
  }

  let {license}: Props = $props();
</script>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow
          categoryName="Level"
          rowValue={license?.level ?? ""}
  >
    {#snippet icon()}
      <LicenseLevelIndicator level={license?.level ?? ""}/>
    {/snippet}
  </GenericDetailRow>

  {#if license?.baseball}
    <hr class="my-2">
    <GenericDetailRow
            categoryName="Baseball"
            rowValue={"yes"}
    >
      {#snippet icon()}
        <div class="ms-1">⚾️</div>
      {/snippet}
    </GenericDetailRow>

  {/if}

  {#if license?.softball}
    <hr class="my-2">
    <GenericDetailRow
            categoryName="Softball"
            rowValue={"yes"}
    >
      {#snippet icon()}
        <div class="ms-1">🥎</div>
      {/snippet}
    </GenericDetailRow>
  {/if}

  {#if license?.sleeve_number}
    <hr class="my-2">
    <GenericDetailRow
            categoryName="Sleeve Number"
            rowValue={license?.sleeve_number ?? ""}
    >
      {#snippet icon()}
        <Shirt/>
      {/snippet}
    </GenericDetailRow>
  {/if}

</article>

<style>
    article {
        margin-bottom: 1.5em;
    }
</style>