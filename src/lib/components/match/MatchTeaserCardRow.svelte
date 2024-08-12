<script lang="ts">
    export let logo: string
    export let teamName: string
    export let score: number | undefined
    export let subduedScore: boolean
</script>

<div class="flex justify-between items-center">
    <img src="{logo}" alt="team logo for {teamName}" width="50" loading="lazy"/>
    <span>{teamName}</span>
    <span class:text-surface-600-300-token={subduedScore} class="text-xl font-extrabold ">{score ?? " "}</span>
</div>