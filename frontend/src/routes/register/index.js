import { h } from "preact";
import StoreContext from "storeon/preact/context";

// Init store
import RegisterStore from "../../storages/register";

// UI components
import RegisterForm from "../../components/forms/register";

// Style
import style from "./style";

const Register = () => {
  return (
    <StoreContext.Provider value={RegisterStore}>
      <div class={style.home}>
        <h2>Register</h2>
        <p>This is the Home component.</p>
        <RegisterForm />
      </div>
    </StoreContext.Provider>
  );
};

export default Register;
