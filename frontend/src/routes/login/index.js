// Style
import style from "./style";

// Hooks
import { useEffect } from "preact/hooks";

// Init store
import StoreContext from "storeon/preact/context";
import LoginStore from "../../storages/login";

// Router link
import { Link } from "preact-router/match";

// UI Component
import LoginForm from "../../components/forms/login";

const Login = () => {
  useEffect(() => {
    document.title = "Login to Account | getopi.";
  }, []);

  return (
    <StoreContext.Provider value={LoginStore}>
      <div class={style.login}>
        <div class={style.image}></div>
        <div class={style.form}>
          <p>
            &larr;&nbsp;<Link href="/">Back to Home</Link>
          </p>
          <h1>Login to Account</h1>
          <p>This is the Login component.</p>
          <LoginForm />
        </div>
      </div>
    </StoreContext.Provider>
  );
};

export default Login;
