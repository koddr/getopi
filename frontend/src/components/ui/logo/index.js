import { h } from "preact";
import { Link } from "preact-router/match";

// Style
import style from "./style.css";

const Logo = () => (
  <Link class={style.logo} href="/">
    getopi<span class={style.dot}>.</span>
  </Link>
);

export default Logo;
