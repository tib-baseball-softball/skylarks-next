<script lang="ts">
  import {toastController} from "$lib/service/ToastController.svelte";
  import {fly} from 'svelte/transition';
  import {flip} from 'svelte/animate';
</script>

<section class="gui-toast-group">
  {#each toastController.toasts as toast (toast.id)}
    <output
            transition:fly={{ y: 200, duration: 200 }}
            animate:flip={{duration: 400}}
            class="gui-toast {toast.background} rounded-base p-4"
    >
      {toast.message}
    </output>
  {/each}
</section>

<style>
    .gui-toast-group {
        position: fixed;
        z-index: 999999;
        inset-block-end: 0;
        inset-inline: 0;
        padding-block-end: 5vh;
        display: grid;
        justify-items: center;
        justify-content: center;
        gap: 1.5vh;
        pointer-events: none;
    }
</style>