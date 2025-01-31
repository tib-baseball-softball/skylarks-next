<script lang="ts">
  import {
    BookOutline,
    FireOutline,
    PenOutline,
    UsersGroupOutline,
    UsersGroupSolid,
    UsersOutline,
  } from "flowbite-svelte-icons";
  import ClubInfoSection from "$lib/components/club/ClubInfoSection.svelte";
  import ClubGridItem from "$lib/components/club/ClubGridItem.svelte";
  import type {PageProps} from './$types';

  let {data}: PageProps = $props();
</script>

<!-- Club Info Section -->
<ClubInfoSection club={data.bsmClubData}/>

<!-- Grid Section -->
<div class="grid grid-cols-2 md:grid-cols-3 gap-4 !mb-4">
  <ClubGridItem
          href="/club/details"
          label="Details"
          description="More Info"
  >
    {#snippet icon()}
      <BookOutline size="lg"/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/teams"
          label="Teams"
          description="Adult and Youth"
  >
    {#snippet icon()}
      <UsersGroupSolid size="lg"/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/scorer"
          label="Scorer"
          description="Keeping statistics"
  >
    {#snippet icon()}
      <PenOutline size="lg"/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/umpire"
          label="Umpire"
          description="Our Referees"
  >
    {#snippet icon()}
      <UsersGroupOutline size="lg"/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/ballpark"
          label="Ballpark"
          description="The Lark's Nest"
  >
    {#snippet icon()}
      <FireOutline size="lg"/>
    {/snippet}
  </ClubGridItem>

  <ClubGridItem
          href="/club/officials"
          label="Officials"
          description="Club Administration"
  >
    {#snippet icon()}
      <UsersOutline size="lg"/>
    {/snippet}
  </ClubGridItem>
</div>
