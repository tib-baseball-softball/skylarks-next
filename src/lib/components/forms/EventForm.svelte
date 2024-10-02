<script lang="ts">
    import { invalidate } from "$app/navigation";
    import type { ExpandedEvent } from "$lib/model/ExpandedResponse";
    import { type UniformsetsResponse } from "$lib/model/pb-types";
    import { client } from "$lib/pocketbase";
    import {
        getDrawerStore,
        RadioGroup,
        RadioItem,
        SlideToggle,
        type ToastSettings,
    } from "@skeletonlabs/skeleton";
    import { getToastStore } from "@skeletonlabs/skeleton";
    import { CloseOutline } from "flowbite-svelte-icons";
    import Flatpickr from "../utility/Flatpickr.svelte";
    import type { Options } from "flatpickr/dist/types/options";

    const toastStore = getToastStore();
    const drawerStore = getDrawerStore();

    const datePickerOptions: Options = {
        enableTime: true,
        dateFormat: "c", // ISO 8601
        altInput: true,
        altFormat: "j F Y - H:i",
        time_24hr: true,
    };

    const toastSettingsSuccess: ToastSettings = {
        message: "Event saved successfully.",
        background: "variant-filled-success",
    };

    const toastSettingsError: ToastSettings = {
        message: "An error occurred while saving the event.",
        background: "variant-filled-error",
    };

    const form: ExpandedEvent = $state(
        $drawerStore.meta.event ?? {
            id: "",
            title: "",
            starttime: "",
            meetingtime: "",
            endtime: "",
            desc: "",
            location: "",
            type: "",
            attire: "",
            cancelled: false,
            bsm_id: "",
            team: $drawerStore.meta?.team?.id,
        },
    );

    const attireOptions = client
        .collection("uniformsets")
        .getFullList<UniformsetsResponse>({
            filter: `club = "${$drawerStore.meta.club}"`,
        });

    async function submitForm(e: SubmitEvent) {
        e.preventDefault();

        let result: ExpandedEvent | null = null;

        try {
            if (form.id) {
                result = await client
                    .collection("events")
                    .update<ExpandedEvent>(form.id, form);
            } else {
                result = await client
                    .collection("events")
                    .create<ExpandedEvent>(form);
            }
        } catch {
            toastStore.trigger(toastSettingsError);
            drawerStore.close();
        }

        if (result) {
            toastStore.trigger(toastSettingsSuccess);
        }
        invalidate("event:list");
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
            <CloseOutline />
        </button>
        <header class="text-xl font-semibold">
            {#if form.id}
                <h2 class="h3">Edit Event "{form?.title}"</h2>
            {:else}
                <h2 class="h3">Create new Event</h2>
            {/if}
        </header>
    </div>

    <form onsubmit={submitForm} class="mt-4 space-y-3">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-2 lg:gap-3 xl:gap-4">
            <input
                name="id"
                autocomplete="off"
                class="input"
                type="hidden"
                readonly
                bind:value={form.id}
            />

            <label class="label">
                Title
                <input
                    name="title"
                    class="input"
                    required
                    type="text"
                    bind:value={form.title}
                />
            </label>

            <label class="label">
                BSM ID
                <input
                    name="bsm_id"
                    class="input"
                    readonly
                    type="text"
                    bind:value={form.bsm_id}
                />
            </label>

            <label class="label">
                Start
                <Flatpickr
                    bind:value={form.starttime}
                    options={datePickerOptions}
                />
            </label>

            <label class="label">
                Meeting
                <Flatpickr
                    bind:value={form.meetingtime}
                    options={datePickerOptions}
                />
            </label>

            <label class="label">
                End
                <Flatpickr
                    bind:value={form.endtime}
                    options={datePickerOptions}
                />
            </label>

            <label class="label col-span-2">
                Description
                <textarea name="desc" class="textarea" bind:value={form.desc}
                ></textarea>
            </label>

            <label class="label col-span-2">
                Location
                <textarea
                    name="location"
                    class="textarea"
                    bind:value={form.location}
                ></textarea>
            </label>

            <label class="label flex flex-col gap-1 col-span-2">
                Type
                <RadioGroup>
                    <RadioItem
                        bind:group={form.type}
                        name="type"
                        value={"game"}
                    >
                        Game
                    </RadioItem>
                    <RadioItem
                        bind:group={form.type}
                        name="type"
                        value={"practice"}
                    >
                        Practice
                    </RadioItem>
                    <RadioItem
                        bind:group={form.type}
                        name="type"
                        value={"misc"}
                    >
                        Other
                    </RadioItem>
                </RadioGroup>
            </label>

            {#await attireOptions then options}
                <label class="label col-span-2">
                    Uniform Set
                    <select class="select" bind:value={form.attire}>
                        {#each options as option}
                            <option value={option.id}>{option.name}</option>
                        {/each}
                    </select>
                </label>
            {/await}

            <SlideToggle
                name="cancelled"
                active={"bg-primary-500"}
                bind:checked={form.cancelled}
            >
                Cancelled
            </SlideToggle>
        </div>

        <hr class="!my-5" />

        <div class="flex justify-center gap-3">
            <button type="submit" class="mt-2 btn variant-ghost-primary">
                Submit
            </button>
        </div>
    </form>
</article>
