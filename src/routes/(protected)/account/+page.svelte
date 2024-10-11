<script lang="ts">
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import UserDataCard from "$lib/components/diamondplanner/user/UserDataCard.svelte";
    import type { CustomAuthModel } from "$lib/model/ExpandedResponse.js";
    import { authModel } from "$lib/pocketbase/Auth";
    import {
        PlusOutline,
        InfoCircleOutline,
        ShieldOutline,
        TagOutline,
        ArrowLeftToBracketOutline,
    } from "flowbite-svelte-icons";

    const model = $authModel as CustomAuthModel;

    let { data } = $props();
</script>

<h1 class="h1 lg:mt-4">My Dashboard</h1>

<div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <UserDataCard {model} />

    {#each model?.expand?.club as club}
        <div class="card variant-glass-primary shadow-lg">
            <header class="card-header">
                <h2 class="h4 font-semibold">Club</h2>
            </header>

            <section class="p-4 space-y-2">
                <div class="flex items-center gap-3">
                    <ShieldOutline />
                    <div>
                        <p>{club.name}</p>
                        <p class="text-sm font-light">Name</p>
                    </div>
                </div>

                <div class="flex items-center gap-3">
                    <TagOutline />
                    <div>
                        <p>{club.acronym}</p>
                        <p class="text-sm font-light">Acronym</p>
                    </div>
                </div>

                <div class="flex items-center gap-3">
                    <InfoCircleOutline />
                    <div>
                        <p>{club.bsm_id}</p>
                        <p class="text-sm font-light">BSM-ID</p>
                    </div>
                </div>
            </section>
        </div>
    {/each}

    {#if !model?.expand?.club}
        <div class="card variant-glass-primary shadow-lg">
            <header class="card-header">
                <h2 class="h4 font-semibold">Club</h2>
            </header>

            <section class="p-4 space-y-2">
                <div class="flex items-center gap-3">
                    <p>You are not a member of any clubs yet.</p>
                </div>
            </section>

            <footer class="card-footer">
                <div class="flex items-center gap-3">
                    <button class="btn variant-ghost-primary">
                        <PlusOutline />
                        <span>Create a Club</span>
                    </button>
                    <button class="btn variant-ghost-primary">
                        <ArrowLeftToBracketOutline />
                        <span>Join a Club</span>
                    </button>
                </div>
            </footer>
        </div>
    {/if}
</div>

<h2 class="h2 mt-3">My Teams</h2>

{#await data.teams then teams}
    <div class="grid grid-cols-1 md:grid-cols-2 gap-3 mb-3">
        {#each teams as team}
            <a href="/account/team/{team.id}">
                <TeamTeaserCard {team} link={true} />
            </a>
        {/each}

        {#if teams?.length === 0}
            <p>You are not a member of any teams yet.</p>
        {/if}
    </div>
{/await}
