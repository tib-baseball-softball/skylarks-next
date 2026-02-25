<script lang="ts">
  import { File, IdCard, ShieldHalf, Users } from "lucide-svelte";
  import type { ExpandedTeam } from "$lib/dp/types/ExpandedResponse.ts";

  /**
   * Used on Team Detail page.
   */

  interface props {
    team: ExpandedTeam;
    link: boolean;
  }

  let { team, link = false }: props = $props();
</script>

<article class="card card-content preset-tonal-surface" class:card-hover={link}>
  <div class="detail-row">
    <File />
    <div>
      <p>{team.name}</p>
      <p class="detail-label">Name</p>
    </div>
  </div>

  <hr class="divider" />

  <div class="detail-row">
    <IdCard />
    <div>
      <p>{team.bsm_league_group}</p>
      <p class="detail-label">BSM-Liga (für aktuelle Saison)</p>
    </div>
  </div>

  <hr class="divider" />

  <div class="detail-row">
    <ShieldHalf />
    <div>
      <p>{team?.expand?.club.name}</p>
      <p class="detail-label">Club</p>
    </div>
  </div>

  <hr class="divider" />

  <div class="detail-row">
    <Users />
    <dl>
      <dd class="team-age-group">{team.age_group}</dd>
      <dt class="detail-label">Age Group</dt>
    </dl>
  </div>
</article>

<style>
  .card-content {
    display: block;
    padding: calc(var(--spacing) * 3);
  }

  .detail-row {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 3);
    align-self: flex-end;
  }

  .detail-label {
    font-size: var(--text-sm);
    line-height: var(--tw-leading, var(--text-sm--line-height));
    font-weight: var(--font-weight-light);
  }

  .divider {
    margin-block: calc(var(--spacing) * 2);
  }

  .team-age-group {
    text-transform: capitalize;
  }
</style>
