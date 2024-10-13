<script lang="ts">
  import {EditOutline} from "flowbite-svelte-icons";
  import Cap from "$lib/components/icons/Cap.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import Pants from "$lib/components/icons/Pants.svelte";
  import {getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
  import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
  import type {CustomAuthModel, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
  import {authModel} from "$lib/pocketbase/Auth";

  interface Props {
    uniformSet: ExpandedUniformSet
  }

  let {uniformSet}: Props = $props()

  const model = $authModel as CustomAuthModel;

  const modalStore = getModalStore();

  function triggerModal() {
    const modalComponent: ModalComponent = {
      ref: UniformSetForm,
      props: {
        uniformSet: uniformSet,
        clubID: uniformSet.club,
      },
    };

    const modal: ModalSettings = {
      type: "component",
      component: modalComponent,
    };
    modalStore.trigger(modal);
  }
</script>

<article class="card variant-soft-surface shadow-md">
    <header class="card-header">
        <h3 class="h4 font-semibold">{uniformSet.name}</h3>
    </header>

    <section class="p-4">
        <div class="grid grid-cols-3 place-items-center">
            <div class="max-w-20">
                <Cap identifier={uniformSet.cap} classes="w-20"/>
            </div>

            <div>
                <Shirt fillColor={uniformSet.jersey} classes="w-16 h-16"/>
            </div>

            <div>
                <Pants fillColor={uniformSet.pants} classes="w-16 h-16"/>
            </div>
        </div>
    </section>

    {#if uniformSet?.expand?.club?.admins.includes(model.id)}
        <footer class="card-footer flex justify-end">
            <button class="btn variant-ghost-tertiary" onclick={triggerModal}>
                <EditOutline/>
                <span>Edit</span>
            </button>
        </footer>
    {/if}
</article>