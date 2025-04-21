<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {Avatar} from "@skeletonlabs/skeleton-svelte";
  import {client} from "$lib/pocketbase/index.svelte";
  import UserDetailsForm from "$lib/auth/UserDetailsForm.svelte";
  import PasswordRequestButton from "$lib/auth/PasswordRequestButton.svelte";
  import {Lock, Mail, User} from "lucide-svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";
  import ChangeEmailForm from "$lib/auth/ChangeEmailForm.svelte";

  interface props {
    model: CustomAuthModel;
  }

  let {model}: props = $props();
</script>

<div class="card preset-tonal-surface lg:col-span-2 shadow-lg">
  <header class="card-header">
    <h2 class="h4 font-semibold">User Data</h2>
  </header>

  <section class="p-4 flex flex-col sm:flex-row gap-4 lg:gap-12 sm:items-center">
    <Avatar src={client.files.getURL(model, model?.avatar)}/>

    <div class="grid grid-cols-1 gap-2 col-span-4">
      <div class="flex items-center gap-3">
        <User/>
        <div>
          <p>{`${model?.first_name ?? ""} ${model?.last_name ?? ""}`}</p>
          <p class="text-sm font-light">Name</p>
        </div>
      </div>

      <div class="flex items-center gap-3">
        <Mail/>
        <div>
          <p>{model?.email}</p>
          <p class="text-sm font-light">email address</p>
        </div>
      </div>
    </div>
  </section>

  <footer class="card-footer flex justify-end items-center gap-3 md:gap-4">
    <p>Edit User Data:</p>
    <PasswordRequestButton email={model.email} classes="btn preset-tonal-secondary border border-secondary-500">
      <Lock/>
    </PasswordRequestButton>

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
  </footer>
</div>