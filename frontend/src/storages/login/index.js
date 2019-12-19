import createStore from 'storeon';
import persistState from '@storeon/localstorage';

// Create store schema
let store = store => {
	store.on('@init', () => ({ loginEmail: '' }));
	store.on('login/pre-save/email', (state, email) => ({ loginEmail: email }));
};

// Init /login store
const LoginStore = createStore([store, persistState(['loginEmail'])]);

export default LoginStore;
