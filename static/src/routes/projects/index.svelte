<script context="module">
  export async function preload({ params, query }) {
    const res = await this.fetch(
      `http://192.168.88.100:5000/api/public/projects`
    );
    const data = await res.json();

    if (res.status === 200 && !data.error) {
      return {
        count: data.count,
        projects: data.projects,
      };
    }

    this.error(res.status, data.msg);
  }
</script>

<script>
  export let count;
  export let projects;
</script>

<style>

</style>

<svelte:head>
  <title>Projects</title>
  <meta name="keywords" content="" />
  <meta name="description" content="" />
</svelte:head>

<h1>Projects</h1>

<strong>{count}</strong>

<ul>
  {#each projects as project (project.id)}
    <li>
      <a href={`/projects/${project.alias}`}>{project.project_attrs.title}</a>
    </li>
  {/each}
</ul>
