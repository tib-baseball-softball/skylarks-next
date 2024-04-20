<script lang="ts">
    import type {MatchBoxscore} from "../../../../../bsm.js/src/model/MatchBoxscore";
    import Linescore from "$lib/components/match/Linescore.svelte";

    export let boxscore: MatchBoxscore
</script>

<section>
    <Linescore linescore="{boxscore.linescore}"/>
</section>

{@html boxscore.full_boxscore_html}