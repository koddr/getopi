// Style
import style from './style';

// Use store
import useStoreon from 'storeon/preact';

// UI elements
import Input from '../../ui/input';

const LoginForm = () => {
	const { dispatch, loginEmail } = useStoreon('loginEmail');

	return (
		<div class={style.form}>
			<Input
				id="email"
				label="E-mail"
				type="email"
				placeholder="E-mail"
				value={loginEmail}
				onInput={e => {
					dispatch('login/pre-save/email', e.target.value);
				}}
			/>
		</div>
	);
};

export default LoginForm;
