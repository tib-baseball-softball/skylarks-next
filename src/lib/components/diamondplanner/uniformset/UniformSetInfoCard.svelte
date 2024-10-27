<script lang="ts">
    import {EditOutline, TrashBinOutline} from "flowbite-svelte-icons";
    import Cap from "$lib/components/icons/Cap.svelte";
    import Shirt from "$lib/components/icons/Shirt.svelte";
    import Pants from "$lib/components/icons/Pants.svelte";
    import {
        getModalStore,
        getToastStore,
        type ModalComponent,
        type ModalSettings,
        type ToastSettings
    } from "@skeletonlabs/skeleton";
    import UniformSetForm from "$lib/components/forms/UniformSetForm.svelte";
    import type {CustomAuthModel, ExpandedUniformSet} from "$lib/model/ExpandedResponse";
    import {authModel} from "$lib/pocketbase/Auth";
    import {client} from "$lib/pocketbase";
    import {invalidate} from "$app/navigation";

    interface Props {
        uniformSet: ExpandedUniformSet
    }

    let {uniformSet}: Props = $props()

    const model = $authModel as CustomAuthModel;

    const modalStore = getModalStore();
    const toastStore = getToastStore();

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

    function triggerDeleteModal() {
        const modal: ModalSettings = {
            type: 'confirm',
            title: 'Please Confirm',
            body: 'Are you sure you wish to delete this uniform set?',
            response: (r: boolean) => {
                if (r) {
                    try {
                        client.collection("uniformsets").delete(uniformSet.id)
                        invalidate("club:single")

                        const toastSettingsDeletions: ToastSettings = {
                            message: "Uniform Set deleted successfully.",
                            background: "variant-filled-success",
                        };

                        toastStore.trigger(toastSettingsDeletions)
                    } catch {
                        const toastSettingsDeletions: ToastSettings = {
                            message: "Error deleting Uniform Set.",
                            background: "variant-filled-error",
                        };

                        toastStore.trigger(toastSettingsDeletions)
                    }
                }
            },
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

    <footer class="card-footer flex justify-end gap-2">
        {#if uniformSet?.expand?.club?.admins.includes(model.id)}
            <button class="btn btn-sm btn-icon variant-ghost-tertiary" aria-label="edit uniform set"
                    onclick={triggerEditModal}>
                <EditOutline/>
            </button>
            <button class="btn btn-sm btn-icon variant-ghost-error" aria-label="delete uniform set"
                    onclick="{triggerDeleteModal}">
                <TrashBinOutline/>
            </button>
        {/if}
    </footer>
</article>