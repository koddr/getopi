<style lang="postcss">
  .footer {
    background-color: var(--blue);
    color: var(--white);
    font-size: calc(var(--font-size) - 3px);
    padding: 12px 8px;
  }

  @media (width >= 768px) {
    .footer {
      padding: 16px 24px;
    }
  }

  .footer a {
    color: var(--white);
    border-bottom: 1.5px solid var(--blue-light);
  }

  .footer__wrapper {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    align-items: center;
    max-width: 1440px;
    margin: 0 auto;
  }

  @media (width <= 480px) {
    .footer__wrapper {
      grid-template-columns: auto;
      row-gap: 16px;
      text-align: center;
    }
  }

  .footer__logo {
    width: 32px;
    justify-self: center;
  }

  @media (width <= 480px) {
    .footer__logo {
      display: none;
    }
  }

  .footer__list {
    display: grid;
    grid-template-columns: repeat(3, auto);
    column-gap: 8px;
    justify-self: right;
  }

  @media (width <= 480px) {
    .footer__list {
      grid-template-columns: auto;
      row-gap: 8px;
      justify-self: center;
    }
  }
</style>

<footer class="footer">
  <div class="footer__wrapper">
    <div>
      &copy; 2020. Crafted by
      <a href="https://1wa.co">True web artisans</a>
      with &hearts; for people and robots
    </div>
    <img
      class="footer__logo"
      src="images/getopi_short-logo-white.svg"
      alt="getopi logo" />
    <ul class="footer__list">
      <li>
        <a href="about">About Service</a>
      </li>
      <li>
        <a href="about">Terms of Usage</a>
      </li>
      <li>
        <a href="about">Contact Us</a>
      </li>
    </ul>
  </div>
</footer>
