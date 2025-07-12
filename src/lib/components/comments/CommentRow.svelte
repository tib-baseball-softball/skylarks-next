<script lang="ts">
  import type {CustomAuthModel, ExpandedComment} from "$lib/model/ExpandedResponse.ts";
  import {authSettings, client} from "$lib/pocketbase/index.svelte.ts";
  import {Avatar} from "@skeletonlabs/skeleton-svelte";
  import {DateTimeUtility} from "$lib/service/DateTimeUtility.ts";
  import DeleteButton from "$lib/components/utility/DeleteButton.svelte";
  import {invalidateAll} from "$app/navigation";

  interface Props {
    comment: ExpandedComment;
  }

  let {comment}: Props = $props();

  const userFullName = $derived(comment?.expand?.user?.first_name + " " + comment?.expand?.user?.last_name);
  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const isLoggedInUser = $derived(authRecord.id === comment?.expand?.user?.id);

  async function deleteComment(id: string) {
    const result = await client.collection("comments").delete(id);

    if (result) {
      await invalidateAll()
    }
  }
</script>

<div class="avatar-container flex flex-col justify-between gap-2">
  <Avatar
          src={client.files.getURL(comment?.expand?.user ?? {}, comment?.expand?.user?.avatar ?? "")}
          name={userFullName}
          size="w-11 h-11"
          background="preset-tonal-primary"
  />

  {#if isLoggedInUser}
    <DeleteButton modelName="Comment" id={comment.id} action={deleteComment} />
  {/if}
</div>

<div class={["w-full card p-2 rounded-base", isLoggedInUser ? "preset-tonal-primary" : "preset-tonal-surface"]}>
  <div class="flex justify-between items-center">
    <p class="font-bold text-wrap">{userFullName}</p>

    <div class={["text-sm font-light", isLoggedInUser ? "" : "text-gray-600 dark:text-gray-400"]}>
      {DateTimeUtility.dateTimeFormatShort.format(new Date(comment.updated))}
    </div>
  </div>

  <p class="mt-1">
    {comment.text}
  </p>
</div>
