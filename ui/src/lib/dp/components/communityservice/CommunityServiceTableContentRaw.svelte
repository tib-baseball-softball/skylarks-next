<script lang="ts">
  import { page } from "$app/state";
  import { appLocale } from "$lib/dp/locale.svelte";
  import { DateTimeUtility } from "$lib/dp/service/DateTimeUtility";
  import type { ExpandedServiceEntry } from "$lib/dp/types/ExpandedResponse";
  import { toHours } from "$lib/dp/utility/toHours";
  import type { TableHandler } from "@vincjo/datatables";
  import Dialog from "../modal/Dialog.svelte";
  import { SquarePen } from "lucide-svelte";
  import ServiceEntryForm from "../forms/ServiceEntryForm.svelte";
  import DeleteButton from "../utils/DeleteButton.svelte";
  import { client } from "$lib/dp/client.svelte";
  import { Collection } from "$lib/dp/enum/Collection";
  import { invalidate } from "$app/navigation";

  interface Props {
    handler: TableHandler<ExpandedServiceEntry>;
  }

  let { handler }: Props = $props();
  const rows = $derived(handler.rows) as ExpandedServiceEntry[];

  const season = $derived(page.url.searchParams.get("season"));
  
  async function deleteEntry(id: string) {
    const result = await client.collection(Collection.ServiceEntries).delete(id);

    if (result) {
      await invalidate("communityservice:admin");
    }
  }
</script>

{#each rows as row (row.id)}
  <tr>
    <td>
      <strong>
        {DateTimeUtility.dateFormatMedium(appLocale.current).format(
          new Date(row.service_date),
        )}
      </strong>
    </td>

    <td>
      <a
        class="anchor"
        href="/account/{row?.expand?.member
          ?.id}/communityservice?season={season}"
        title="Go to community service page for user {row?.expand?.member
          ?.last_name}"
      >
        {row?.expand?.member?.first_name}
        {row?.expand?.member?.last_name}
      </a>
    </td>

    <td>
      <code class="code">
        {toHours(row.minutes)}h
      </code>
    </td>

    <td>{row.title}</td>

    <td>{row.description || "---"}</td>

    <td>{row.admin_comment || "---"}</td>

    <td
      >{row?.expand?.board_member?.first_name}
      {row?.expand?.board_member?.last_name}</td
    >
    
    <td class="button-box">
      {#if row.expand?.club}
        <Dialog triggerClasses="btn btn-sm preset-tonal-tertiary border-tertiary">

          {#snippet triggerContent()}
            <SquarePen size="14"/>
          {/snippet}

          {#snippet title()}
            <header>
              <h2 class="h4">Edit Community Service Entry "{row.title}"</h2>
            </header>
          {/snippet}

          <ServiceEntryForm serviceEntry={row} club={row?.expand?.club}/>
        </Dialog>
      {/if}

      <DeleteButton
        action={deleteEntry}
        classes="btn-sm preset-tonal-error border border-error-500"
        iconSize={14}
        id={row.id}
        modelName="Community Service Entry"
      />
    </td>
  </tr>
{/each}

<style>
  .code {
    font-size:  var(--text-sm)
  }
  
  .button-box {
    display: flex;
    align-items: center;
    gap: calc(var(--spacing) * 1)
  }
</style>
