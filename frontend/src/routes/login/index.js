import { h } from "preact";
import StoreContext from "storeon/preact/context";

// Init store
import LoginStore from "../../storages/login";

// UI Component
import LoginForm from "../../components/forms/login";

// Style
import style from "./style";

const Login = () => {
  return (
    <StoreContext.Provider value={LoginStore}>
      <div class={style.login}>
        <h2>Login</h2>
        <p>This is the Login component.</p>
        <LoginForm />
      </div>
    </StoreContext.Provider>
  );
};

export default Login;
