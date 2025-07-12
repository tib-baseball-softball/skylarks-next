<script lang="ts">
  import type {CustomAuthModel, ExpandedComment} from "$lib/model/ExpandedResponse.ts";
  import CommentRow from "$lib/components/comments/CommentRow.svelte";
  import {authSettings, client} from "$lib/pocketbase/index.svelte.ts";
  import {Send} from "lucide-svelte";
  import type {CommentsCreate, CommentsResponse} from "$lib/model/pb-types.ts";
  import {toastController} from "$lib/service/ToastController.svelte.ts";
  import {invalidateAll} from "$app/navigation";

  interface Props {
    targetID: string;
    targetType: "announcement" | "event";
    comments: ExpandedComment[];
  }

  let {targetID, targetType, comments}: Props = $props();

  let commentText = $state("")

  const authRecord = $derived(authSettings.record as CustomAuthModel);

  async function addComment(event: Event) {
    event.preventDefault();

    let formValues: CommentsCreate = {
      text: commentText,
      user: authRecord.id,
    }

    if (targetType === "event") {
      formValues.event = targetID
    } else if (targetType === "announcement") {
      formValues.announcement = targetID
    }

    try {
      await client.collection("comments").create<CommentsResponse>(formValues)
    } catch (error) {
      toastController.trigger({
        message: "Failed to add comment.",
        background: "preset-filled-error-500",
      })
    }

    commentText = ""
    await invalidateAll()
  }
</script>

{#if comments.length > 0}
  <ul>
    {#each comments as comment}
      <li class={["my-3 md:my-4 flex gap-2", authRecord.id === comment?.expand?.user?.id && "flex-row-reverse"]}>
        <CommentRow {comment}/>
      </li>
    {/each}
  </ul>
{:else}
  <p>No comments yet.</p>
{/if}

<form class="mt-6">
  <div class="input-group grid-cols-[1fr_auto]">
    <input
            class="ig-input rounded-s-base"
            type="text"
            placeholder="Your comment..."
            bind:value={commentText}
    />
    <button type="submit" class="ig-btn preset-filled" title="Add comment" onclick={addComment} disabled={!commentText}>
      <Send size={18}/>
    </button>
  </div>
</form>
