<script lang="ts">
  import type {UniformsetsResponse} from "$lib/model/pb-types";
  import {EditOutline} from "flowbite-svelte-icons";
  import Cap from "$lib/components/icons/Cap.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import Pants from "$lib/components/icons/Pants.svelte";

  interface Props {
    uniformSet: UniformsetsResponse
  }

  let {uniformSet}: Props = $props()
</script>

<article class="card variant-soft-surface shadow-md">
    <header class="card-header">
        <h3 class="h4 font-semibold">{uniformSet.name}</h3>
    </header>
    <section class="p-4">
        <div class="grid grid-cols-3 place-items-center">
            <div class="max-w-20">
                <Cap identifier={uniformSet.cap} classes="w-20"/>
            </div>

            <div>
                <Shirt fillColor={uniformSet.jersey} classes="w-16 h-16"/>
            </div>

            <div>
                <Pants fillColor={uniformSet.pants} classes="w-16 h-16"/>
            </div>
        </div>
    </section>
    <footer class="card-footer flex justify-end">
        <button class="btn variant-ghost-tertiary">
            <EditOutline/>
            <span>Edit</span>
        </button>
    </footer>
</article>