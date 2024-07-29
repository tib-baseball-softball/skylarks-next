<script lang="ts">
  import { onDestroy } from "svelte";
  import { authModel, client } from "../pocketbase";
  import LoginForm from "./LoginForm.svelte";
  import Dialog from "$lib/auth/Dialog.svelte";
  import {Avatar} from "@skeletonlabs/skeleton";
  const { signupAllowed = true } = $props();
  async function logout() {
    client.authStore.clear();
  }

  const unsubscribe = client.authStore.onChange((token, model) => {
    if (model) {
      const { name, username } = model;
      alert(`Signed in as ${name || username || "Admin"}`);
    } else {
      alert(`Signed out`);
    }
  }, false);
  onDestroy(() => {
    unsubscribe();
  });
</script>

{#if $authModel}
  <Dialog>
  
    {#snippet trigger(show)}
      <button class="badge" onclick={show}>
        {#if $authModel.avatar}
          <img
            src={client.getFileUrl($authModel, $authModel.avatar)}
            alt="profile pic"
          />
        {/if}
        
        <Avatar initials={$authModel?.first_name?.charAt(0).toUpperCase() + $authModel?.last_name?.charAt(0).toUpperCase()} background="bg-primary-500" />
       
      </button>
    {/snippet}
    
    <div class="wrapper flex m-6 gap-5 items-center">
    <p>Eingeloggt als:</p>
      <div class="badge variant-filled-primary">
        {#if $authModel.avatar}
          <img
            src={client.getFileUrl($authModel, $authModel.avatar)}
            alt="profile pic"
          />
        {/if}
        <samp>{`${$authModel?.first_name ?? ""} ${$authModel?.last_name ?? ""}`}</samp>
      </div>
      <button class="btn variant-ghost-primary" onclick={logout}>Abmelden</button>
    </div>
    
  </Dialog>
{:else}
  <Dialog>
    {#snippet trigger(show)}
      <button class="btn variant-ghost-primary" onclick={show}>
        {signupAllowed ? "Sign In / Sign Up" : "Sign In"}
      </button>
    {/snippet}
    
    <LoginForm {signupAllowed} />
    
  </Dialog>
{/if}