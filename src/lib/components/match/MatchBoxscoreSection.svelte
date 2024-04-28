<script lang="ts">
    import type {MatchBoxscore} from "../../../../../bsm.js/src/model/MatchBoxscore";
    import Linescore from "$lib/components/match/Linescore.svelte";

    export let boxscore: MatchBoxscore
</script>

<section>
    <h2 class="h3 my-2">Linescore</h2>
    <Linescore linescore="{boxscore.linescore}"/>
</section>

{@html boxscore.full_boxscore_html}