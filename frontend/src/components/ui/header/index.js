import { h } from "preact";
import { Link } from "preact-router/match";

// Style
import style from "./style.css";

// UI elements
import Logo from "../logo";

const Header = () => (
  <header class={style.header}>
    <Logo />
    <nav class={style.menu}>
      <Link activeClassName={style.active} href="/login">
        Login
      </Link>
      |
      <Link activeClassName={style.active} href="/register">
        Register
      </Link>
    </nav>
  </header>
);

export default Header;
