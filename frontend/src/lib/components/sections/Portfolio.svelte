<script lang="ts">
    import ProjectCard from '$lib/components/ProjectCard.svelte';
    let { projects }: { projects: any[] } = $props();
</script>

<section id="portfolio" class="pt-36 pb-16 bg-slate-100">
    <div class="container">
        <div class="w-full px-4">
            <div class="max-w-xl mx-auto text-center mb-16">
                <h4 class="font-bold text-3xl text-primary mt-4">Portfolio</h4>
                <h2 class="font-semibold text-dark text-2xl mb-2">My Project</h2>
                <p class="font-medium text-md text-secondary">Here's some of My Projects!</p>
            </div>
        </div>

        <div class="w-full px-4 flex flex-wrap justify-center xl:w-10/12 xl:mx-auto">
            {#if projects && projects.length > 0}
                {#each projects as project (project.id)}
                    <ProjectCard {project} />
                {/each}
            {:else}
                <p class="font-semibold text-dark text-lg mb-5">Projects are coming soon. Stay tuned!</p>
            {/if}
        </div>
    </div>
</section>