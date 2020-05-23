/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

const Header = () => (
  <header class={style.header}>
    <div class={style.logo}>
      <Link activeClassName={style.active} href="/">
        <img src="/assets/images/getopi_short-logo.svg" alt="GetOpi Logo" />
      </Link>
    </div>
    <nav>
      <Link activeClassName={style.active} href="/">
        <img src="/assets/icons/search.svg" alt="search icon" />
      </Link>
    </nav>
  </header>
);

export default Header;
