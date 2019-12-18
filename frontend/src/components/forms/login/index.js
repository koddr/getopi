import { h } from "preact";
import useStoreon from "storeon/preact";

// Style
import style from "./style";

// UI elements
import Input from "../../ui/input";

const LoginForm = () => {
  const { dispatch, loginEmail } = useStoreon("loginEmail");

  return (
    <div class={style.form}>
      <Input
        id="email"
        label="E-mail"
        type="email"
        placeholder="E-mail"
        value={loginEmail}
        onInput={e => {
          dispatch("login/pre-save/email", e.target.value);
        }}
      />
    </div>
  );
};

export default LoginForm;
