// inferno module
import Inferno from 'inferno';

// routing modules
import { Router, Route } from 'inferno-router';
import createBrowserHistory from 'history/createBrowserHistory';

// app components
import { App, FormComponent } from './components';
// import App from './components/AppComponent';
// import FormComponent from './components/FormComponent';

if (module.hot) {
    require('inferno-devtools');
}

const browserHistory = createBrowserHistory();

const routes = (
	<Router history={ browserHistory }>
		<Route component={ App }>
			<Route path="/" component={ FormComponent } />
		</Route>
	</Router>
);

Inferno.render(routes, document.getElementById('app'));

if (module.hot) {
    module.hot.accept()
}

