<script lang="ts">
  import {SquarePen} from "lucide-svelte";
  import {invalidate} from "$app/navigation";
  import UniformSetForm from "$lib/dp/components/forms/UniformSetForm.svelte";
  import Cap from "$lib/dp/components/icons/Cap.svelte";
  import Pants from "$lib/dp/components/icons/Pants.svelte";
  import Shirt from "$lib/dp/components/icons/Shirt.svelte";
  import DeleteButton from "$lib/dp/components/utils/DeleteButton.svelte";
  import Dialog from "$lib/dp/components/modal/Dialog.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import type {CustomAuthModel, ExpandedUniformSet} from "$lib/dp/types/ExpandedResponse.ts";

  interface Props {
    uniformSet: ExpandedUniformSet;
  }

  const {uniformSet}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("uniformsets").delete(id);
    invalidate("club:single");
  }
</script>

<article class="card preset-tonal-surface shadow-md">
  <header class="card-header">
    <h3 class="h4 uniform-title">{uniformSet.name}</h3>
  </header>

  <section class="uniform-content">
    <div class="uniform-grid">
      <div class="cap-container">
        <Cap identifier={uniformSet.cap} size={70}/>
      </div>

      <div>
        <Shirt fillColor={uniformSet.jersey} size={64}/>
      </div>

      <div>
        <Pants fillColor={uniformSet.pants} size={64}/>
      </div>
    </div>
  </section>

  <footer class="card-footer">
    {#if uniformSet?.expand?.club?.admins.includes(authRecord.id)}
      <hr class="divider">
      <div class="actions-container">
        <Dialog triggerClasses="btn btn-sm btn-icon preset-tonal-tertiary border border-tertiary-500">

          {#snippet triggerContent()}
            <SquarePen size={16}/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2>Edit Uniform Set "{uniformSet.name}"</h2>
            </header>
          {/snippet}

          <UniformSetForm {uniformSet} clubID={uniformSet.club}/>
        </Dialog>

        <DeleteButton id={uniformSet.id} modelName="Uniform Set" action={deleteAction} iconSize={16}/>
      </div>
    {/if}
  </footer>
</article>

<style>
  .uniform-title {
    font-weight: var(--font-weight-semibold);
  }

  .uniform-content {
    padding: calc(var(--spacing) * 4);
  }

  .uniform-grid {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
    place-items: center;
  }

  .cap-container {
    max-width: calc(var(--spacing) * 20);
  }

  .divider {
    margin-block: calc(var(--spacing) * 2);
  }

  .actions-container {
    display: flex;
    justify-content: flex-end;
    gap: calc(var(--spacing) * 2);
  }
</style>
