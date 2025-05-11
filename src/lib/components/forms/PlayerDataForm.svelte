<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {Segment, TagsInput} from "@skeletonlabs/skeleton-svelte";
  import {
    getAllBaseballPositionStringValues,
    positionEnumStringValuesToKeys,
    positionKeysToEnumStringValues,
  } from "$lib/types/BaseballPosition";
  import {Edit} from "lucide-svelte";
  //@ts-ignore
  import * as Sheet from "$lib/components/utility/sheet/index.js";
  import type {Toast} from "$lib/types/Toast.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";

  interface ValidateArgs {
    inputValue: string;
    value: string[];
  }

  interface Props {
    buttonClasses?: string;
  }

  let { buttonClasses = "" }: Props = $props();

  const toastSettingsSuccess: Toast = {
    message: "Player data saved successfully.",
    background: "preset-filled-success-500",
  };

  const toastSettingsError: Toast = {
    message: "An error occurred while saving player data.",
    background: "preset-filled-error-500",
  };

  // Mark: here we do not need to be reactive as we're editing
  const authRecord = authSettings.record as CustomAuthModel;
  const form = $state({
    id: authRecord.id ?? "",
    number: authRecord.number ?? "",
    bats: authRecord.bats ?? "",
    position: [""],
    throws: authRecord.throws ?? "",
    image: authRecord.image ?? "",
    umpire: authRecord.umpire ?? "0",
    scorer: authRecord.scorer ?? "0",
    bsm_id: authRecord.bsm_id ?? 0,
  });

  let open = $state(false)

  const possiblePositionValues = getAllBaseballPositionStringValues();
  let selectedPositions: string[] = $state(
      positionKeysToEnumStringValues(authRecord.position),
  );

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
      })
    }
  }

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    form.position = positionEnumStringValuesToKeys(selectedPositions);

    let result: CustomAuthModel | null = null;

    try {
      result = await client
          .collection("users")
          .update<CustomAuthModel>(form.id, form);
    } catch {
      toastController.trigger(toastSettingsError);
    }

    if (result) {
      toastController.trigger(toastSettingsSuccess);
      open = false;
    }
  }
</script>

<Sheet.Root bind:open={open}>
  <Sheet.Trigger class={buttonClasses}>
    <Edit/>
    <span>Edit Player Data</span>
  </Sheet.Trigger>

  <Sheet.Content>
    <Sheet.Header></Sheet.Header>

    <header class="text-xl font-semibold">
      <h2 class="h3">
        Edit Player Data for "{`${authRecord.first_name} ${authRecord.last_name}`}"
      </h2>
    </header>

  <form onsubmit={submitForm} class="mt-4 space-y-3">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-3 xl:gap-4">
      <input
              name="id"
              autocomplete="off"
              class="input"
              type="hidden"
              readonly
              bind:value={form.id}
      />

      <label class="label">
        Jersey Number
        <input
                name="number"
                autocomplete="off"
                class="input"
                type="number"
                bind:value={form.number}
        />
      </label>

      <label class="label">
        BSM ID (Player Pass Number)
        <input
                name="bsm_id"
                autocomplete="off"
                class="input"
                type="text"
                readonly
                bind:value={form.bsm_id}
        />
      </label>

      <label class="label flex flex-col gap-1 md:gap-2 md:col-span-2">
        Positions
        <TagsInput
                value={selectedPositions}
                name="positions"
                onValueChange={(e) => (selectedPositions = e.value)}
                placeholder="Enter positions or click buttons..."
                validate={validatePositionValue}
                tagBackground="preset-filled-primary-500"
        />
        <span class="flex flex-wrap gap-2 md:col-span-2">
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

      <label class="label flex flex-col gap-1 md:col-span-2">
        Bats
        <Segment name="bats" value={form.bats} onValueChange={(e) => (form.bats = e.value ?? "1")}>
          <Segment.Item value={"1"} classes="flex-grow">
            Left
          </Segment.Item>
          <Segment.Item value={"2"} classes="flex-grow">
            Right
          </Segment.Item>
          <Segment.Item value={"3"} classes="flex-grow">
            Switch
          </Segment.Item>
        </Segment>
      </label>

      <label class="label flex flex-col gap-1 md:col-span-2">
        Throws
        <Segment name="throws" value={form.throws} onValueChange={(e) => (form.throws = e.value ?? "1")}>
          <Segment.Item value={"1"} classes="flex-grow">
            Left
          </Segment.Item>
          <Segment.Item value={"2"} classes="flex-grow">
            Right
          </Segment.Item>
          <Segment.Item value={"3"} classes="flex-grow">
            Switch
          </Segment.Item>
        </Segment>
      </label>

      <label class="label flex flex-col gap-1 md:col-span-2">
        Umpire License
        <Segment name="umpire" value={form.umpire} onValueChange={(e) => (form.umpire = e.value ?? "0")}>
          <Segment.Item value={"0"} classes="flex-grow">
            None
          </Segment.Item>
          <Segment.Item value={"1"} classes="flex-grow">
            A
          </Segment.Item>
          <Segment.Item value={"2"} classes="flex-grow">
            B
          </Segment.Item>
          <Segment.Item value={"3"} classes="flex-grow">
            C
          </Segment.Item>
          <Segment.Item value={"4"} classes="flex-grow">
            D
          </Segment.Item>
        </Segment>
      </label>

      <label class="label flex flex-col gap-1 md:col-span-2">
        Scorer License
        <Segment name="scorer" value={form.scorer} onValueChange={(e) => (form.scorer = e.value ?? "0")}>
          <Segment.Item value={"0"} classes="flex-grow">
            None
          </Segment.Item>
          <Segment.Item value={"1"} classes="flex-grow">
            A
          </Segment.Item>
          <Segment.Item value={"2"} classes="flex-grow">
            B
          </Segment.Item>
          <Segment.Item value={"3"} classes="flex-grow">
            C
          </Segment.Item>
        </Segment>
      </label>
    </div>

    <hr class="my-5!"/>

    <div class="flex justify-center gap-3">
      <button type="submit" class="mt-2 btn preset-tonal-primary border border-primary-500">
        Submit
      </button>
    </div>
  </form>
  </Sheet.Content>
</Sheet.Root>
