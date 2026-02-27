<script lang="ts">
  import {SquarePen} from "lucide-svelte";
  import TabsRadioGroup from "$lib/dp/components/formElements/TabsRadioGroup.svelte";
  import Switch from "$lib/dp/components/formElements/Switch.svelte";
  //@ts-ignore
  import * as Sheet from "$lib/dp/components/modal/sheet";
  import TagsInput from "$lib/dp/components/formElements/TagsInput.svelte";
  import {authSettings, client} from "$lib/dp/client.svelte.js";
  import {toastController} from "$lib/dp/service/ToastController.svelte.ts";
  import {
    getAllBaseballPositionStringValues,
    positionEnumStringValuesToKeys,
    positionKeysToEnumStringValues,
  } from "$lib/dp/types/BaseballPosition.ts";
  import type {CustomAuthModel} from "$lib/dp/types/ExpandedResponse.ts";
  import {Collection} from "$lib/dp/enum/Collection.ts";

  interface ValidateArgs {
    inputValue: string;
    value: string[];
  }

  interface Props {
    triggerVariant?: "filled-primary" | "tonal-primary" | "tonal-secondary" | "tonal-tertiary" | "tonal-surface";
    triggerSize?: "default" | "sm";
    triggerIcon?: boolean;
    triggerSpaced?: boolean;
  }

  const {
    triggerVariant = "tonal-primary",
    triggerSize = "default",
    triggerIcon = false,
    triggerSpaced = false,
  }: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  function formFromProps(record: CustomAuthModel) {
    return {
      id: record.id ?? "",
      number: record.number ?? "",
      bats: record.bats ?? "",
      position: [""],
      throws: record.throws ?? "",
      image: record.image ?? "",
      umpire: record.umpire ?? "0",
      scorer: record.scorer ?? "0",
      bsm_id: record.bsm_id ?? 0,
      display_on_website: record.display_on_website ?? false,
    };
  }

  let form = $derived.by(() => {
    const created = $state(formFromProps(authRecord));
    return created;
  });

  let open = $state(false);

  const possiblePositionValues = getAllBaseballPositionStringValues();
  // svelte-ignore state_referenced_locally - this component uses manual state management
  let selectedPositions: string[] = $state(positionKeysToEnumStringValues(authRecord.position));

  function validatePositionValue(details: ValidateArgs): boolean {
    return possiblePositionValues.includes(details.inputValue);
  }

  function addPosition(value: string) {
    if (!selectedPositions.includes(value)) {
      selectedPositions.push(value);
    } else {
      toastController.trigger({
        message: "Position is already selected.",
        background: "preset-filled-error-500",
      });
    }
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    form.position = positionEnumStringValuesToKeys(selectedPositions);

    let result: CustomAuthModel | null = null;

    try {
      result = await client.collection(Collection.Users).update<CustomAuthModel>(form.id, form);
    } catch {
      toastController.triggerGenericFormErrorMessage("Player data");
    }

    if (result) {
      toastController.triggerGenericFormSuccessMessage("Player data");
      open = false;
    }
  }

  const onValueChange = (e: { value: string[] }) => (selectedPositions = e.value);
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger
    class={[
      "btn",
      "trigger-button",
      `trigger-variant-${triggerVariant}`,
      triggerSize === "sm" && "btn-sm",
      triggerIcon && "btn-icon",
      triggerSpaced && "trigger-spaced",
      triggerVariant === "filled-primary" && "preset-filled-primary-500",
      triggerVariant === "tonal-primary" && "preset-tonal-primary",
      triggerVariant === "tonal-secondary" && "preset-tonal-secondary",
      triggerVariant === "tonal-tertiary" && "preset-tonal-tertiary",
      triggerVariant === "tonal-surface" && "preset-tonal-surface",
    ]}
  >
    <SquarePen/>
    <span>Edit Player Data</span>
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header>
      <h2 class="h3">
        Edit Player Data for "{`${authRecord.first_name} ${authRecord.last_name}`}"
      </h2>
    </header>

    <form onsubmit={submitForm}>
      <div class="edit-form-grid">
        <input
          autocomplete="off"
          bind:value={form.id}
          class="input"
          name="id"
          readonly
          type="hidden"
        />

        <label class="label">
          Jersey Number
          <input
            autocomplete="off"
            bind:value={form.number}
            class="input"
            name="number"
            type="number"
          />
        </label>

        <label class="label">
          BSM ID (Player Pass Number)
          <input
            autocomplete="off"
            bind:value={form.bsm_id}
            class="input"
            name="bsm_id"
            readonly
            type="text"
          />
        </label>

        <span class="label-wide">
          <Switch
            bind:checked={form.display_on_website}
            name="display_on_website"
          >
            Display this data publicly on team website
          </Switch>
        </span>

        <label class="label label-wide">
          Positions
          <TagsInput
            name="positions"
            onValueChange={onValueChange}
            placeholder="Enter positions or click buttons..."
            validate={validatePositionValue}
            value={selectedPositions}
          />
          <span class="label-wide position-labels">
          {#each possiblePositionValues as value}
            <button
              type="button"
              class="btn btn-sm preset-outlined-primary-500"
              onclick={() => addPosition(value)}
            >
              {value}
            </button>
          {/each}
        </span>
        </label>

        <TabsRadioGroup
          bind:value={form.bats}
          label="Bats"
          name="bats"
          options={["left", "right", "switch"]}
        />

        <TabsRadioGroup
          bind:value={form.throws}
          label="Throws"
          name="throws"
          options={["left", "right", "switch"]}
        />

        <TabsRadioGroup
          bind:value={form.umpire}
          label="Umpire License"
          name="umpire"
          options={["none", "A", "B", "C", "D"]}
        />

        <TabsRadioGroup
          bind:value={form.scorer}
          label="Scorer License"
          name="scorer"
          options={["none", "A", "B", "C", "D"]}
        />
      </div>
      
      <hr>

      <div class="submit-container">
        <button class="btn preset-tonal-primary border border-primary-500" type="submit">
          Submit
        </button>
      </div>
    </form>
  </Sheet.Content>
</Sheet.Root>

<style>
  form {
    margin-block: calc(var(--spacing) * 3);
  }
  
  .label-wide {
    grid-column: span 2 / span 2;
    gap: calc(var(--spacing) * 2);
  }
  
  .position-labels {
    display: flex;
    flex-wrap: wrap;
    gap: calc(var(--spacing) * 2);
  }
  
  .submit-container {
    display: flex;
    justify-content: center;
    gap: calc(var(--spacing) * 3);
    
    > button {
      margin-block-start: calc(var(--spacing) * 2);
    }
  }
  
  hr {
    margin-block: calc(var(--spacing) * 4);
  }
</style>
