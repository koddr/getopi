import { h } from "preact";
import useStoreon from "storeon/preact";

// Style
import style from "./style";

const LoginForm = () => {
  const { dispatch, email } = useStoreon("email");

  return (
    <div class={style.form}>
      <label>Email</label>
      <input
        type="text"
        value={email}
        onInput={e => {
          dispatch("login/pre-save/email", e.target.value);
        }}
      />
      {email}
    </div>
  );
};

export default LoginForm;
