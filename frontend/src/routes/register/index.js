import { h } from "preact";
import StoreContext from "storeon/preact/context";

// Style
import style from "./style";

// Init store
import RegisterStore from "../../storages/register";

// UI components
import RegisterForm from "../../components/forms/register";

const Register = () => {
  return (
    <StoreContext.Provider value={RegisterStore}>
      <div class={style.register}>
        <h2>Register</h2>
        <p>This is the Register component.</p>
        <RegisterForm />
      </div>
    </StoreContext.Provider>
  );
};

export default Register;
