<script lang="ts">
    import { run } from 'svelte/legacy';

    import {browser} from "$app/environment";
    import {preferences} from "$lib/stores";
    import {goto} from "$app/navigation";

    const reloadUponSeasonChange = () => {
        if (browser) {
            let queryString = `?season=${$preferences.selectedSeason}`

            goto(queryString)
        }
    }

    run(() => {
        console.log(`preferences ${$preferences.toString()} changed - reload`)
        reloadUponSeasonChange()
    });
</script>