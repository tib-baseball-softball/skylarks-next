<script lang="ts">
  import {invalidateAll} from "$app/navigation";
  import MultiSelectCombobox from "$lib/dp/components/formElements/MultiSelectCombobox.svelte";
  import {authSettings, client, manualAuthRefresh} from "$lib/dp/client.svelte.ts";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import type {CustomAuthModel, ExpandedTeam} from "$lib/dp/types/ExpandedResponse.ts";
  import type {ClubsResponse, UsersResponse, UsersUpdate} from "$lib/dp/types/pb-types.ts";
  import {closeModal} from "$lib/dp/utility/closeModal.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface Props {
    club: ClubsResponse;
    team: ExpandedTeam;
  }

  const {club, team}: Props = $props();

  let form = $derived.by(() => {
    const ret = $state({
      id: team.id,
    });
    return ret;
  });

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  let selectedUsers: UsersResponse[] = $state([]);

  const allUsersForClub = $derived(client.collection(Collection.Users).getFullList<UsersResponse>({
    filter: `club ?~ '${club.id}' && teams !~ '${team.id}'`, // all users that are club members, but not members of this team
    sort: "+last_name",
  }));

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    try {
      let changedAuthRecord = false;
      for (const user of selectedUsers) {
        if (user.id === authRecord.id) {
          changedAuthRecord = true;
        }
        await client.collection(Collection.Users).update<UsersUpdate>(user.id, {
          "teams+": team.id,
        });
      }
      toastController.triggerGenericFormSuccessMessage("Team members");

      if (changedAuthRecord) {
        await manualAuthRefresh();
      }
    } catch (error) {
      console.error(error);
      toastController.triggerGenericFormErrorMessage("Team members");
    }
    await invalidateAll();
    closeModal();
  }
</script>

<form class="mt-4 space-y-3" onsubmit={submitForm}>
  <div class="edit-form-grid">
    <input
      autocomplete="off"
      bind:value={form.id}
      class="input"
      name="id"
      readonly
      type="hidden"
    />

    <label class="label space-y-3 field-wide">
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
