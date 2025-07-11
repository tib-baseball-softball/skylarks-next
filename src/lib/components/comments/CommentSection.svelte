<script lang="ts">
  import type {CustomAuthModel, ExpandedComment} from "$lib/model/ExpandedResponse.ts";
  import CommentRow from "$lib/components/comments/CommentRow.svelte";
  import {authSettings} from "$lib/pocketbase/index.svelte.ts";

  interface Props {
    comments: ExpandedComment[];
  }

  let {comments}: Props = $props();

  const authRecord = $derived(authSettings.record as CustomAuthModel);
</script>

{#if comments.length > 0}
  <ul>
    {#each comments as comment}
      <li class={["my-4 flex gap-2", authRecord.id === comment?.expand?.user?.id && "flex-row-reverse"]}>
        <CommentRow {comment} />
      </li>
    {/each}
  </ul>
{:else}
  <p>No comments yet.</p>
{/if}