<script lang="ts">
  import type {PageProps} from "./$types";
  import {User} from "lucide-svelte";
  import PriorityBadge from "$lib/components/announcements/PriorityBadge.svelte";
  import type {CustomAuthModel} from "$lib/model/ExpandedResponse";
  import {authSettings, client} from "$lib/pocketbase/index.svelte";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";

  let {data}: PageProps = $props();
  const announcement = $derived(data.announcement);

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  function deleteAction(id: string) {
    client.collection("announcements").delete(id);
  }
</script>

<div class="flex justify-between items-center gap-4 lg:gap-6">
  <h1 class="h1 my-3!">{$announcement.title}</h1>
  <PriorityBadge priority={$announcement.priority}/>
</div>

<div class="flex justify-center">
  <div>
    <section>
      <p class="my-6 flex gap-3 items-center">
        <User size="28"/>
        <span class="h6">
        {$announcement.expand?.author.first_name}
          {$announcement.expand?.author.last_name}
      </span>
      </p>
    </section>

    <section class="prose text-justify">
      {@html $announcement.bodytext}

    </section>

    {#if $announcement.link}
      <section class="my-4">
        <header>
          <h2 class="h4">Link</h2>
        </header>

        <a class="anchor block mt-2" href={$announcement.link} target="_blank"
           title="Link associated with this announcement: {$announcement.link}">
          {#if $announcement.link_text}
            {$announcement.link_text}
          {:else }
            {$announcement.link}
          {/if}
        </a>
      </section>
    {/if}

    <section class="my-4">

      <header>
        <h2 class="h4">Admin Options</h2>
      </header>
      {#if $announcement?.expand?.club?.admins.includes(authRecord.id)}
        <div class="flex mt-3">
          <DeleteButton
                  id={$announcement.id}
                  modelName="Announcement"
                  action={deleteAction}
          />
        </div>
      {/if}
    </section>
  </div>
</div>
