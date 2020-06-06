<script>
  import Nav from "../components/Nav.svelte";
  import Footer from "../components/Footer.svelte";

  export let segment;
</script>

<style>
  .layout {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  .main {
    flex: 1;
    padding: 48px 8px;
  }

  @media (max-width: 480px) {
    .main {
      padding: 24px 8px;
    }
  }

  .main__wrapper {
    max-width: 1440px;
    margin: 0 auto;
  }
</style>

<div class="layout">
  <Nav {segment} />
  <main class="main">
    <div class="main__wrapper">
      <slot />
    </div>
  </main>
  <Footer />
</div>
