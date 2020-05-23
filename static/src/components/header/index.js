/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

const Header = () => (
  <header class={style.header}>
    <div class={style.logo}>
      <Link href="/">
        <img src="/assets/images/getopi_short-logo.svg" alt="getopi logo" />
      </Link>
    </div>
    <nav>
      <Link href="/">
        <img src="/assets/icons/search.svg" alt="search icon" />
      </Link>
      <Link href="/">
        <img src="/assets/icons/plus-circle.svg" alt="plus circle icon" />
      </Link>
      <Link href="/">
        <img src="/assets/icons/bell.svg" alt="bell icon" />
      </Link>
      <Link href="/">
        <img src="/assets/icons/settings.svg" alt="settings icon" />
      </Link>
      <Link href="/">
        <img src="/assets/icons/no-avatar.svg" alt="no avatar icon" />
      </Link>
    </nav>
  </header>
);

export default Header;
