<script lang="ts">
  import LoginBadge from "$lib/auth/LoginBadge.svelte";
  import {env} from "$env/dynamic/public";
  import LocalePicker from "$lib/components/meta/LocalePicker.svelte";
</script>

{#if env.PUBLIC_APPLICATION_CONTEXT !== "Production"}
  <div class="preset-tonal-warning p-1.5 rounded-container border border-warning-500 text-sm">
    {env.PUBLIC_APPLICATION_CONTEXT}
  </div>
{/if}

<LocalePicker/>

{#if env.PUBLIC_AUTH_FUNCS_ENABLED === "true"}
  <LoginBadge signupAllowed={true}/>
{/if}