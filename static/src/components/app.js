/** @jsx h */

import { h, Component } from "preact";
import { Router } from "preact-router";
import AsyncRoute from "preact-async-route";
import { StoreContext } from "storeon/preact";

// Store
import store from "../store";

// Components
import Loader from "./ui/loader";

export default class App extends Component {
  render() {
    return (
      <div id="getopi_app">
        <StoreContext.Provider value={store}>
          <Router onChange={(e) => (this.currentUrl = e.url)}>
            <AsyncRoute
              path="/project/:alias"
              getComponent={() =>
                import("../routes/project").then((module) => module.default)
              }
              loading={() => <Loader />}
            />
            <AsyncRoute
              default
              path="/error/:status?"
              getComponent={() =>
                import("../routes/error").then((module) => module.default)
              }
              loading={() => <Loader />}
            />
          </Router>
        </StoreContext.Provider>
      </div>
    );
  }
}
