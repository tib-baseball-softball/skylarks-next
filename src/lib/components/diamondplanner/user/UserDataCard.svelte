<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {Avatar, getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
  import {EnvelopeOutline, UserEditOutline, UserOutline} from "flowbite-svelte-icons";
  import {client} from "$lib/pocketbase/index.svelte";
  import UserDetailsForm from "$lib/auth/UserDetailsForm.svelte";

  interface props {
    model: CustomAuthModel
  }

  let {model}: props = $props()

  const modalStore = getModalStore();

  function triggerModal() {
    const modalComponent: ModalComponent = {
      ref: UserDetailsForm,
      props: {},
    };

    const modal: ModalSettings = {
      type: "component",
      component: modalComponent,
    };
    modalStore.trigger(modal);
  }
</script>

<div class="card variant-glass-surface lg:col-span-2 shadow-lg">
    <header class="card-header">
        <h2 class="h4 font-semibold">User Data</h2>
    </header>

    <section class="p-4 flex flex-col sm:flex-row gap-4 lg:gap-12 sm:items-center">
        <Avatar src={client.files.getURL(model, model?.avatar)}/>

        <div class="grid grid-cols-1 gap-2 col-span-4">
            <div class="flex items-center gap-3">
                <UserOutline size="lg"/>
                <div>
                    <p>{`${model?.first_name ?? ""} ${model?.last_name ?? ""}`}</p>
                    <p class="text-sm font-light">Name</p>
                </div>
            </div>

            <div class="flex items-center gap-3">
                <EnvelopeOutline size="lg"/>
                <div>
                    <p>{model?.email}</p>
                    <p class="text-sm font-light">email address</p>
                </div>
            </div>
        </div>
    </section>

    <footer class="card-footer flex justify-end">
        <button aria-label="edit user data" class="btn variant-ghost-primary" onclick="{() => triggerModal()}">
            <UserEditOutline/>
        </button>
    </footer>
</div>