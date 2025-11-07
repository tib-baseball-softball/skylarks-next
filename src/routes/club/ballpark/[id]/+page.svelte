<script lang="ts">
import type { PageProps } from "./$types"
import BallparkDetailCard from "$lib/components/club/ballpark/BallparkDetailCard.svelte"
import BallparkTrafficCard from "$lib/components/club/ballpark/BallparkTrafficCard.svelte"

let { data }: PageProps = $props()
</script>

<h1 class="h1">Ballpark Details</h1>

<div class="grid grid-cols-1 md:grid-cols-2 gap-4 lg:gap-6 xl:gap-8">
  <img class="rounded-base" loading="lazy" src="{data.field.photo_url}"
       alt="a baseball or softball field: {data.field.name}">

  <BallparkDetailCard field={data.field}/>

  <div>
    <h2 class="h3 mb-2">How to get there</h2>
    <BallparkTrafficCard field={data.field}/>
  </div>
</div>
