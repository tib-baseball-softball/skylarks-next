<script lang="ts">
interface Props {
  streak: string
}

let { streak }: Props = $props()

const START_VALUE = 10
const MAX_VALUE = 20

/**
 * Converted from Swift code (iOS app).
 *
 * Internal logic: Losing Streak from L10 to Winning Streak W10.
 * This gets converted to a simple scale from 0 to 20.
 * Every streak longer than 10 gets subsumed (should it ever happen).
 */
function getStreak(): number {
  // We start at 10 - right in the middle if there is no other data
  let streakNumber = START_VALUE

  if (streak.includes("W")) {
    const numberMatch = streak.match(/\d+/)
    if (numberMatch) {
      const number = parseInt(numberMatch[0], 10)
      streakNumber = number + START_VALUE
    }
  } else if (streak.includes("L")) {
    const numberMatch = streak.match(/\d+/)
    if (numberMatch) {
      const number = parseInt(numberMatch[0], 10)
      streakNumber = START_VALUE - number
    }
  }

  return streakNumber
}

function getEmoji(streakNumber: number): string {
  if (streakNumber <= 0) return "🪦"
  if (streakNumber >= 1 && streakNumber <= 2) return "😖"
  if (streakNumber >= 3 && streakNumber <= 4) return "☹️"
  if (streakNumber >= 5 && streakNumber <= 6) return "🙁"
  if (streakNumber >= 7 && streakNumber <= 8) return "😕"
  if (streakNumber === 9) return "😐"
  if (streakNumber === 10) return "😶"
  if (streakNumber === 11) return "🙂"
  if (streakNumber >= 12 && streakNumber <= 14) return "😀"
  if (streakNumber >= 15 && streakNumber <= 16) return "😄"
  if (streakNumber >= 17 && streakNumber <= 19) return "🤩"
  if (streakNumber >= MAX_VALUE) return "🏆"
  return "😐"
}

const streakValue = $derived(getStreak())
const emoji = $derived(getEmoji(streakValue))
</script>

<h5 class="h5">Team Mood</h5>
<div class="flex justify-center">
  <div class="flex flex-col items-center gap-6 xl:gap-8">
    <p class="text-4xl font-bold">{streak}</p>
    <p class="text-8xl">{emoji}</p>
    <p class="font-light text-sm">Clubhouse champagne celebration or dugout fights?</p>
  </div>
</div>
