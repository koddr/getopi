// Style
import style from "./style";

// Router link
import { Link } from "preact-router/match";

const Footer = () => (
  <footer class={style.footer}>
    <div class={style.copyright}>
      &copy; 2020. <a href="https://1wa.co">True web artisans</a>
    </div>
    <nav class={style.menu}>
      <Link href="/login">Login</Link>
      <Link href="/register">Register</Link>
    </nav>
  </footer>
);

export default Footer;
