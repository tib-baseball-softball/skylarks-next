<script lang="ts">
  import type {CustomAuthModel, ExpandedTeam} from "$lib/model/ExpandedResponse";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {invalidate} from "$app/navigation";
  import {ClipboardList, Dumbbell, ShieldHalf} from "lucide-svelte";
  import TeamForm from "$lib/components/forms/TeamForm.svelte";

  /**
   * Used for Teams in List Teasers and on Club Page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  let {team, link = false}: props = $props();

  const model = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("teams").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface block shadow-lg" class:card-hover={link}>
  <a href="/account/team/{team.id}">
    <header class="card-header">
      <h3 class="h4 font-semibold">{team.name}</h3>
    </header>

    <section class="p-4">

      <div class="space-y-2">

        <div class="flex items-center gap-4">
          <ShieldHalf/>
          <div>
            <p>{team?.expand?.club.name}</p>
            <p class="text-sm font-light">Club</p>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <Dumbbell/>
          <div>
            <p>{team.bsm_league_group}</p>
            <p class="text-sm font-light">BSM-Liga (für aktuelle Saison)</p>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <ClipboardList/>
          <dl>
            <dd class="capitalize">{team.age_group}</dd>
            <dt class="text-sm font-light">Age Group</dt>
          </dl>
        </div>
      </div>
    </section>
  </a>

  <footer class="card-footer">
    <hr class="my-2">
    <div class="flex justify-end gap-2">

      {#if team?.expand?.club?.admins.includes(model.id) || team?.admins.includes(model.id)}
        <TeamForm
                club={team.expand?.club}
                team={team}
                buttonClasses="btn btn-icon preset-tonal-tertiary border border-tertiary-500"
                showLabel={false}
        />
      {/if}

      {#if team?.expand?.club?.admins.includes(model.id)}
        <DeleteButton
                id={team.id}
                modelName="Team"
                action={deleteAction}
        />
      {/if}

    </div>
  </footer>
</article>
