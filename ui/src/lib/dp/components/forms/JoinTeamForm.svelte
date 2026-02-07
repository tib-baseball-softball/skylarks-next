<script lang="ts">
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {client, manualAuthRefresh} from "$lib/dp/client.svelte.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {invalidateAll} from "$app/navigation";

  type JoinTeamPayload = {
    signupKey: string;
    userID: string;
  }

  interface Props {
    authRecord: CustomAuthModel;
    teamID: string;
  }

  let {authRecord, teamID}: Props = $props();


  let form = $state({
    signupKey: ""
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    const payload: JoinTeamPayload = {
      signupKey: form.signupKey,
      userID: authRecord.id
    };

    try {
      const response = await client.send(`/api/dp/teams/${teamID}/join`, {
        method: "POST",
        body: payload
      });

      if (response) {
        toastController.trigger({
          message: "Successfully joined team",
          background: "preset-filled-success-500"
        });
      }
      await manualAuthRefresh();
      await invalidateAll();
    } catch (error) {
      console.error(error);
      toastController.trigger({
        message: "An error occurred while joining the team.",
        background: "preset-filled-error-500",
      });
    }
  }
</script>

<form onsubmit={submitForm}>
  <label class="label">
    <span>Signup Key</span>
    <input
      autocomplete="one-time-code"
      bind:value={form.signupKey}
      class="input"
      placeholder="*********"
      required
      type="password"
    />
  </label>
  <button class="btn preset-tonal-primary" type="submit">Submit</button>
</form>

<style>
  button {
    border: 1px solid var(--color-primary-500);
  }

  input {
    margin-block-end: calc(var(--spacing) * 4);
  }
</style>
