<script lang="ts">
  import type {DataHandler} from "@vincjo/datatables";
  import ThSort from "./ThSort.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {Avatar} from "@skeletonlabs/skeleton";
  import {client} from "$lib/pocketbase";
  import {CheckCircleSolid, EditOutline, LockOpenOutline,} from "flowbite-svelte-icons";
  import LocalDate from "../utility/LocalDate.svelte";

  interface Props {
    handler: DataHandler<CustomAuthModel>;
    showAdminSection?: boolean;
  }

  let {handler, showAdminSection = false}: Props = $props();

  const rows = $derived(handler.getRows());
</script>

<thead>
<tr class="sticky">
    <ThSort {handler} orderBy="last_name">Name</ThSort>
    <ThSort {handler} orderBy="verified">Status</ThSort>
    <ThSort {handler} orderBy="number">Number</ThSort>
    <ThSort {handler} orderBy="bsm_id">BSM ID</ThSort>
    <ThSort {handler} orderBy="position">Position</ThSort>
    <ThSort {handler} orderBy="last_login">Last Login</ThSort>

    {#if showAdminSection}
        <th>Actions</th>
    {/if}
</tr>
</thead>

<tbody>
{#each $rows as row}
    <tr>
        <td>
            <div class="flex items-center gap-2">
                <Avatar
                        src={client.files.getURL(row, row.avatar)}
                        width="w-10"
                />
                <div>
                    <span>{row.first_name} {row.last_name}</span>
                    <div>
                        <a class="font-extralight" href="mailto:{row.email}"
                        >{row.email}</a
                        >
                    </div>
                </div>
            </div>
        </td>

        <td class="flex gap-1">
            {#if row.verified}
                <div>
                    <CheckCircleSolid
                            class="text-success-600 dark:text-success-500"
                    />
                </div>
                <div>Verified</div>
            {:else}
                <div>
                    <!-- Icon for some reason not in lib -->
                    <svg
                            class="w-5 h-5 text-error-500-400-token"
                            aria-hidden="true"
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            fill="currentColor"
                            viewBox="0 0 24 24"
                    >
                        <path
                                fill-rule="evenodd"
                                d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm5.757-1a1 1 0 1 0 0 2h8.486a1 1 0 1 0 0-2H7.757Z"
                                clip-rule="evenodd"
                        />
                    </svg>
                </div>
                <div>Unverified</div>
            {/if}
        </td>

        <td>
            {#if row.number}
                    <span class="badge-icon p-3 variant-filled-primary"
                    >{row.number}</span
                    >
            {:else}
                <span>None</span>
            {/if}
        </td>

        <td>
            {#if row.bsm_id}
                {row.bsm_id}
            {:else}
                <span>None</span>
            {/if}
        </td>

        <td>{row.position}</td>

        <td>
            <LocalDate date={row.last_login}/>
        </td>

        {#if showAdminSection}
            <td class="space-x-1">
                <button class="btn btn-sm variant-ghost-primary">
                    <EditOutline/>
                    Edit
                </button>

                <button class="btn btn-sm variant-ghost-tertiary">
                    <LockOpenOutline/>
                    Make Admin
                </button>

                <button class="btn btn-sm variant-ghost-warning">
                    <EditOutline/>
                    Remove
                </button>
            </td>
        {/if}
    </tr>
{/each}

{#if $rows.length === 0}
    <tr>
        <th class="py-4" colspan="8">No team members to display.</th>
    </tr>
{/if}
</tbody>
