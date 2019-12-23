// Style
import style from './style';

// Hooks
import { useEffect } from 'preact/hooks';

// Init store
import StoreContext from 'storeon/preact/context';
import LoginStore from '../../storages/login';

// Router link
import { Link } from 'preact-router/match';

// UI Component
import LoginForm from '../../components/forms/login';
import Separator from '../../components/ui/separator';
import { ButtonLink } from '../../components/ui/button';

const Login = () => {
	useEffect(() => {
		document.title = 'Login to Account | getopi.';
	}, []);

	return (
		<StoreContext.Provider value={LoginStore}>
			<div class={style.login}>
				<div class={style.image} />
				<div class={style.form}>
					<p>
            &larr;&nbsp;<Link href="/">Back to Home</Link>
					</p>
					<Separator />
					<h1>Login to Account</h1>
					<Separator />
					<LoginForm />
					<Separator text="Don't have an Account?" />
					<ButtonLink
						name="Create your own Account!"
						icon=":)"
						color="green"
						fullwidth
						outline
						href="/register"
					/>
				</div>
			</div>
		</StoreContext.Provider>
	);
};

export default Login;
