<script lang="ts">
    import {ClipboardListOutline, ClipboardOutline, InfoCircleOutline} from "flowbite-svelte-icons";
    import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
    import {authModel} from "$lib/pocketbase/Auth";
    import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
    import {client} from "$lib/pocketbase";
    import {invalidate} from "$app/navigation";

    /**
     * Used for Teams in List Teasers and on Club Page.
     */

    interface props {
        team: ExpandedTeam
        link: boolean
    }

    let {team, link = false}: props = $props()

    const model = $authModel as CustomAuthModel;

    function deleteAction(id: string) {
        client.collection("teams").delete(id)
        invalidate("club:single")
    }
</script>

<article class="card variant-glass-surface block shadow-lg" class:card-hover={link}>
    <header class="card-header">
        <h3 class="h4 font-semibold">{team.name}</h3>
    </header>

    <section class="p-4">

        <div class="space-y-2">

            <div class="flex items-center gap-3">
                <ClipboardOutline/>
                <div>
                    <p>{team?.expand?.club.name}</p>
                    <p class="text-sm font-light">Club</p>
                </div>
            </div>

            <div class="flex items-center gap-3">
                <InfoCircleOutline/>
                <div>
                    <p>{team.bsm_league_group}</p>
                    <p class="text-sm font-light">BSM-Liga (für aktuelle Saison)</p>
                </div>
            </div>


            <div class="flex items-center gap-3">
                <ClipboardListOutline/>
                <div>
                    <p>{team.age_group}</p>
                    <p class="text-sm font-light">Altersgruppe</p>
                </div>
            </div>
        </div>
    </section>

    <footer class="card-footer">
        {#if team?.expand?.club?.admins.includes(model.id) || team?.admins.includes(model.id)}
            <hr class="my-2">
            <div class="flex justify-end">
                <DeleteButton id={team.id} modelName="Team" action={deleteAction}/>
            </div>
        {/if}
    </footer>
</article>
