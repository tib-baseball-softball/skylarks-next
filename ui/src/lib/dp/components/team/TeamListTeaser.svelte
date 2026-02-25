<script lang="ts">
  import {ClipboardList, Dumbbell, ShieldHalf} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import TeamForm from "$lib/dp/components/forms/TeamForm.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";

  /**
   * Used for Teams in List Teasers and on Club Page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  const {team, link = false}: props = $props();

  const model = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("teams").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface shadow-lg" class:card-hover={link}>
  <a href="/account/team/{team.id}">
    <header class="card-header">
      <h3 class="h4 font-semibold">{team.name}</h3>
    </header>

    <section class="team-content">

      <div class="team-details">

        <div class="item-row">
          <ShieldHalf/>
          <div>
            <dd>{team?.expand?.club?.name}</dd>
            <dt class="detail-label">Club</dt>
          </div>
        </div>

        <div class="item-row">
          <Dumbbell/>
          <div>
            <dd>{team.bsm_league_group}</dd>
            <dt class="detail-label">BSM-Liga (für aktuelle Saison)</dt>
          </div>
        </div>

        <div class="item-row">
          <ClipboardList/>
          <dl>
            <dd class="team-age-group">{team.age_group}</dd>
            <dt class="detail-label">Age Group</dt>
          </dl>
        </div>
      </div>
    </section>
  </a>

  <footer class="card-footer">
    <hr class="footer-divider">
    <div class="footer-actions">

      {#if team?.expand?.club?.admins.includes(model.id) || team?.admins.includes(model.id)}
        <TeamForm
                club={team.expand?.club}
                team={team}
                triggerVariant="tonal-tertiary"
                triggerSize="sm"
                triggerIcon={true}
                showLabel={false}
        />
      {/if}

      {#if team?.expand?.club?.admins.includes(model.id)}
        <DeleteButton
                id={team.id}
                modelName="Team"
                action={deleteAction}
                iconSize={16}
        />
      {/if}

    </div>
  </footer>
</article>

<style>
    .card {
        display: block;
    }

    .team-content {
        padding: calc(var(--spacing) * 4);
    }

    .team-details {
        display: flex;
        flex-direction: column;
        gap: calc(var(--spacing) * 2);
    }

    .item-row {
        display: flex;
        align-items: center;
        gap: calc(var(--spacing) * 4);
    }

    .detail-label {
        font-size: var(--text-sm);
        line-height: var(--tw-leading, var(--text-sm--line-height));
        font-weight: var(--font-weight-light);
    }

    .footer-divider {
        margin-block: calc(var(--spacing) * 2);
    }

    .footer-actions {
        display: flex;
        justify-content: flex-end;
        gap: calc(var(--spacing) * 2);
    }
    
    .team-age-group {
      text-transform: capitalize;
    }
</style>
