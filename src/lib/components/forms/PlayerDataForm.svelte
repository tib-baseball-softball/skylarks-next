<script lang="ts">
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import {
    getDrawerStore,
    getToastStore,
    InputChip,
    RadioGroup,
    RadioItem,
    type ToastSettings,
  } from "@skeletonlabs/skeleton";
  import {CloseOutline} from "flowbite-svelte-icons";
  import {
    getAllBaseballPositionStringValues,
    positionEnumStringValuesToKeys,
    positionKeysToEnumStringValues,
  } from "$lib/types/BaseballPosition";

  const toastStore = getToastStore();
  const drawerStore = getDrawerStore();

  const toastSettingsSuccess: ToastSettings = {
    message: "Player data saved successfully.",
    background: "variant-filled-success",
  };

  const toastSettingsError: ToastSettings = {
    message: "An error occurred while saving player data.",
    background: "variant-filled-error",
  };

  const model = authSettings.record as CustomAuthModel;
  const form = $state({
    id: model.id ?? "",
    number: model.number ?? "",
    bats: model.bats ?? "",
    position: [""],
    throws: model.throws ?? "",
    image: model.image ?? "",
    umpire: model.umpire ?? "0",
    scorer: model.scorer ?? "0",
    bsm_id: model.bsm_id ?? 0,
  });

  const possiblePositionValues = getAllBaseballPositionStringValues();
  let selectedPositions: string[] = $state(
      positionKeysToEnumStringValues(model.position),
  );
  let inputChip: InputChip;

  async function submitForm(e: SubmitEvent) {
    e.preventDefault();

    form.position = positionEnumStringValuesToKeys(selectedPositions);

    let result: CustomAuthModel | null = null;

    try {
      result = await client
          .collection("users")
          .update<CustomAuthModel>(form.id, form);
    } catch {
      toastStore.trigger(toastSettingsError);
      drawerStore.close();
    }

    if (result) {
      toastStore.trigger(toastSettingsSuccess);
    }
    drawerStore.close();
  }
</script>

<article class="p-6">
    <div class="flex items-center gap-5">
        <button
                aria-label="cancel and close"
                class="btn variant-ghost-surface"
                onclick={drawerStore.close}
        >
            <CloseOutline/>
        </button>
        <header class="text-xl font-semibold">
            <h2 class="h3">
                Edit Player Data for "{`${model.first_name} ${model.last_name}`}"
            </h2>
        </header>
    </div>

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
                <InputChip
                        bind:this={inputChip}
                        bind:value={selectedPositions}
                        name="positions"
                        placeholder="Enter positions or click buttons..."
                        whitelist={possiblePositionValues}
                        allowUpperCase={true}
                        chips="variant-filled-primary"
                />
                <div class="flex flex-wrap gap-2 md:col-span-2">
                    {#each possiblePositionValues as value}
                        <button
                                type="button"
                                class="btn btn-sm variant-ringed-primary"
                                onclick={() => inputChip.addChip(value)}
                        >
                            {value}
                        </button>
                    {/each}
                </div>
            </label>

            <label class="label flex flex-col gap-1 md:col-span-2">
                Bats
                <RadioGroup>
                    <RadioItem bind:group={form.bats} name="bats" value={"1"}>
                        Left
                    </RadioItem>
                    <RadioItem bind:group={form.bats} name="bats" value={"2"}>
                        Right
                    </RadioItem>
                    <RadioItem bind:group={form.bats} name="bats" value={"3"}>
                        Switch
                    </RadioItem>
                </RadioGroup>
            </label>

            <label class="label flex flex-col gap-1 md:col-span-2">
                Throws
                <RadioGroup>
                    <RadioItem
                            bind:group={form.throws}
                            name="throws"
                            value={"1"}
                    >
                        Left
                    </RadioItem>
                    <RadioItem
                            bind:group={form.throws}
                            name="throws"
                            value={"2"}
                    >
                        Right
                    </RadioItem>
                    <RadioItem
                            bind:group={form.throws}
                            name="throws"
                            value={"3"}
                    >
                        Switch
                    </RadioItem>
                </RadioGroup>
            </label>

            <label class="label flex flex-col gap-1 md:col-span-2">
                Umpire License
                <RadioGroup>
                    <RadioItem
                            bind:group={form.umpire}
                            name="umpire"
                            value={"0"}
                    >
                        None
                    </RadioItem>
                    <RadioItem
                            bind:group={form.umpire}
                            name="umpire"
                            value={"1"}
                    >
                        A
                    </RadioItem>
                    <RadioItem
                            bind:group={form.umpire}
                            name="umpire"
                            value={"2"}
                    >
                        B
                    </RadioItem>
                    <RadioItem
                            bind:group={form.umpire}
                            name="umpire"
                            value={"3"}
                    >
                        C
                    </RadioItem>
                    <RadioItem
                            bind:group={form.umpire}
                            name="umpire"
                            value={"4"}
                    >
                        D
                    </RadioItem>
                </RadioGroup>
            </label>

            <label class="label flex flex-col gap-1 md:col-span-2">
                Scorer License
                <RadioGroup>
                    <RadioItem
                            bind:group={form.scorer}
                            name="scorer"
                            value={"0"}
                    >
                        None
                    </RadioItem>
                    <RadioItem
                            bind:group={form.scorer}
                            name="scorer"
                            value={"1"}
                    >
                        A
                    </RadioItem>
                    <RadioItem
                            bind:group={form.scorer}
                            name="scorer"
                            value={"2"}
                    >
                        B
                    </RadioItem>
                    <RadioItem
                            bind:group={form.scorer}
                            name="scorer"
                            value={"3"}
                    >
                        C
                    </RadioItem>
                </RadioGroup>
            </label>
        </div>

        <hr class="!my-5"/>

        <div class="flex justify-center gap-3">
            <button type="submit" class="mt-2 btn variant-ghost-primary">
                Submit
            </button>
        </div>
    </form>
</article>
