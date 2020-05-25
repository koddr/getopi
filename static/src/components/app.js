/** @jsx h */

import { h, Component } from "preact";
import { Router } from "preact-router";
import AsyncRoute from "preact-async-route";

// Components
import Loader from "./ui/loader";

export default class App extends Component {
  render() {
    return (
      <div id="getopi_app">
        <Router onChange={(e) => (this.currentUrl = e.url)}>
          <AsyncRoute
            path="/project/:alias"
            getComponent={() =>
              import("../routes/project").then((module) => module.default)
            }
            loading={() => <Loader />}
          />
        </Router>
      </div>
    );
  }
}
