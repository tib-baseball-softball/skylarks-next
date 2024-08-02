<script lang="ts">
    import {authModel, client} from "../pocketbase";

    const { parent } = $props()

    async function logout() {
        //@ts-ignore
        parent.onClose()
        client.authStore.clear();
    }
</script>

<div class="w-modal-slim">
    <div class="card p-4">

        <div class="flex mt-1 mb-5 gap-5 items-center">
            <p>Eingeloggt als:</p>
            <div class="badge variant-filled-primary">
                {#if $authModel?.avatar}
                    <img
                        src={client.getFileUrl($authModel, $authModel?.avatar)}
                        alt="profile pic"
                    />
                {/if}
                <samp>{`${$authModel?.first_name ?? ""} ${$authModel?.last_name ?? ""}`}</samp>
            </div>

        </div>
        
        <button class="btn variant-ghost-primary mb-1" onclick={logout}>Abmelden</button>

    </div>
</div>