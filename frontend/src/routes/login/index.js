// Style
import style from './style';

// Hooks
import { useEffect } from 'preact/hooks';

// Init store
import StoreContext from 'storeon/preact/context';
import LoginStore from '../../storages/login';

// UI Component
import LoginForm from '../../components/forms/login';

const Login = () => {
	useEffect(() => {
		document.title = 'Login to Account | getopi.';
	}, []);

	return (
		<StoreContext.Provider value={LoginStore}>
			<div class={style.login}>
				<h2>Login</h2>
				<p>This is the Login component.</p>
				<LoginForm />
			</div>
		</StoreContext.Provider>
	);
};

export default Login;
