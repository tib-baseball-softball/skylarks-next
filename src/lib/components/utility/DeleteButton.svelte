<script lang="ts">
    import {getModalStore, getToastStore, type ModalSettings, type ToastSettings} from "@skeletonlabs/skeleton";
    import {TrashBinOutline} from "flowbite-svelte-icons";
    import {invalidate} from "$app/navigation";

    const modalStore = getModalStore();
    const toastStore = getToastStore();

    interface Props {
        id: string,
        modelName: string,
        action: (id: string) => void,
        classes?: string,
        buttonText?: string,
    }

    let {id, modelName, action, classes = "btn-sm btn-icon variant-ghost-error", buttonText = ""}: Props = $props()

    function triggerDeleteModal() {
        const modal: ModalSettings = {
            type: 'confirm',
            title: 'Please Confirm',
            body: `Are you sure you wish to delete this ${modelName}?`,
            response: (r: boolean) => {
                if (r) {
                    try {
                        action(id)

                        const toastSettingsDeletions: ToastSettings = {
                            message: `${modelName} deleted successfully.`,
                            background: "variant-filled-success",
                        };

                        toastStore.trigger(toastSettingsDeletions)
                        invalidate("nav:load")
                    } catch {
                        const toastSettingsDeletions: ToastSettings = {
                            message: `Error deleting ${modelName}.`,
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

<button class="btn {classes}" aria-label="delete ${modelName}"
        onclick="{triggerDeleteModal}">
    <TrashBinOutline/>

    {#if buttonText}
        <span>{buttonText}</span>
    {/if}
</button>