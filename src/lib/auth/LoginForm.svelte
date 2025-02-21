<script lang="ts">
  import {getToastStore, SlideToggle, Tab, TabGroup, type ToastSettings} from "@skeletonlabs/skeleton";
  import {client} from "../pocketbase/index.svelte";
  import {goto} from "$app/navigation";
  import OAuthProviderButton from "$lib/auth/OAuthProviderButton.svelte";
  import PasswordRequestButton from "$lib/auth/PasswordRequestButton.svelte";
  import {fade, slide} from "svelte/transition";

  const {
    authCollection = "users",
    passwordLogin = true,
    signupAllowed = true,
  } = $props();

  const coll = $derived(client.collection(authCollection));
  const toastStore = getToastStore();

  const failSettings: ToastSettings = {
    message: "There was an error processing your authentication request.",
    background: "variant-filled-error"
  };

  const form = $state({
    email: "",
    password: "",
    passwordConfirm: "",
    signup_key: "",
  });
  let signup = false;

  async function submit(e: SubmitEvent) {
    e.preventDefault();

    if (signup) {
      try {
        await coll.create({...form});
        const signupSuccessful = await coll.requestVerification(form.email);

        if (signupSuccessful) {
          await goto("/signupconfirm");
        } else {
          toastStore.trigger(failSettings);
        }
      } catch {
        toastStore.trigger(failSettings);
      }

    } else {
      try {
        const authResponse = await coll.authWithPassword(form.email, form.password, {expand: "club"});

        if (authResponse) {
          await goto("/account", {invalidateAll: true});
        }
      } catch (error) {
        console.error(error);
        toastStore.trigger(failSettings);
      }
    }
  }

  let tabSet = $state(0);
  let forgotPassword = $state(false);
</script>

{#snippet signin()}
  <label class="label">
    <span class="">Your email</span>
    <input class="input" bind:value={form.email} required type="email" placeholder="name@provider.com"/>
  </label>

  {#if !forgotPassword}
    <label class="label" transition:slide>
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

    <button
            transition:slide
            class="btn variant-ghost-primary block mt-4"
            type="submit"
            onclick={() => (signup = false)}
            disabled={form.email === "" || form.password === ""}
    >
      Login to your account
    </button>
  {/if}

  <div class="block md:flex gap-6 items-center justify-start my-6">
    <SlideToggle size="sm" name="forgot password" bind:checked={forgotPassword}>Forgot Password?</SlideToggle>

    {#if forgotPassword}
      <div class="mt-2 md:mt-0" transition:fade>
        <PasswordRequestButton
                email={form.email}
                disabled={form.email === ""}
                classes="btn anchor p-0"
        >
          Reset Password
        </PasswordRequestButton>
      </div>
    {/if}
  </div>
{/snippet}

<div class="flex justify-center !mt-3 !mb-10">
  <article>
    <header class="!mb-6">
      <h1 class="h2">Sign in to Skylarks Diamond Planner</h1>
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
                          bind:value={form.signup_key}
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

                <button
                        class="btn variant-ghost-primary my-2"
                        type="submit"
                        onclick={() => (signup = true)}
                        disabled={form.email === "" || form.password === "" || form.passwordConfirm === "" || form.signup_key === ""}
                >
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

      {#await coll.listAuthMethods({requestKey: null}) then methods}
        {#if methods.oauth2.providers.length > 0}
          <hr class="my-2">

          <div class="mx-2 mt-3 text-surface-600-300-token font-light flex justify-center">
            <span>or sign in with</span>
          </div>

          <div class="my-3 md:my-5 flex flex-wrap gap-6 justify-center">
            {#each methods.oauth2.providers as provider}

              {#if tabSet === 0}
                <!--Login Buttons - no signupKey-->
                <OAuthProviderButton authProvider={provider} collection={coll} disabled={false}/>
              {:else}
                <!--Signup Buttons - with signupKey state prop-->
                <OAuthProviderButton authProvider={provider} collection={coll}
                                     signup_key={form.signup_key} disabled={form.signup_key === ""}/>
              {/if}

            {/each}
          </div>

          {#if tabSet === 1 && form.signup_key === ""}
            <div class="mx-2 mt-3 text-surface-600-300-token font-light flex justify-center" transition:slide>
              <span>
                Even when using an external login provider,
                a signup key (see above) still required for account creation.
              </span>
            </div>
          {/if}
          <hr class="my-6">

          <p class="font-light text-sm mt-10">
            Note: it is possible to associate a local account with an
            external provider later by using the same email address. A single account can be associated with more
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
        @apply md:my-2;
    }
</style>
