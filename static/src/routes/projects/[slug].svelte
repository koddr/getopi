<script context="module">
  import fetchWithAuth from "../../utils/jwt";

  export async function preload({ params }) {
    // The `slug` parameter is available because this file is called [slug].svelte
    const res = await fetchWithAuth(
      `http://192.168.88.100:5000/api/public/projects/${params.slug}`,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    const data = await res.json();

    if (res.status === 200 && !res.error) {
      return {
        project: data.project.project_attrs,
        author: data.author,
      };
    }

    this.error(res.status, data.msg);
  }
</script>

<script>
  export let project;
  export let author;
</script>

<style>

</style>

<svelte:head>
  <title>{project.title}</title>
</svelte:head>

<div>
  <h1>{project.title}</h1>
  <div>{author.first_name}</div>
</div>
