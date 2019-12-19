// Style
import style from './style';

// Use store
import useStoreon from 'storeon/preact';

// UI elements
import Input from '../../ui/input';

const RegisterForm = () => {
	const { dispatch, registerEmail, registerName } = useStoreon(
		'registerEmail',
		'registerName'
	);

	return (
		<form class={style.form}>
			<Input
				id="email"
				label="E-mail"
				type="email"
				placeholder="E-mail"
				value={registerEmail}
				onInput={e => {
					dispatch('register/pre-save/email', e.target.value);
				}}
			/>
			<Input
				id="name"
				label="Name"
				type="text"
				placeholder="Name"
				value={registerName}
				onInput={e => {
					dispatch('register/pre-save/name', e.target.value);
				}}
			/>
		</form>
	);
};

export default RegisterForm;
