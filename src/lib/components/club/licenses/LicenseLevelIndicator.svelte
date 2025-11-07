<script lang="ts">
interface Props {
  level: string
}

let { level }: Props = $props()

let colorClasses: string = $derived.by(() => {
  switch (level) {
    case "A":
      return "bg-red-500 text-black"
    case "B":
      return "bg-orange-500 text-black"
    case "C":
      return "bg-yellow-500 text-black"
    case "D":
      return "bg-teal-500 text-black"
    default:
      return "bg-gray-500 text-white"
  }
})
</script>

<div aria-label="License level {level}:" class="py-0.5 px-2 rounded-sm {colorClasses}">
  {level}
</div>