<script lang="ts">
  import {Tab, TabGroup} from "@skeletonlabs/skeleton";
  import {client} from "../pocketbase";
  import {providerLogin} from "$lib/pocketbase/Auth";
  import {goto} from "$app/navigation";
  import type {AuthProviderInfo} from "pocketbase";

  const {
    authCollection = "users",
    passwordLogin = true,
    signupAllowed = true,
    parent = null,
  } = $props()

  const coll = $derived(client.collection(authCollection))

  const form = $state({
    email: "",
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
      await coll.authWithPassword(form.email, form.password, {expand: "club"})
      goto("/account", {invalidateAll: true})
    }
  }

  async function submitOAuthRequest(provider: AuthProviderInfo) {
    providerLogin(provider, coll)
    goto("/account", {invalidateAll: true})

    //@ts-ignore
    parent.onClose()
  }

  let tabSet = $state(0);
</script>

{#snippet signin()}
    <label class="label">
        <span class="ps-2">Your email</span>
        <input class="input" bind:value={form.email} required type="text" placeholder="name@provider.com"/>
    </label>

    <label class="label">
        <span class="ps-2">Your password</span>
        <input
                bind:value={form.password}
                class="input"
                required
                type="password"
                placeholder="**********"
                minlength="8"
                maxlength="72"
        />
    </label>

    <button class="btn variant-ghost-primary" type="submit" onclick={() => (signup = false)}>
        Login to your account
    </button>
{/snippet}

<div class="w-modal">
    <article class="card p-6">
        <header class="mb-4">
            <h2 class="h4 font-semibold">Sign in to DiamondPlanner</h2>
        </header>

        <form onsubmit={submit}>
            {#if passwordLogin}
                {#if signupAllowed}
                    <TabGroup flex="grow" class="my-2" active="variant-ghost border-b-2 border-surface-900-50-token">
                        <Tab bind:group={tabSet} name="tab1" value={0}>Log In</Tab>
                        <Tab bind:group={tabSet} name="tab2" value={1}>Sign Up</Tab>

                        <svelte:fragment slot="panel">
                            {#if tabSet === 0}

                                {@render signin()}

                            {:else if tabSet === 1}

                                <label class="label">
                                    <span class="ps-2">Your email</span>
                                    <input
                                            class="input"
                                            bind:value={form.email}
                                            required
                                            type="email"
                                            placeholder="name@provider.com"
                                    />
                                </label>

                                <label class="label">
                                    <span class="ps-2">Your password</span>
                                    <input
                                            class="input"
                                            bind:value={form.password}
                                            required
                                            type="password"
                                            placeholder="**********"
                                            minlength="8"
                                            maxlength="72"
                                    />
                                </label>

                                <label class="label">
                                    <span class="ps-2">Confirm password</span>
                                    <input
                                            class="input"
                                            bind:value={form.passwordConfirm}
                                            required
                                            type="password"
                                            placeholder="confirm password"
                                            minlength="8"
                                            maxlength="72"
                                    />
                                </label>

                                <input type="hidden" name="register" value={true}/>

                                <button class="btn variant-ghost-primary" type="submit" onclick={() => (signup = true)}>
                                    Register new account
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
                {#if methods.authProviders.length > 0}

                    <p class="px-2 py-2 text-surface-600-300-token">or sign in with</p>

                    {#each methods.authProviders as provider}
                        <button
                                class="btn variant-ghost-tertiary"
                                type="button"
                                onclick={() => submitOAuthRequest(provider)}
                        >

                            <span class="capitalize">{provider.name}</span>
                        </button>
                    {/each}
                {/if}
            {:catch error}
                <!-- pocketbase not working -->
                <p>There seems to be an error while contacting the backend. Please try again later.</p>
            {/await}
        </form>
    </article>
</div>

<style lang="postcss">
    input, .btn {
        @apply m-2;
    }

    .label {
        @apply my-2;
    }
</style>
