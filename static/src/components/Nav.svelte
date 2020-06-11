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

  @media (width >= 768px) {
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

  @media (width >= 768px) {
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

  @media (width >= 768px) {
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

  @media (width >= 768px) {
    .nav__title {
      font-size: calc(var(--font-size) + 9px);
      line-height: calc(var(--font-size) + 9px);
    }
  }

  .menu__list {
    display: grid;
    grid-template-columns: repeat(3, min-content);
    column-gap: 20px;
    justify-self: end;
  }

  .menu__item {
    align-self: center;
  }

  .menu__item .item__icon {
    width: 26px;
  }

  .menu__item .item__avatar {
    width: 32px;
    border: 2px solid var(--blue);
    border-radius: 50%;
  }

  .menu__item a {
    display: grid;
    grid-template-columns: repeat(2, min-content);
    column-gap: 8px;
    border-bottom: 0;
    align-items: center;
  }

  .menu__item a:visited {
    color: var(--blue);
  }

  @media (max-width: 480px) {
    .menu__item a {
      display: block;
    }

    .menu__item a span {
      display: none;
    }
  }
</style>

<nav class="nav">
  <div class="nav__wrapper">
    <a
      class="nav__logo"
      aria-current={segment === undefined ? 'page' : undefined}
      href=".">
      <img src="images/getopi_short-logo.svg" alt="getopi logo" />
    </a>
    <div class="nav__title">GetOpi</div>
    <ul class="menu__list">
      <li class="menu__item">
        <a aria-current={segment === 'about' ? 'page' : undefined} href="about">
          <img class="item__icon" src="icons/search.svg" alt="search icon" />
          <span>Search</span>
        </a>
      </li>
      <li class="menu__item">
        <a aria-current={segment === 'about' ? 'page' : undefined} href="about">
          <img class="item__icon" src="icons/bell.svg" alt="bell icon" />
          <span>Notifications</span>
        </a>
      </li>
      <li class="menu__item">
        <a
          rel="prefetch"
          aria-current={segment === 'projects' ? 'page' : undefined}
          href="projects">
          <img
            class="item__avatar"
            src="icons/no-avatar.svg"
            alt="no avatar icon" />
          <span>@koddr</span>
        </a>
      </li>
    </ul>
  </div>
</nav>
