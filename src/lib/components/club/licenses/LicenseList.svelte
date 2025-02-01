<script lang="ts">
  import type {License} from "bsm.js";
  import LicenseRow from "$lib/components/club/licenses/LicenseRow.svelte";

  interface Props {
    licenses: License[]
  }

  let {licenses}: Props = $props()
  let baseballLicenses = $derived(licenses.filter(license => license.baseball === true))
  let softballLicenses = $derived(licenses.filter(license => license.softball === true))
</script>

<article class="!mb-4">
  <section>
    <header class="mb-2">
      <h2 class="h2">Baseball</h2>
    </header>

    <ul class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">

      {#each baseballLicenses as license, index}
        <li>
          <LicenseRow {license}/>
        </li>

        {#if index + 1 < baseballLicenses.length}
          <hr class="my-2">
        {/if}
      {:else}
        <li>
          <p>No baseball licenses available.</p>
        </li>
      {/each}
    </ul>
  </section>

  <section class="mt-3">
    <header class="mb-2">
      <h2 class="h2">Softball</h2>
    </header>

    <ul class="card p-3 variant-soft dark:border dark:border-surface-500 shadow-xl">

      {#each softballLicenses as license, index}
        <li>
          <LicenseRow {license}/>
        </li>

        {#if index + 1 < softballLicenses.length}
          <hr class="my-2">
        {/if}
      {:else}
        <li>
          <p>No softball licenses available.</p>
        </li>
      {/each}
    </ul>
  </section>
</article>