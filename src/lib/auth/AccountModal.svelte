<script lang="ts">
    import { authModel, client } from "../pocketbase";

    const { parent } = $props();

    async function logout() {
        //@ts-ignore
        parent.onClose();
        client.authStore.clear();
    }

    async function closeModal() {
        // @ts-ignore
        parent.onClose();
    }
</script>

<div class="w-modal-slim">
    <div class="card p-4">
        <div class="flex mt-1 mb-5 justify-between items-center">
            <p>Eingeloggt als:</p>

            <a
                href="/account"
                class="badge variant-filled-primary"
                onclick={closeModal}
            >
                {#if $authModel?.avatar}
                    <img
                        src={client.getFileUrl($authModel, $authModel?.avatar)}
                        alt="profile pic"
                    />
                {/if}
                <samp
                    >{`${$authModel?.first_name ?? ""} ${$authModel?.last_name ?? ""}`}</samp
                >
            </a>
        </div>

        <div class="flex mt-1 mb-5 justify-between items-start">
            <p>Club:</p>

            <div class="flex flex-wrap gap-2 justify-end">
                {#each $authModel?.expand?.club as club}
                    <span class="badge variant-filled-tertiary"
                        >{club.name}</span
                    >
                {/each}
            </div>
        </div>

        <button class="btn variant-ghost-primary mb-1" onclick={logout}
            >Abmelden</button
        >
    </div>
</div>
