<script lang="ts">
import LoginBadge from "$lib/auth/LoginBadge.svelte"
import LocalePicker from "$lib/components/meta/LocalePicker.svelte"
</script>

<LocalePicker/>

<LoginBadge signupAllowed={true}/>