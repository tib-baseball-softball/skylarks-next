<script lang="ts">
  import type {ClubFunction} from "bsm.js";
  import ClubOfficialIcon from "$lib/tib/components/club/officials/ClubOfficialIcon.svelte";

  interface Props {
    clubOfficial: ClubFunction;
  }

  let {clubOfficial}: Props = $props();
</script>

<div class="flex items-center gap-4">
  <ClubOfficialIcon {clubOfficial}/>

  <dl class="space-y-0.5">
    <dt class="font-light text-sm">
      {clubOfficial.function}
    </dt>

    <dd>
      {clubOfficial.person.last_name}, {clubOfficial.person.first_name}
    </dd>
  </dl>
</div>