<script lang="ts">
    import {getModalStore, getToastStore, type ModalSettings, type ToastSettings} from "@skeletonlabs/skeleton";
    import {TrashBinOutline} from "flowbite-svelte-icons";

    const modalStore = getModalStore();
    const toastStore = getToastStore();

    interface Props {
        id: string,
        modelName: string,
        action: (id: string) => void
    }

    let {id, modelName, action}: Props = $props()

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

<button class="btn btn-sm btn-icon variant-ghost-error" aria-label="delete ${modelName}"
        onclick="{triggerDeleteModal}">
    <TrashBinOutline/>
</button>