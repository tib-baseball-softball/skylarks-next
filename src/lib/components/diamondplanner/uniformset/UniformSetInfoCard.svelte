<script lang="ts">
  import Cap from "$lib/components/icons/Cap.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import Pants from "$lib/components/icons/Pants.svelte";
  import {getModalStore, type ModalComponent, type ModalSettings} from "@skeletonlabs/skeleton";
  import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
  import type {CustomAuthModel, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {invalidate} from "$app/navigation";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {Edit} from "lucide-svelte";

  interface Props {
    uniformSet: ExpandedUniformSet;
  }

  let {uniformSet}: Props = $props();

  const authRecord = authSettings.record as CustomAuthModel;

  const modalStore = getModalStore();

  function triggerEditModal() {
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

  function deleteAction(id: string) {
    client.collection("uniformsets").delete(id);
    invalidate("club:single");
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

  <footer class="card-footer flex justify-end gap-2">
    {#if uniformSet?.expand?.club?.admins.includes(authRecord.id)}
      <button class="btn btn-sm btn-icon variant-ghost-tertiary" aria-label="edit uniform set"
              onclick={triggerEditModal}>
        <Edit/>
      </button>

      <DeleteButton id={uniformSet.id} modelName="Uniform Set" action={deleteAction}/>
    {/if}
  </footer>
</article>