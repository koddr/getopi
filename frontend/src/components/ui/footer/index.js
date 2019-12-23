// Style
import style from './style';

// Router link
import { Link } from 'preact-router/match';

const Footer = () => (
	<header class={style.footer}>
		<div class={style.copyright}>
      &copy; 2019. <a href="https://1wa.co">True web artisans</a>
		</div>
		<nav class={style.menu}>
			<Link href="/login">Login</Link>
			<Link href="/register">Register</Link>
		</nav>
	</header>
);

export default Footer;
