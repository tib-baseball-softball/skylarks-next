<script lang="ts">
  import type {ClubTeam} from "bsm.js";

  interface Props {
    clubTeam: ClubTeam
  }
  let {clubTeam}: Props = $props()
</script>

{clubTeam.team.name}
