// Style
import style from './style';

// Hooks
import { useEffect } from 'preact/hooks';

// Init store
import StoreContext from 'storeon/preact/context';
import RegisterStore from '../../storages/register';

// UI components
import RegisterForm from '../../components/forms/register';

const Register = () => {
	useEffect(() => {
		document.title = 'Register a new Account | getopi.';
	}, []);

	return (
		<StoreContext.Provider value={RegisterStore}>
			<div class={style.register}>
				<h2>Register</h2>
				<p>This is the Register component.</p>
				<RegisterForm />
			</div>
		</StoreContext.Provider>
	);
};

export default Register;
