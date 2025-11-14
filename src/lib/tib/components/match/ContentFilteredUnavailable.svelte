<script lang="ts">
import { BookX } from "lucide-svelte"

interface Props {
  text?: string
}

let { text = "There is no data to display for the selected filters." }: Props = $props()
</script>

<div class="flex items-center gap-4 my-3">
  <BookX size="35"/>
  <p>{text}</p>
</div>