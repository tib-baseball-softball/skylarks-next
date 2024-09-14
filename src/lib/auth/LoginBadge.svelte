<script lang="ts">
    import { onDestroy } from "svelte";
    import { client } from "../pocketbase";
    import { Avatar, type ModalSettings } from "@skeletonlabs/skeleton";
    import { getToastStore } from "@skeletonlabs/skeleton";
    import { getModalStore } from "@skeletonlabs/skeleton";
    import { browser } from "$app/environment";
    import { invalidateAll } from "$app/navigation";
    import { authModel } from "$lib/pocketbase/Auth";

    const modalStore = getModalStore();
    const toastStore = getToastStore();

    const { signupAllowed = true } = $props();

    const loginModal: ModalSettings = {
        type: "component",
        component: "loginForm",
    };

    const accountOverview: ModalSettings = {
        type: "component",
        component: "accountOverview",
    };

    async function logout() {
        client.authStore.clear();
    }

    const unsubscribe = client.authStore.onChange((token, model) => {
        // do not do any auth stuff on the server
        if (browser) {
            if (model) {
                const { email } = model;
                toastStore.trigger({
                    message: `Angemeldet als ${email}`,
                    background: "variant-filled-success",
                });
                invalidateAll();
            } else {
                toastStore.trigger({ message: "Abmeldung erfolgreich" });
                invalidateAll();
            }
        }
    }, false);

    onDestroy(() => {
        unsubscribe();
    });
</script>

{#if $authModel}
    <button class="badge" onclick={() => modalStore.trigger(accountOverview)}>
        {#if $authModel.avatar}
            <img
                src={client.getFileUrl($authModel, $authModel.avatar)}
                alt="profile pic"
            />
        {/if}

        <Avatar
            initials={$authModel?.first_name?.charAt(0).toUpperCase() +
                $authModel?.last_name?.charAt(0).toUpperCase()}
            background="variant-filled-primary"
            width="w-14"
            fill="fill-white"
        />
    </button>
{:else}
    <button
        class="btn variant-ghost-primary"
        onclick={() => modalStore.trigger(loginModal)}
    >
        {signupAllowed ? "Login / Registrieren" : "Login"}
    </button>
{/if}
