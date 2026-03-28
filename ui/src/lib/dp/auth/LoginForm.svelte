<script lang="ts">
  import {fade, slide} from "svelte/transition";
  import {goto} from "$app/navigation";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  import OAuthProviderButton from "$lib/dp/auth/OAuthProviderButton.svelte";
  import PasswordRequestButton from "$lib/dp/auth/PasswordRequestButton.svelte";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {Extension} from "$lib/dp/types/ExpandedResponse.js";
  import type {UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import type {Toast} from "$lib/dp/types/Toast.ts";
  import {client} from "../client.svelte.js";
  //@ts-ignore
  import {Tabs} from "bits-ui";
  import {page} from "$app/state";

  const {authCollection = "users", passwordLogin = true, signupAllowed = true} = $props();

  const coll = $derived(client.collection(authCollection));

  const failSettings: Toast = {
    message: "There was an error processing your authentication request.",
    background: "preset-filled-error-500",
  };

  const prefilledSignupKey = page.url.searchParams.get("signup_key") ?? "";

  let form: Extension<Partial<UsersUpdate>, { signup_key: string }> = $state({
    email: "",
    password: "",
    first_name: "",
    last_name: "",
    passwordConfirm: "",
    signup_key: prefilledSignupKey,
  });
  let signup = false;

  async function submit(e: SubmitEvent) {
    e.preventDefault();

    if (signup) {
      try {
        await coll.create({...form});
        const signupSuccessful = await coll.requestVerification(form.email ?? "");

        if (signupSuccessful) {
          await goto("/signupconfirm");
        } else {
          toastController.trigger(failSettings);
        }
      } catch {
        toastController.trigger(failSettings);
      }
    } else {
      try {
        const authResponse = await coll.authWithPassword(form.email ?? "", form.password ?? "", {
          expand: "club",
        });

        if (authResponse) {
          await goto("/account", {invalidateAll: true});
        }
      } catch (error) {
        console.error(error);
        toastController.trigger(failSettings);
      }
    }
  }

  let tabSet: "login" | "signup" | string = $state(page.url.searchParams.get("action") === "signup" ? "signup" : "login");
  let forgotPassword = $state(false);
</script>

{#snippet login()}
  <label class="label">
    <span>Your email</span>
    <input
      class="input"
      bind:value={form.email}
      required
      type="email"
      placeholder="name@provider.com"
      autocomplete="email"
    />
  </label>

  {#if !forgotPassword}
    <label class="label" transition:slide>
      <span>Your password</span>
      <input
        bind:value={form.password}
        class="input"
        required
        type="password"
        placeholder="**********"
        minlength="8"
        maxlength="72"
        autocomplete="current-password"
      />
    </label>

    <button
      transition:slide
      class="btn preset-tonal-primary login-btn"
      type="submit"
      onclick={() => (signup = false)}
      disabled={form.email === "" || form.password === ""}
    >
      Login to your account
    </button>
  {/if}

  <div class="support-row">
    <Switch
      name="forgot password"
      onCheckedChange={(e) => (forgotPassword = e.checked)}
    >
      Forgot Password?
    </Switch>

    {#if forgotPassword}
      <div class="forgot-cta" transition:fade>
        <PasswordRequestButton
          email={form.email ?? ""}
          disabled={form.email === ""}
          classes="btn anchor p-0"
        >
          Reset Password
        </PasswordRequestButton>
      </div>
    {/if}
  </div>
{/snippet}

<div class="outer-wrapper">
  <div class="inner-wrapper">
    <article>
      <header>
        <h1 class="h2">Sign in to Diamond Planner</h1>
      </header>

      <form onsubmit={submit}>
        {#if passwordLogin}
          {#if signupAllowed}
            <Tabs.Root
              bind:value={tabSet}
              class="tabs-wrap"
            >
              <Tabs.List
                class="tabs-list preset-tonal-surface"
              >
                <Tabs.Trigger
                  value="login"
                  class="tabs-trigger btn"
                >Log In
                </Tabs.Trigger>
                <Tabs.Trigger
                  value="signup"
                  class="tabs-trigger btn"
                >Create Account
                </Tabs.Trigger>
              </Tabs.List>

              <Tabs.Content value="login" class="tab-panel">
                {@render login()}
              </Tabs.Content>

              <Tabs.Content value="signup" class="tab-panel">

                {#if tabSet === "signup"}
                  <label class="label">
                    <span class="">Your email</span>
                    <input
                      class="input"
                      bind:value={form.email}
                      required
                      type="email"
                      placeholder="name@provider.com"
                      autocomplete="email"
                    />
                  </label>

                  <div class="form-grid">

                    <label class="label">
                      <span class="">First Name</span>
                      <input
                        class="input"
                        bind:value={form.first_name}
                        placeholder="John"
                        required
                        type="text"
                        autocomplete="given-name"
                      />
                    </label>

                    <label class="label">
                      <span class="">Last Name</span>
                      <input
                        class="input"
                        bind:value={form.last_name}
                        placeholder="Doe"
                        required
                        type="text"
                        autocomplete="family-name"
                      />
                    </label>
                  </div>

                  <div class="form-grid">

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
                        autocomplete="new-password"
                      />
                    </label>

                    <label class="label">
                      <span class="">Confirm password</span>
                      <input
                        class="input"
                        bind:value={form.passwordConfirm}
                        required
                        type="password"
                        placeholder="**********"
                        minlength="8"
                        maxlength="72"
                        autocomplete="new-password"
                      />
                    </label>
                  </div>

                  <label class="label">
                    <span class="">Signup Key</span>
                    <input
                      bind:value={form.signup_key}
                      class="input"
                      name="signup_key"
                      placeholder="**********"
                      minlength="8"
                      required
                      type="text"
                      autocomplete="one-time-code"
                    />
                    <span class="help-block">
                                          A valid signup key needs to be entered upon user account creation.
                                          If you do not have a signup key, please contact your team manager.
                                      </span>
                  </label>

                  <input type="hidden" name="register" value={true}/>

                  <button
                    class="btn preset-tonal-primary register-btn"
                    type="submit"
                    onclick={() => (signup = true)}
                    disabled={form.email === "" || form.password === "" || form.passwordConfirm === "" || form.signup_key === ""}
                  >
                    Register new account
                  </button>
                {/if}
              </Tabs.Content>

            </Tabs.Root>

          {/if}
        {/if}

        {#await coll.listAuthMethods({requestKey: null}) then methods}
          {#if methods.oauth2.providers.length > 0 && forgotPassword === false}
            <hr class="divider-sm">

            <div class="muted-row">
              <span>or sign in with</span>
            </div>

            <div class="providers">
              {#each methods.oauth2.providers as provider}

                {#if tabSet === "login"}
                  <!--Login Buttons - no signupKey-->
                  <OAuthProviderButton authProvider={provider} collection={coll} disabled={false}/>
                {:else}
                  <!--Signup Buttons - with signupKey state prop-->
                  <OAuthProviderButton authProvider={provider} collection={coll}
                                       signup_key={form.signup_key} disabled={form.signup_key === ""}/>
                {/if}

              {/each}
            </div>

            {#if tabSet === "signup" && form.signup_key === ""}
              <div class="muted-row" transition:slide>
              <span>
                Even when using an external login provider,
                a signup key (see above) is still required for account creation.
              </span>
              </div>
            {/if}
            <hr class="divider-lg">

            <p class="note">
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

</div>

<style>
  .outer-wrapper {
    display: flex;
    justify-content: center;
  }

  .inner-wrapper {
    display: flex;
    justify-content: center;
    margin-block-start: calc(var(--spacing) * 3) !important;
    margin-block-end: calc(var(--spacing) * 10) !important;

    @media (min-width: 48rem) {
      width: 60%;
    }

    @media (min-width: 64rem) {
      width: 55%;
    }
  }

  .label {
    margin-block: calc(var(--spacing) * 1);
    @media (width >= 48rem /* 768px */
    ) {
      margin-block: calc(var(--spacing) * 2); /* 0.5rem = 8px */
    }
  }

  header {
    margin-bottom: 2em;
  }

  .login-btn {
    border: 1px solid var(--color-primary-500);
    display: block;
    margin-top: calc(var(--spacing) * 4);
  }

  .support-row {
    display: block;
    margin-block: calc(var(--spacing) * 6);
  }

  @media (min-width: 48rem) {
    .support-row {
      display: flex;
      gap: calc(var(--spacing) * 6);
      align-items: center;
      justify-content: flex-start;
    }
  }

  .forgot-cta {
    margin-top: calc(var(--spacing) * 2);
  }

  @media (min-width: 48rem) {
    .forgot-cta {
      margin-top: 0;
    }
  }

  .form-grid {
    display: grid;
    gap: calc(var(--spacing) * 2);
    grid-template-columns: 1fr;
  }

  @media (min-width: 40rem) {
    .form-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .help-block {
    display: block;
    font-size: var(--text-sm);
    font-weight: var(--font-weight-light);
  }

  .register-btn {
    border: 1px solid var(--color-primary-500);
    margin-block: calc(var(--spacing) * 2);
  }

  .divider-sm {
    margin-block: calc(var(--spacing) * 2);
  }

  .divider-lg {
    margin-block: calc(var(--spacing) * 6);
  }

  .muted-row {
    margin-inline: calc(var(--spacing) * 2);
    margin-top: calc(var(--spacing) * 3);
    display: flex;
    justify-content: center;
    font-weight: var(--font-weight-light);
    color: var(--color-surface-700);
  }

  @media (prefers-color-scheme: dark) {
    .muted-row {
      color: var(--color-surface-300);
    }
  }

  .providers {
    margin-block: calc(var(--spacing) * 3);
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 6);
    justify-content: center;
  }

  @media (min-width: 48rem) {
    .providers {
      margin-block: calc(var(--spacing) * 5);
    }
  }

  .note {
    font-weight: var(--font-weight-light);
    font-size: var(--text-sm);
    margin-top: calc(var(--spacing) * 10);
  }
</style>
