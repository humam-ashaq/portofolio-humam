<script lang="ts">
  // Data 'projects' akan dikirim dari +page.ts
  export let data;
</script>

<main class="container">
  <h1 class="text-4xl font-bold text-center text-blue-600">My Portfolio</h1>
  <div class="project-grid">
    {#if data.projects && data.projects.length > 0}
      {#each data.projects as project}
        <div class="card">
          <img src={project.image_url} alt={project.title} />
          <div class="card-content">
              <h2>{project.title}</h2>
              <p>{project.description}</p>
              <a href={project.project_url} target="_blank" rel="noopener noreferrer">
                View Project
              </a>
          </div>
        </div>
      {/each}
    {:else}
      <p>Loading projects or no projects found...</p>
    {/if}
  </div>
</main>

<style>
  /* Anda bisa menggunakan style dari panduan sebelumnya atau style Anda sendiri */
  .container { max-width: 960px; margin: 2rem auto; padding: 1rem; }
  h1 { text-align: center; margin-bottom: 2rem; }
  .project-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(280px, 1fr)); gap: 1.5rem; }
  .card { border: 1px solid #ddd; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 5px rgba(0,0,0,0.1); display: flex; flex-direction: column; }
  .card img { width: 100%; height: 180px; object-fit: cover; }
  .card-content { padding: 1rem; flex-grow: 1; }
  .card a { display: inline-block; margin-top: 1rem; color: #007bff; text-decoration: none; }
  .card a:hover { text-decoration: underline; }
</style>