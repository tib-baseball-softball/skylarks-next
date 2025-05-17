<script lang="ts">
  import Cap from "$lib/components/icons/Cap.svelte";
  import Shirt from "$lib/components/icons/Shirt.svelte";
  import Pants from "$lib/components/icons/Pants.svelte";
  import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
  import type {CustomAuthModel, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {invalidate} from "$app/navigation";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {Edit} from "lucide-svelte";
  import Dialog from "$lib/components/utility/Dialog.svelte";

  interface Props {
    uniformSet: ExpandedUniformSet;
  }

  let {uniformSet}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("uniformsets").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface shadow-md">
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

  <footer class="card-footer">
    {#if uniformSet?.expand?.club?.admins.includes(authRecord.id)}
      <hr class="my-2">
      <div class="flex justify-end gap-2">
        <Dialog triggerClasses="btn btn-sm btn-icon preset-tonal-tertiary border border-tertiary-500">

          {#snippet triggerContent()}
            <Edit/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2>Edit Uniform Set "{uniformSet.name}"</h2>
            </header>
          {/snippet}

          <UniformSetForm {uniformSet} clubID={uniformSet.club}/>
        </Dialog>

        <DeleteButton id={uniformSet.id} modelName="Uniform Set" action={deleteAction}/>
      </div>
    {/if}
  </footer>
</article>