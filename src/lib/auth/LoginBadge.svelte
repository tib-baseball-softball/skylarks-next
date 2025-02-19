<script lang="ts">
    import {client} from "../pocketbase/index.svelte";
    import {Avatar, getModalStore, getToastStore, type ModalSettings} from "@skeletonlabs/skeleton";
    import {browser} from "$app/environment";
    import {invalidateAll} from "$app/navigation";
    import {authSettings} from "$lib/pocketbase/index.svelte";
    import {LucideLogIn} from "lucide-svelte";

    const modalStore = getModalStore();
  const toastStore = getToastStore();

  const {signupAllowed = true} = $props();

  const loginModal: ModalSettings = {
    type: "component",
    component: "loginForm",
  };

  const accountOverview: ModalSettings = {
    type: "component",
    component: "accountOverview",
  };

  client.authStore.onChange((_token, model) => {
    // do not do any auth stuff on the server
    if (browser) {
      if (model) {
        const {email} = model;
        toastStore.trigger({
          message: `Signed in as ${email}`,
          background: "variant-filled-success",
        });
        invalidateAll();
      } else {
        toastStore.trigger({message: "Logout successful"});
        invalidateAll();
      }
    }
  });

</script>

{#if authSettings.record}
  <button class="badge" onclick={() => modalStore.trigger(accountOverview)}>

    {#if authSettings.record?.avatar}
      <Avatar
              src={client.files.getURL(authSettings.record, authSettings.record?.avatar)}
              width="w-14"
      />

    {:else}

      <Avatar
              initials={authSettings.record?.first_name?.charAt(0).toUpperCase() +
          authSettings.record?.last_name?.charAt(0)?.toUpperCase()}
              background="variant-filled-primary"
              width="w-14"
              fill="fill-white"
      />

    {/if}
  </button>
{:else}
  <button
          class="btn variant-ghost-primary"
          onclick={() => modalStore.trigger(loginModal)}
          title={signupAllowed ? "Login / Register" : "Login"}
  >
    <LucideLogIn/>
  </button>
{/if}
