<script lang="ts">
import LicenseLevelIndicator from "$lib/components/club/licenses/LicenseLevelIndicator.svelte"
import type { License } from "bsm.js"

interface Props {
  license: License
}

let { license }: Props = $props()
</script>

<a href="/club/license/{license.id}" class="container">
  <LicenseLevelIndicator level={license.level}/>
  <span>{license.person.last_name}, {license.person.first_name}</span>
</a>

<style lang="postcss">
    .container {
        display: flex;
        align-items: center;
        gap: 1rem;
    }
</style>
