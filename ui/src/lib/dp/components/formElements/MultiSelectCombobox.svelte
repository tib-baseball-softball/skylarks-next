<script lang="ts">
  import { X } from "@lucide/svelte";
  import { toastController } from "$lib/dp/service/ToastController.svelte.ts";
  import type { BaseCollectionResponse } from "$lib/dp/types/pb-types";

  interface Props {
    itemName: string;
    selectedItems: BaseCollectionResponse[];
    allItems: BaseCollectionResponse[];
    labelFunc: (item: BaseCollectionResponse) => string; // TODO: type magic to please compiler
    allowDeletionOfLastItem?: boolean;
  }

  let {
    itemName,
    selectedItems = $bindable(), // MARK: be careful when binding to derived state here that might be undefined
    allItems,
    labelFunc,
    allowDeletionOfLastItem = false,
  }: Props = $props();

  let selectElement: HTMLSelectElement;

  function addItemToSelection(allItems: BaseCollectionResponse[]) {
    if (!selectElement || selectElement.value === "") {
      return;
    }
    const selectedItem = allItems.find(
      (item) => item.id === selectElement?.value,
    );
    const itemExists = selectedItems.find(
      (item) => item.id === selectedItem?.id,
    );

    if (selectedItem && !itemExists) {
      selectedItems.push(selectedItem);
    }
    selectElement.value = "";
  }

  function removeItemFromSelection(itemToRemove: BaseCollectionResponse) {
    if (!allowDeletionOfLastItem && selectedItems.length === 1) {
      toastController.trigger({
        message: `You cannot remove the last ${itemName}!`,
        background: "preset-filled-warning-500",
      });
      return;
    }
    const itemRef = selectedItems.find((entry) => entry.id === itemToRemove.id);

    if (itemRef) {
      const index = selectedItems.indexOf(itemRef);

      if (index !== -1) {
        selectedItems.splice(index, 1);
      }
    }
  }
</script>

{#each selectedItems as selectItem}
  <button
    type="button"
    class="box-button chip preset-filled-primary-500"
    onclick={() => removeItemFromSelection(selectItem)}
  >
    <span>{labelFunc(selectItem)}</span>
    <X size="12" />
  </button>
{/each}

<div>Select to add as {itemName}:</div>
<select
  bind:this={selectElement}
  class="select"
  onchange={() => addItemToSelection(allItems)}
>
  <option selected value="">None</option>
  {#each allItems as item (item.id)}
    <option value={item.id}>{labelFunc(item)}</option>
  {/each}
</select>

<style>
  .box-button {
    margin-inline-end: var(--spacing);

    @media (min-width: 48rem) {
      margin-inline-end: calc(var(--spacing) * 2);
    }
  }
</style>
