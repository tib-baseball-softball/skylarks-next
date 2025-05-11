<script lang="ts">
  import type {ClubsResponse, UsersResponse, UsersUpdate} from "$lib/model/pb-types.ts";
  import type {ExpandedTeam} from "$lib/model/ExpandedResponse.ts";
  import {client, manualAuthRefresh} from "$lib/pocketbase/index.svelte.ts";
  import MultiSelectCombobox from "$lib/components/utility/MultiSelectCombobox.svelte";
  import {invalidateAll} from "$app/navigation";
  import {closeModal} from "$lib/functions/closeModal.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface Props {
    club: ClubsResponse,
    team: ExpandedTeam,
  }

  let {club, team}: Props = $props();

  let form = $state({
    id: team.id,
  });

  let selectedUsers: UsersResponse[] = $state([]);

  const allUsersForClub = client.collection("users").getFullList<UsersResponse>({
    filter: `club ?~ '${club.id}' && teams !~ '${team.id}'` // all users that are club members, but not members of this team
  });

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    try {
      for (const user of selectedUsers) {
        await client.collection("users").update<UsersUpdate>(user.id, {
          "teams+": team.id
        });
      }
      toastController.trigger({
        message: `All members have been added to team "${team.name}"`,
        background: "preset-filled-success-500",
      });

      // set manually in case we have updated the currently logged-in user
      await manualAuthRefresh();

    } catch (error) {
      console.error(error);
      toastController.trigger({
        message: "An error occurred while saving user team data",
        background: "preset-filled-error-500",
      });
    }
    await invalidateAll();
    closeModal();
  }
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <div class="grid grid-cols-1 md:grid-cols-2 gap-2 md:gap-3 xl:gap-4">
    <input
            autocomplete="off"
            bind:value={form.id}
            class="input"
            name="id"
            readonly
            type="hidden"
    />

    <label class="label space-y-3 md:col-span-2">
      <span>Selected Users:</span><br>

      {#if selectedUsers.length === 0}
        <span>---</span>
      {/if}

      {#await allUsersForClub then users}
        <MultiSelectCombobox
                itemName="Member"
                bind:selectedItems={selectedUsers}
                allItems={users}
                allowDeletionOfLastItem={true}
        />
      {/await}
    </label>
  </div>

  <hr class="my-5!"/>

  <div class="flex justify-center gap-3">
    <button class="mt-2 btn preset-tonal-primary border border-primary-500" type="submit">
      Submit
    </button>
  </div>
</form>