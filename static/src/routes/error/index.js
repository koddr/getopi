/** @jsx h */

import { h } from "preact";
import { useEffect } from "preact/hooks";

// Styles
import style from "./style";

// Components
import Header from "../../components/header";
import Footer from "../../components/footer";

const Error = (props) => {
  // Save status
  const status = props.status ? props.status : 404;

  useEffect(() => {
    // Set META attributes
    document.title = `Error ${status}`;
  }, [status]);

  // Render component
  return (
    <>
      <Header />
      <main class={style.error}>
        <div class={style.wrapper}>
          <h1>Error {status}</h1>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Error;
