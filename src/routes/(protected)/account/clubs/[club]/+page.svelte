<script lang="ts">
    import ClubDetailCard from "$lib/components/diamondplanner/club/ClubDetailCard.svelte";
    import TeamTeaserCard from "$lib/components/diamondplanner/team/TeamTeaserCard.svelte";
    import UniformSetInfoCard from "$lib/components/diamondplanner/uniformset/UniformSetInfoCard.svelte";
    import {PlusOutline} from "flowbite-svelte-icons";
    import type {CustomAuthModel, ExpandedClub, ExpandedTeam, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
    import {authModel} from "$lib/pocketbase/Auth";
    import {getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
    import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";

    let {data} = $props()

    const model = $authModel as CustomAuthModel;

    let club: ExpandedClub = $derived(data.club)
    let teams: ExpandedTeam[] = $derived(data.teams)
    let uniformSets: ExpandedUniformSet[] = $derived(data.uniformSets)

    const modalStore = getModalStore();

    function triggerModal() {
        const modalComponent: ModalComponent = {
            ref: UniformSetForm,
            props: {
                uniformSet: null,
                clubID: club.id,
            },
        };

        const modal: ModalSettings = {
            type: "component",
            component: modalComponent,
        };
        modalStore.trigger(modal);
    }
</script>

<h1 class="h1">{club.name}</h1>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    <ClubDetailCard {club}/>
</section>

<h2 class="h2">Club Teams</h2>

<section class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
    {#each teams as team}
        <a href="/account/team/{team.id}">
            <TeamTeaserCard {team} link={true}/>
        </a>
    {/each}
</section>

<section>
    <header>
        <h2 class="h2 mb-3">Uniform Sets</h2>
    </header>
    <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-3 mb-3">
        {#each uniformSets as uniformSet}
            <UniformSetInfoCard {uniformSet}/>
        {/each}
    </div>

    {#if club?.admins.includes(model.id)}
        <button class="btn variant-ghost-primary" onclick={triggerModal}>
            <PlusOutline/>
            <span>Create new</span>
        </button>
    {/if}
</section>