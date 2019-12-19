import createStore from 'storeon';
import persistState from '@storeon/localstorage';

// Create store schema
let store = store => {
	store.on('@init', () => ({ registerEmail: '', registerName: '' }));
	store.on('register/pre-save/email', (state, email) => ({
		registerEmail: email
	}));
	store.on('register/pre-save/name', (state, name) => ({
		registerName: name
	}));
};

// Init /register store
const RegisterStore = createStore([
	store,
	persistState(['registerEmail', 'registerName'])
]);

export default RegisterStore;
