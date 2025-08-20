<script lang="ts">
    // 'project' adalah data yang akan kita kirim dari luar komponen ini
    let { project }: { project: { title: string; description: string; image_url: string; project_url: string; }} = $props();
</script>

<div class="mb-12 p-4 md:w-1/2">
    <a href={project.project_url} target="_blank" rel="noopener noreferrer">
        <div class="rounded-md shadow-md overflow-hidden">
            <img src={project.image_url} alt={project.title} class="w-full h-56 object-cover" />
        </div>
        <h3 class="font-semibold text-dark text-xl mt-5 mb-3">{project.title}</h3>
        <p class="font-medium text-base text-secondary">{project.description}</p>
    </a>
</div>