<script lang="ts">
  import type {CustomAuthModel, ExpandedComment} from "$lib/model/ExpandedResponse.ts";
  import {authSettings, client} from "$lib/pocketbase/index.svelte.ts";
  import {Avatar} from "@skeletonlabs/skeleton-svelte";

  interface Props {
    comment: ExpandedComment;
  }

  let {comment}: Props = $props();

  const userFullName = $derived(comment?.expand?.user?.first_name + " " + comment?.expand?.user?.last_name);
  const authRecord = $derived(authSettings.record as CustomAuthModel);
  const isLoggedInUser = $derived(authRecord.id === comment?.expand?.user?.id);
</script>

<div class="avatar-container">
  <Avatar
          src={client.files.getURL(comment?.expand?.user ?? {}, comment?.expand?.user?.avatar ?? "")}
          name={userFullName}
          size="w-11"
  />
</div>

<div class={["w-full card p-2 rounded-base", isLoggedInUser ? "preset-tonal-primary" : "preset-tonal-surface"]}>
  <div class="flex justify-between items-center">
    <p class="font-bold text-wrap">{userFullName}</p>

    <div class={["text-sm font-light", isLoggedInUser ? "" : "text-gray-600 dark:text-gray-400"]}>
      {new Date(comment.updated).toLocaleString()}
    </div>
  </div>

  <p class="mt-1">
    {comment.text}
  </p>
</div>
