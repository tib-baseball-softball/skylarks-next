<script lang="ts">
  import ClubInfoSection from "$lib/tib/components/club/ClubInfoSection.svelte";
  import ClubGridItem from "$lib/tib/components/club/ClubGridItem.svelte";
  import type {PageProps} from "../../../../.svelte-kit/types/src/routes";
  import {BookText, Diamond, Pencil, Users} from "lucide-svelte";

  let {data}: PageProps = $props();
</script>

<h1 class="h1">Club</h1>

<!-- Club Info Section -->
<ClubInfoSection club={data.bsmClubData}/>

<!-- Grid Section -->
<div class="grid grid-cols-2 md:grid-cols-3 gap-4 lg:gap-6 mb-4!">
  <ClubGridItem
          description="More Info"
          href="/club/details"
          label="Details"
  >
    {#snippet icon()}
      <BookText/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          description="Adult and Youth"
          href="/teams"
          label="Teams"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          description="Keeping statistics"
          href="/club/scorer"
          label="Scorer"
  >
    {#snippet icon()}
      <Pencil/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          description="Our Referees"
          href="/club/umpire"
          label="Umpire"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          description="The Lark's Nest"
          href="/club/ballpark"
          label="Ballpark"
  >
    {#snippet icon()}
      <Diamond/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          description="Club Administration"
          href="/club/officials"
          label="Officials"
  >
    {#snippet icon()}
      <Users/>
    {/snippet}
  </ClubGridItem>
</div>
