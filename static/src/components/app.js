/** @jsx h */

import { h, Component } from "preact";
import { Router } from "preact-router";

import Header from "./header";
import Footer from "./footer";

// Code-splitting is automated for routes
import Project from "../routes/project";

export default class App extends Component {
  /** Gets fired when the route changes.
   *	@param {Object} event		"change" event from [preact-router](http://git.io/preact-router)
   *	@param {string} event.url	The newly routed URL
   */
  handleRoute = (e) => (this.currentUrl = e.url);

  render() {
    return (
      <div id="getoppi_app">
        <Header />
        <div class="wrapper">
          <Router onChange={this.handleRoute}>
            <Project path="/project/:alias" />
          </Router>
        </div>
        <Footer />
      </div>
    );
  }
}
