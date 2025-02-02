<script lang="ts">
  import type {License} from "bsm.js";
  import LicenseDetailRow from "$lib/components/club/licenses/LicenseDetailRow.svelte";
  import {Calendar1, CalendarClock, Hash, User} from "lucide-svelte";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.ts";

  interface Props {
    license: License,
  }

  let {license}: Props = $props();
  let expiryDate = $derived(new Date(license.valid_until));
</script>

<article class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">
  <LicenseDetailRow
          categoryName="Holder"
          rowValue={`${license?.person.first_name} ${license?.person.last_name}`}
  >
    {#snippet icon()}
      <User/>
    {/snippet}
  </LicenseDetailRow>

  <hr class="my-2">

  <LicenseDetailRow
          categoryName="Number"
          rowValue={license.number}
  >
    {#snippet icon()}
      <Hash/>
    {/snippet}
  </LicenseDetailRow>

  <hr class="my-2">

  <LicenseDetailRow
          categoryName="Valid until"
          rowValue={expiryDate.toLocaleDateString()}
  >
    {#snippet icon()}
      <Calendar1/>
    {/snippet}
  </LicenseDetailRow>

  <hr class="my-2">

  <LicenseDetailRow
          categoryName="Relative Time"
          rowValue={DateTimeUtility.getRelativeTimeString(expiryDate)}
  >
    {#snippet icon()}
      <CalendarClock/>
    {/snippet}
  </LicenseDetailRow>

</article>