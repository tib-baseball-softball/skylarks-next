<script lang="ts">
import type { Match } from "bsm.js"
import { ShareUtility } from "$lib/service/ShareUtility.ts"
import { Share } from "lucide-svelte"

interface Props {
  match: Match
}

let { match }: Props = $props()

const shareUtility = new ShareUtility()
</script>

<button
        class="btn preset-filled"
        onclick={() => shareUtility.shareGame(match, "de-DE")}
        title="Share game data using your browser"
>
  Share Game
  <Share/>
</button>