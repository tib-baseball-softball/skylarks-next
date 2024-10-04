<script>
    import { HeartSolid } from "flowbite-svelte-icons";
    import Svelte from "../icons/Svelte.svelte";

</script>
<section class="flex justify-center items-center gap-8 mt-5">
    <a href="https://www.facebook.com/TiBBaseball/" target="_blank" rel="noreferrer" aria-label="to Skylarks Facebook page">
        <img src="/Facebook_Logo_Primary.svg" alt="Facebook brand logo" width="50em">
    </a>

    <a href="https://www.instagram.com/berlinskylarks/" target="_blank" rel="noreferrer" aria-label="to Skylarks Instagram profile">
        <img class="m-1" src="/Instagram_Glyph_Gradient.png" alt="Instagram brand logo" width="40em">
    </a>

    <a href="https://tib1848ev.de/" target="_blank" aria-label="to Turngemeinde in Berlin main website">
        <img class="min-w-8 max-w-14" src="/tib_logo.svg" alt="TiB Logo" width="38em">
    </a>

</section>

<nav class="list-nav pt-1 pb-3 px-6">

    <ul class="flex items-end justify-center gap-4">
        <li><a href="/legalnotice">Impressum</a></li>
        <li><a href="/privacy">Datenschutzerklärung</a></li>
    </ul>

</nav>

<div class="flex justify-end items-center gap-2 pb-3 mx-4">
    <span>Built with</span>
    <HeartSolid size="lg"/>
    and
    <Svelte fillColor={"currentColor"} classes="w-6"/>
</div>