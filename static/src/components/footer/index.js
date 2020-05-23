/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

const Footer = () => (
  <footer class={style.footer}>
    <nav>
      <Link href="/">About Service</Link>
      <Link href="/">Private Policy</Link>
      <Link href="/">Contact Us</Link>
    </nav>
    <div class={style.logo}>
      <Link href="/">
        <img src="/assets/images/getopi_full-logo.svg" alt="getopi full logo" />
      </Link>
    </div>
  </footer>
);

export default Footer;
