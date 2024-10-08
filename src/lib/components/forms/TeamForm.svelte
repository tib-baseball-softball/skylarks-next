<script lang="ts">
    import { invalidate } from "$app/navigation";
    import type { ExpandedTeam } from "$lib/model/ExpandedResponse";
    import { client } from "$lib/pocketbase";
    import {
        getDrawerStore,
        RadioGroup,
        RadioItem,
        type ToastSettings,
    } from "@skeletonlabs/skeleton";
    import { getToastStore } from "@skeletonlabs/skeleton";
    import { CloseOutline } from "flowbite-svelte-icons";

    const toastStore = getToastStore();
    const drawerStore = getDrawerStore();

    const toastSettingsSuccess: ToastSettings = {
        message: "Team data saved successfully.",
        background: "variant-filled-success",
    };

    const toastSettingsError: ToastSettings = {
        message: "An error occurred while saving team data.",
        background: "variant-filled-error",
    };

    const form: ExpandedTeam = $state(
        $drawerStore.meta.team ?? {
            id: "",
            name: "",
            age_group: "",
            club: $drawerStore.meta.club.id, // no binding, cannot be changed via this form
            description: "",
        },
    );

    async function submitForm(e: SubmitEvent) {
        e.preventDefault();

        let result: ExpandedTeam | null = null;

        try {
            if (form.id) {
                result = await client
                    .collection("teams")
                    .update<ExpandedTeam>(form.id, form);
            } else {
                result = await client
                    .collection("teams")
                    .create<ExpandedTeam>(form);
            }
        } catch {
            toastStore.trigger(toastSettingsError);
            drawerStore.close();
        }

        if (result) {
            toastStore.trigger(toastSettingsSuccess);
        }
        invalidate("teams:list");
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
                <h2 class="h3">Edit Team "{form?.name}"</h2>
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
                Name
                <input
                    name="title"
                    class="input"
                    required
                    type="text"
                    bind:value={form.name}
                />
            </label>

            <label class="label">
                Club
                <input
                    name="id"
                    autocomplete="off"
                    class="input"
                    type="text"
                    readonly
                    value={$drawerStore.meta.club?.name}
                />
            </label>

            <label class="label flex flex-col gap-1 col-span-2">
                Type
                <RadioGroup>
                    <RadioItem
                        bind:group={form.age_group}
                        name="age_group"
                        value={"adults"}
                    >
                        Adults
                    </RadioItem>
                    <RadioItem
                        bind:group={form.age_group}
                        name="type"
                        value={"minors"}
                    >
                        Minors
                    </RadioItem>
                </RadioGroup>
            </label>

            <label class="label col-span-2">
                Description
                <textarea
                    name="desc"
                    class="textarea"
                    rows="8"
                    bind:value={form.description}
                ></textarea>
            </label>
        </div>

        <hr class="!my-5" />

        <div class="flex justify-center gap-3">
            <button type="submit" class="mt-2 btn variant-ghost-primary">
                Submit
            </button>
        </div>
    </form>
</article>
