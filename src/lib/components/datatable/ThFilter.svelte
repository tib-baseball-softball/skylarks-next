<script lang="ts">
    import type { DataHandler } from '@vincjo/datatables';
    interface Props {
        handler: DataHandler;
        filterBy: string;
    }

    let { handler, filterBy }: Props = $props();
    let value: string = $state("");
</script>

<th>
    <input
            class="input text-sm w-full preset-outlined-primary-500"
            type="text"
            bind:value
            oninput={() => {
			if (filterBy) handler.filter(value, filterBy);
		}}
    />
</th>