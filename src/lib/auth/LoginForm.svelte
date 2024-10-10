<script lang="ts">
    import {Tab, TabGroup} from "@skeletonlabs/skeleton";
    import {client} from "../pocketbase";
    import { providerLogin } from "$lib/pocketbase/Auth";
    import { goto } from "$app/navigation";
    import type { AuthProviderInfo } from "pocketbase";

    const {
        authCollection = "users",
        passwordLogin = true,
        signupAllowed = true,
        parent = null,
    } = $props()

    const coll = client.collection(authCollection)

    const form = $state({
        email: "",
        name: "",
        password: "",
        passwordConfirm: "",
    })
    let signup = false

    async function submit(e: SubmitEvent) {
        e.preventDefault()
        //@ts-ignore
        parent.onClose();

        if (signup) {
            await coll.create({...form})
            goto("/signupconfirm")
        } else {
            await coll.authWithPassword(form.email, form.password, { expand: "club" })
            goto("/account")
        }
    }

    async function submitOAuthRequest(provider: AuthProviderInfo) {
      providerLogin(provider, coll)
      goto("/account")

      //@ts-ignore
      parent.onClose()
    }

    let tabSet = $state(0);
</script>

{#snippet signin()}
    <input class="input" bind:value={form.email} required type="text" placeholder="email"/>
    <input
        bind:value={form.password}
        class="input"
        required
        type="password"
        placeholder="password"
        minlength="8"
        maxlength="72"
    />

    <button class="btn variant-ghost-primary" type="submit" onclick={() => (signup = false)}>Sign In</button>
{/snippet}

<div class="w-modal">
    <div class="card p-4">
        <form onsubmit={submit}>
            {#if passwordLogin}
                {#if signupAllowed}
                    <TabGroup>
                        <Tab bind:group={tabSet} name="tab1" value={0}>Login</Tab>
                        <Tab bind:group={tabSet} name="tab2" value={1}>Registrieren</Tab>

                        <svelte:fragment slot="panel">
                            {#if tabSet === 0}

                                {@render signin()}

                            {:else if tabSet === 1}

                                <input
                                    class="input"
                                    bind:value={form.email}
                                    required
                                    type="email"
                                    placeholder="email"
                                />
                                <input
                                    class="input"
                                    bind:value={form.password}
                                    required
                                    type="password"
                                    placeholder="password"
                                    minlength="8"
                                    maxlength="72"
                                />
                                <input
                                    class="input"
                                    bind:value={form.passwordConfirm}
                                    required
                                    type="password"
                                    placeholder="confirm password"
                                    minlength="8"
                                    maxlength="72"
                                />
                                <input
                                    class="input"
                                    bind:value={form.name}
                                    required
                                    type="text"
                                    placeholder="name / label"
                                />
                                <input type="hidden" name="register" value={true}/>
                                <button class="btn variant-ghost-primary" type="submit" onclick={() => (signup = true)}>
                                    Register
                                </button>

                            {/if}
                        </svelte:fragment>

                    </TabGroup>
                {:else}
                    <h2>Login</h2>
                    {@render signin()}
                {/if}
            {/if}

            {#await coll.listAuthMethods({$autoCancel: false}) then methods}

                {#each methods.authProviders as provider}
                    <button class="btn variant-ghost-tertiary" type="button" onclick={() => submitOAuthRequest(provider)}
                    >Sign-in with {provider.name}</button
                    >
                {/each}

            {:catch error}
                <!-- pocketbase not working -->
                <p>There seems to be an error while contacting the backend. Please try again later.</p>
            {/await}
        </form>
    </div>
</div>

<style lang="postcss">
    input, .btn {
        @apply m-2
    }
</style>
