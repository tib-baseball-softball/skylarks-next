<script lang="ts">
import GenericDetailRow from "$lib/components/utility/GenericDetailRow.svelte"
import type { ClubFunction } from "bsm.js"
import ClubOfficialIcon from "$lib/components/club/officials/ClubOfficialIcon.svelte"
import { Calendar, Diamond, Mail, Tag } from "lucide-svelte"

interface Props {
  clubOfficial: ClubFunction
}

let { clubOfficial }: Props = $props()

const admissionDate = $derived(new Date(clubOfficial.admission_date))
</script>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <GenericDetailRow
          categoryName="Title"
          rowValue={clubOfficial.function}
  >
    {#snippet icon()}
      <ClubOfficialIcon {clubOfficial}/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Category"
          rowValue={clubOfficial.category}
  >
    {#snippet icon()}
      <Tag/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Name"
          rowValue={`${clubOfficial.person.last_name}, ${clubOfficial.person.first_name}`}
  >
    {#snippet icon()}
      <Diamond/>
    {/snippet}
  </GenericDetailRow>

  <hr class="my-2">

  <GenericDetailRow
          categoryName="Since"
          rowValue={admissionDate.toLocaleDateString()}
  >
    {#snippet icon()}
      <Calendar/>
    {/snippet}
  </GenericDetailRow>

</article>

<article class="card p-3 preset-tonal dark:border dark:border-surface-500 shadow-xl">
  <a class="anchor" href="mailto:{clubOfficial.mail}">
    <GenericDetailRow
            categoryName="Contact"
            rowValue={clubOfficial.mail}
    >
      {#snippet icon()}
        <Mail/>
      {/snippet}
    </GenericDetailRow>
  </a>
</article>

<style>
    article {
        margin-bottom: 1.5em;
    }
</style>