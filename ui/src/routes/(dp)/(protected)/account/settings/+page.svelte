<script lang="ts">
  import {authSettings} from "$lib/dp/client.svelte.js";
  import PasswordRequestButton from "$lib/dp/auth/PasswordRequestButton.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {Lock, Mail, User} from "lucide-svelte";
  import ChangeEmailForm from "$lib/dp/auth/ChangeEmailForm.svelte";
  import UserDetailsForm from "$lib/dp/components/forms/UserDetailsForm.svelte";
  import PushSettingsSection from "$lib/dp/components/settings/PushSettingsSection.svelte";
  import ICalSection from "$lib/dp/components/settings/ICalSection.svelte";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse";

  const model = $derived(authSettings.record) as CustomAuthModel;
</script>

<h1 class="h1">User Settings</h1>
{#if model}
  <section class="card preset-tonal-surface">
    <header class="card-header">
      <h2 class="h3">
        Account Data
      </h2>
    </header>

    <div class="card-body">

      <label class="label account-label">
        <span>
        Reset Password
        </span>
        <PasswordRequestButton classes="btn preset-tonal-secondary border border-secondary-500" email={model.email}>
          <Lock/>
        </PasswordRequestButton>
      </label>

      <label class="label account-label">
        <span>
        Change Email Address
        </span>
        <Dialog triggerClasses="btn preset-tonal-secondary border border-secondary-500">


          {#snippet triggerContent()}
            <Mail/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2>
                Change email address for {model.first_name}
                {model.last_name}
              </h2>
            </header>
          {/snippet}

          <ChangeEmailForm/>
        </Dialog>
      </label>

      <label class="label account-label">
        <span>
        Edit Profile Data
        </span>
        <Dialog triggerClasses="btn preset-tonal-primary border border-primary-500">
          {#snippet triggerContent()}
            <User/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2>
                Edit User Data for {model.first_name}
                {model.last_name}
              </h2>
            </header>
          {/snippet}

          <UserDetailsForm/>
        </Dialog>
      </label>
    </div>

    <footer class="card-footer">
      <p>
        Your player profile data can be changed on
        <a class="anchor" href="/account/playerprofile" title="Go to Player profile page">a dedicated page
        </a>.
      </p>
    </footer>
  </section>

  <section class="card preset-tonal-surface">
    <PushSettingsSection userID={model.id}/>
  </section>

  <section class="card preset-tonal-surface">
    <ICalSection link={model?.ical_link ?? ""}>
      {#snippet header()}
        <span>Calendar Import</span>
      {/snippet}

      {#snippet subheader()}
        <h3 class="h4">Your personal calendar link</h3>
        <p class="cal-hint">
          This link includes all events for all teams that you are a member of,
          going back one year.
        </p>
      {/snippet}
    </ICalSection>
  </section>

{:else}
  <p>User not found. That should not happen, if it does anyway, please log in.</p>
{/if}

<style>
  section {
    margin-block-end: calc(var(--spacing) * 6);
  }

  .account-label {
    display: grid;
    justify-items: start;
    align-items: center;
    gap: calc(var(--spacing) * 2);
    margin-block-end: calc(var(--spacing) * 2);
    grid-template-columns: 200px 1fr;
  }

  .cal-hint {
    margin-block-start: var(--spacing);
  }
</style>
