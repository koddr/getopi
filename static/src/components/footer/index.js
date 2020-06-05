/** @jsx h */

import { h } from "preact";
import { Link } from "preact-router/match";
import style from "./style";

const Footer = () => (
  <>
    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 100 19">
      <path fill="#90B7F9" d="M0 30 V15 Q30 3 60 15 V30z" />
      <path fill="#216EF4" d="M0 30 V12 Q30 17 55 12 T100 11 V30z" />
    </svg>
    <footer class={style.footer}>
      <div class={style.footer__wrapper}>
        <div class={style.footer__copyright}>
          &copy; 2020. Crafted with &hearts; to people and robots by{" "}
          <a href="https://1wa.co" target="_blank" rel="noreferrer">
            True web artisans
          </a>
          .
        </div>
        <div class={style.footer__logo}>
          <Link href="/">
            <img
              src="/assets/images/getopi_full-logo.svg"
              alt="getopi full logo"
            />
          </Link>
        </div>
        <nav class={style.footer__nav}>
          <div>
            <Link href="/">About Service</Link>
          </div>
          <div>
            <Link href="/">Private Policy</Link>
          </div>
          <div>
            <Link href="/">Contact Us</Link>
          </div>
        </nav>
      </div>
    </footer>
  </>
);

export default Footer;
