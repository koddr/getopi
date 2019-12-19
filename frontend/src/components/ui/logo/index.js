// Style
import style from './style';

// Router link
import { Link } from 'preact-router/match';

const Logo = () => (
	<Link class={style.logo} href="/">
    getopi<span class={style.dot}>.</span>
	</Link>
);

export default Logo;
