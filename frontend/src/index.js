// @flow
import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { ApolloProvider } from "react-apollo";
import { ApolloClient } from "apollo-client";
import { InMemoryCache } from "apollo-cache-inmemory";
import { BatchHttpLink } from "apollo-link-batch-http";
import { onError } from "apollo-link-error";
import { ApolloLink } from "apollo-link";

// Utils
import registerServiceWorker from "./registerServiceWorker";
import WebFont from "webfontloader";
import * as Stickyfill from "stickyfilljs";

// Styles
import "./style.css";

// Components
import Landing from "./components/Landing";

// Checking root app element
const root = document.getElementById("getopi-app");
if (root === null) throw new Error("Error! No root element.");

// Apollo Client
const client = new ApolloClient({
  link: ApolloLink.from([
    onError(({ graphQLErrors, networkError }) => {
      if (graphQLErrors)
        graphQLErrors.map(({ message, locations, path }) =>
          console.log(
            `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
          )
        );
      if (networkError) console.log(`[Network error]: ${networkError}`);
    }),
    new BatchHttpLink({
      uri: "http://localhost:8000/__gql__/",
      credentials: "same-origin"
    })
  ]),
  cache: new InMemoryCache()
});

// Render App
ReactDOM.render(
  <ApolloProvider client={client}>
    <Router>
      <Switch>
        <Route exact path="/" component={Landing} />
      </Switch>
    </Router>
  </ApolloProvider>,
  root
);

// Web font loader for Google Fonts
WebFont.load({
  google: {
    families: ["Open+Sans:400,800"]
  }
});

// Make sticky box
if (document.querySelectorAll(".sticky").length > 0) {
  Stickyfill.add(document.querySelectorAll(".sticky"));
}

// Register service worker
registerServiceWorker();
