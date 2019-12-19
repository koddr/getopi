if (process.env.NODE_ENV === 'development') {
	require('preact/debug');
}

// Styles
import './styles';

// Init App
import App from './components/app';

export default App;
