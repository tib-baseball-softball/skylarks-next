<script lang="ts">
import type { CustomAuthModel, ExpandedComment } from "$lib/model/ExpandedResponse.ts"
import { authSettings, client } from "$lib/pocketbase/index.svelte.ts"
import { DateTimeUtility } from "$lib/service/DateTimeUtility.ts"
import DeleteButton from "$lib/components/utility/DeleteButton.svelte"
import { invalidate } from "$app/navigation"
import { Send, SquarePen, X } from "lucide-svelte"
import { toastController } from "$lib/service/ToastController.svelte.ts"
import type { ClubsResponse } from "$lib/model/pb-types.ts"
import Avatar from "$lib/components/utility/Avatar.svelte"

interface Props {
  comment: ExpandedComment
  club: ClubsResponse
}

let { comment, club }: Props = $props()
let formText = $derived(comment.text)
let isEditing = $state(false)

const userFullName = $derived(
  comment?.expand?.user?.first_name + " " + comment?.expand?.user?.last_name
)
const authRecord = $derived(authSettings.record as CustomAuthModel)
const isLoggedInUser = $derived(authRecord.id === comment?.expand?.user?.id)

async function deleteComment(id: string) {
  const result = await client.collection("comments").delete(id)

  if (result) {
    await invalidate("comments:list")
  }
}

async function updateComment(event: Event) {
  event.preventDefault()

  try {
    const result = await client.collection("comments").update(comment.id, { text: formText })

    if (result) {
      isEditing = false
      await invalidate("comments:list")
    }
  } catch (e) {
    toastController.trigger({
      message: "Error updating comment",
      background: "preset-filled-error-500",
    })
  }
}
</script>

{#snippet editForm()}
  <form class="mt-1">
    <label class="label mb-2 sr-only" for="edit-comment-input-{comment.id}">Comment</label>
    <div class="input-group grid-cols-[auto_1fr_auto]">
      <button type="button" class="ig-btn preset-filled" title="cancel edit" onclick="{() => isEditing = false}">
        <X/>
      </button>
      <input
              id="edit-comment-input-{comment.id}"
              data-testid="edit-comment-input"
              class="ig-input"
              type="text"
              placeholder="Your comment..."
              bind:value={formText}
      />
      <button type="submit" class="ig-btn preset-filled" title="Add comment" onclick={updateComment}
              disabled={!formText}>
        <Send size={18}/>
      </button>
    </div>
  </form>
{/snippet}

<div class="avatar-container flex flex-col justify-between gap-2">
  <Avatar
          --size="2.5rem"
          background="preset-tonal-primary"
          fallback={`${comment?.expand?.user?.first_name.charAt(0)?.toUpperCase()}${comment?.expand?.user?.last_name.charAt(0)?.toUpperCase()}`}
          src={client.files.getURL(comment?.expand?.user ?? {}, comment?.expand?.user?.avatar ?? "")}
  />

  {#if isLoggedInUser || club?.admins.includes(authRecord.id)}
    <DeleteButton modelName="Comment" id={comment.id} action={deleteComment}/>
  {/if}
</div>

<div class={["w-full card p-2 rounded-base", isLoggedInUser ? "preset-tonal-primary" : "preset-tonal-surface"]}
     data-testid="comment-container">
  <div class="flex justify-between items-center">
    <p class="font-bold text-wrap">{userFullName}</p>

    <div class={["text-sm font-light", isLoggedInUser ? "" : "text-gray-600 dark:text-gray-400"]}>
      {DateTimeUtility.dateTimeFormatShort.format(new Date(comment.created))}
    </div>
  </div>

  {#if isEditing}
    {@render editForm()}
  {:else}
    <div class="flex flex-wrap items-end gap-3">
      <p class="mt-1">
        {comment.text}

      </p>
      {#if isLoggedInUser}
        <button type="button" class="btn btn-icon btn-sm border" title="edit comment text"
                onclick="{() => isEditing = true}">
          <SquarePen/>
        </button>
      {/if}
    </div>
  {/if}
</div>
