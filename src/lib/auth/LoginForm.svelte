<script lang="ts">
  import {getToastStore, Tab, TabGroup, type ToastSettings} from "@skeletonlabs/skeleton";
  import {client} from "../pocketbase";
  import {goto} from "$app/navigation";
  import OAuthProviderButton from "$lib/auth/OAuthProviderButton.svelte";

  const {
    authCollection = "users",
    passwordLogin = true,
    signupAllowed = true,
    parent = null,
  } = $props()

  const coll = $derived(client.collection(authCollection))
  const toastStore = getToastStore()

  const failSettings: ToastSettings = {
    message: "There was an error processing your authentication request.",
    background: "variant-filled-error"
  }

  const form = $state({
    email: "",
    password: "",
    passwordConfirm: "",
    signupKey: "",
  })
  let signup = false

  async function submit(e: SubmitEvent) {
    e.preventDefault()
    //@ts-ignore
    parent.onClose();

    if (signup) {
      try {
        await coll.create({...form});
        const signupSuccessful = await coll.requestVerification(form.email)

        if (signupSuccessful) {
          goto("/signupconfirm")
        } else {
          toastStore.trigger(failSettings)
        }
      } catch {
        toastStore.trigger(failSettings)
      }

    } else {
      try {
        await coll.authWithPassword(form.email, form.password, {expand: "club"})
        goto("/account", {invalidateAll: true})
      } catch (error) {
        console.error(error)
        toastStore.trigger(failSettings)
      }
    }
  }

  let tabSet = $state(0);
</script>

{#snippet signin()}
    <label class="label">
        <span class="">Your email</span>
        <input class="input" bind:value={form.email} required type="text" placeholder="name@provider.com"/>
    </label>

    <label class="label">
        <span class="">Your password</span>
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

    <button class="btn mt-2 variant-ghost-primary" type="submit" onclick={() => (signup = false)}>
        Login to your account
    </button>
{/snippet}

<div class="w-modal">
    <article class="card p-3 md:p-6">
        <header class="mb-4">
            <h2 class="h4 font-semibold">Sign in to Skylarks Diamond Planner</h2>
        </header>

        <form onsubmit={submit}>
            {#if passwordLogin}
                {#if signupAllowed}
                    <TabGroup flex="grow" class="my-2" active="variant-ghost border-b-2 border-surface-900-50-token">
                        <Tab bind:group={tabSet} name="tab1" value={0}>Log In</Tab>
                        <Tab bind:group={tabSet} name="tab2" value={1}>Create Account</Tab>

                        <svelte:fragment slot="panel">
                            {#if tabSet === 0}

                                {@render signin()}

                            {:else if tabSet === 1}

                                <label class="label">
                                    <span class="">Your email</span>
                                    <input
                                            class="input"
                                            bind:value={form.email}
                                            required
                                            type="email"
                                            placeholder="name@provider.com"
                                    />
                                </label>

                                <div class="grid grid-cols-2 gap-2">


                                    <label class="label">
                                        <span class="">Your password</span>
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
                                        <span class="">Confirm password</span>
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
                                </div>

                                <label class="label">
                                    <span class="">Signup Key</span>
                                    <input
                                            bind:value={form.signupKey}
                                            class="input"
                                            name="signup_key"
                                            placeholder="minimum 8 characters"
                                            minlength="8"
                                            required
                                            type="text"
                                    />
                                    <span class="block text-sm font-light">
                                        A valid signup key needs to be entered upon user account creation.
                                        If you do not have a signup key, please contact your team manager.
                                    </span>
                                </label>

                                <input type="hidden" name="register" value={true}/>

                                <button class="btn variant-ghost-primary my-2" type="submit"
                                        onclick={() => (signup = true)}>
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
                    <hr class="my-2">

                    <div class="mx-2 mt-3 text-surface-600-300-token font-light flex justify-center">
                        <span>or sign in with</span>
                    </div>

                    <div class="my-3 md:my-5 flex flex-wrap gap-6 justify-center">
                        {#each methods.authProviders as provider}
                            <OAuthProviderButton authProvider={provider} collection={coll} parent={parent}/>
                        {/each}
                    </div>

                    <button aria-label="open popover with info about " popovertarget="hint-external" type="button"
                            class="anchor">
                        Note
                    </button>

                    <p id="hint-external" popover="auto"
                       class="mt-1 font-light text-sm variant-filled rounded-token p-2 absolute">
                        Note: it is possible to associate a local account with an
                        external provider by using the same email address. A single account can be associated with more
                        than one external provider.
                    </p>
                {/if}
            {:catch error}
                <!-- pocketbase not working -->
                <p>There seems to be an error while contacting the backend. Please try again later.</p>
            {/await}
        </form>
    </article>
</div>

<style lang="postcss">
    .label {
        @apply my-1;
        @apply md:my-2
    }

    /** TODO: this needs some work */
    :popover-open {
        @apply mx-4;
        @apply md:mx-44;
        @apply md:mt-24;
        @apply lg:w-[640px];
    }
</style>
