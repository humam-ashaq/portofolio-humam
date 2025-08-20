<script lang="ts">
    import Hero from '$lib/components/sections/Hero.svelte';
    import About from '$lib/components/sections/About.svelte';
    import Services from '$lib/components/sections/Services.svelte';
    import Portfolio from '$lib/components/sections/Portfolio.svelte';

    // 'data' akan otomatis diterima dari file +page.ts
    export let data;
</script>

<Hero />
<About />
<Services />

<Portfolio projects={data.projects} />