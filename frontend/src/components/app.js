import { h, Component } from "preact";
import { Router } from "preact-router";

// UI elements
import Header from "./header";

// Code-splitting is automated for routes
import Home from "../routes/home";
import Login from "../routes/login";
import Register from "../routes/register";

class GetopiApp extends Component {
  /** Gets fired when the route changes.
   *	@param {Object} event		"change" event from [preact-router](http://git.io/preact-router)
   *	@param {string} event.url	The newly routed URL
   */
  handleRoute = e => {
    this.currentUrl = e.url;
  };

  render() {
    return (
      <div id="getopi-app">
        <Header />
        <Router onChange={this.handleRoute}>
          <Home path="/" />
          <Login path="/login" />
          <Register path="/register" />
        </Router>
      </div>
    );
  }
}

export default GetopiApp;