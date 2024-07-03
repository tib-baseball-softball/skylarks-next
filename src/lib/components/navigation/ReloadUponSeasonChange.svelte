<script lang="ts">
    import {browser} from "$app/environment";
    import {preferences} from "$lib/stores";
    import {goto} from "$app/navigation";

    const reloadUponSeasonChange = () => {
        if (browser) {
            let queryString = `?season=${$preferences.selectedSeason}`

            goto(queryString)
        }
    }

    $: {
        console.log(`preferences ${$preferences.toString()} changed - reload`)
        reloadUponSeasonChange()
    }
</script>