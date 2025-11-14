<script lang="ts">
  import type {License} from "bsm.js";
  import {Calendar1, CalendarClock, Hash, User} from "lucide-svelte";
  import GenericDetailRow from "$lib/components/utility/GenericDetailRow.svelte";
  import {DateTimeUtility} from "$lib/dp/service/DateTimeUtility.ts";

  interface Props {
    license: License;
  }

  const {license}: Props = $props();
  const expiryDate = $derived(new Date(license.valid_until));
</script>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow
          categoryName="Holder"
          rowValue={`${license?.person.first_name} ${license?.person.last_name}`}
  >
    {#snippet icon()}
      <User/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Number"
          rowValue={license.number}
  >
    {#snippet icon()}
      <Hash/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Valid until"
          rowValue={expiryDate.toLocaleDateString()}
  >
    {#snippet icon()}
      <Calendar1/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Relative Time"
          rowValue={DateTimeUtility.getRelativeTimeString(expiryDate)}
  >
    {#snippet icon()}
      <CalendarClock/>
    {/snippet}
  </GenericDetailRow>

</article>