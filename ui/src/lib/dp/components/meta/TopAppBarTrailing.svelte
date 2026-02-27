<script lang="ts">
  import LocalePicker from "$lib/dp/components/meta/LocalePicker.svelte";
  import LoginBadge from "$lib/dp/auth/LoginBadge.svelte";
</script>

<div class="root">
  <LocalePicker/>

  <LoginBadge signupAllowed={true}/>
</div>

<style>
  .root {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 2);
  }
</style>
