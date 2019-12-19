// Style
import style from './style';

// Hooks
import { useEffect } from 'preact/hooks';

const Home = () => {
	useEffect(() => {
		document.title = 'Get opinion for your project here | getopi.';
	}, []);

	return (
		<div class={style.home}>
			<h2>Home</h2>
			<p>This is the Home component.</p>
		</div>
	);
};

export default Home;
