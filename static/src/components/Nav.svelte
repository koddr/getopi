<script>
  export let segment;
</script>

<style>
  .nav {
    background-color: var(--white);
    box-shadow: var(--box-shadow);
    padding: 12px 8px;
    position: sticky;
    top: 0;
  }

  @media (min-width: 768px) {
    .nav {
      padding: 16px 24px;
    }
  }

  .nav__wrapper {
    display: grid;
    grid-template-columns: min-content min-content 1fr;
    column-gap: 6px;
    max-width: 1440px;
    margin: 0 auto;
  }

  @media (min-width: 768px) {
    .nav__wrapper {
      column-gap: 16px;
    }
  }

  .nav__logo {
    border-bottom: 0;
    align-self: center;
  }

  .nav__logo img {
    width: 32px;
  }

  @media (min-width: 768px) {
    .nav__logo img {
      width: 48px;
    }
  }

  .nav__title {
    font-size: calc(var(--font-size) + 3px);
    line-height: calc(var(--font-size) + 3px);
    font-weight: 500;
    color: var(--black);
    align-self: center;
  }

  @media (min-width: 768px) {
    .nav__title {
      font-size: calc(var(--font-size) + 12px);
      line-height: calc(var(--font-size) + 12px);
    }
  }

  .nav__list {
    display: grid;
    grid-template-columns: repeat(3, min-content);
    column-gap: 20px;
    justify-self: end;
  }

  .nav__item {
    align-self: center;
  }

  .nav__icon {
    width: 26px;
  }

  .nav__avatar {
    width: 32px;
    border: 2px solid var(--blue);
    border-radius: 50%;
  }

  .nav__link {
    display: grid;
    grid-template-columns: repeat(2, min-content);
    column-gap: 8px;
    border-bottom: 0;
    align-items: center;
  }

  .nav__link:visited {
    color: var(--blue);
  }

  @media (max-width: 480px) {
    .nav__link {
      display: block;
    }

    .nav__body {
      display: none;
    }
  }
</style>

<nav class="nav" aria-label="Main">
  <div class="nav__wrapper">
    <a
      class="nav__logo"
      aria-current={segment === undefined ? 'page' : undefined}
      href=".">
      <img src="images/getopi_short-logo.svg" alt="getopi logo" />
    </a>
    <div class="nav__title">GetOpi</div>
    <ul class="nav__list">
      <li class="nav__item">
        <a
          class="nav__link"
          aria-current={segment === 'about' ? 'page' : undefined}
          href="about">
          <img class="nav__icon" src="icons/search.svg" alt="search icon" />
          <span>Search</span>
        </a>
      </li>
      <li class="nav__item">
        <a
          class="nav__link"
          aria-current={segment === 'about' ? 'page' : undefined}
          href="about">
          <img class="nav__icon" src="icons/bell.svg" alt="bell icon" />
          <span class="nav__body">Notifications</span>
        </a>
      </li>
      <li class="nav__item">
        <a
          rel="prefetch"
          class="nav__link"
          aria-current={segment === 'projects' ? 'page' : undefined}
          href="projects">
          <img
            class="nav__avatar"
            src="icons/no-avatar.svg"
            alt="no avatar icon" />
          <span>@koddr</span>
        </a>
      </li>
    </ul>
  </div>
</nav>
