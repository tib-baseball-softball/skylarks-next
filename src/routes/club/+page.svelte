<script lang="ts">
  import ClubInfoSection from "$lib/components/club/ClubInfoSection.svelte";
  import ClubGridItem from "$lib/components/club/ClubGridItem.svelte";
  import type {PageProps} from './$types';
  import {BookText, Diamond, Pencil, Users} from "lucide-svelte";

  let {data}: PageProps = $props();
</script>

<!-- Club Info Section -->
<ClubInfoSection club={data.bsmClubData}/>

<!-- Grid Section -->
<div class="grid grid-cols-2 md:grid-cols-3 gap-4 lg:gap-6 mb-4!">
  <ClubGridItem
          href="/club/details"
          label="Details"
          description="More Info"
  >
    {#snippet icon()}
      <BookText/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/teams"
          label="Teams"
          description="Adult and Youth"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/scorer"
          label="Scorer"
          description="Keeping statistics"
  >
    {#snippet icon()}
      <Pencil/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/umpire"
          label="Umpire"
          description="Our Referees"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/ballpark"
          label="Ballpark"
          description="The Lark's Nest"
  >
    {#snippet icon()}
      <Diamond/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/officials"
          label="Officials"
          description="Club Administration"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>
</div>
