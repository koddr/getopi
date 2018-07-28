// @flow
import React from "react";
import { Query } from "react-apollo";
import gql from "graphql-tag";

// Components
import Header from "./ui/Header";

// GraphQL query
const query = gql`
  {
    allProjects {
      id
      name
    }
  }
`;

export default function Landing() {
  return (
    <React.Fragment>
      <Header />
      <section className="container">
        <div className="item">
          <h1 className="heading__blue">Get opinion here.</h1>
          <h2 className="heading__grey">Online focus group for your project</h2>
        </div>
        <div className="item">
          <Query query={query}>
            {({ data, error, loading }) => {
              if (error) return "Ooops...";
              if (loading) return "Loading...";

              return (
                <React.Fragment>
                  Project count: {data.allProjects.length}
                </React.Fragment>
              );
            }}
          </Query>
        </div>
      </section>
    </React.Fragment>
  );
}
