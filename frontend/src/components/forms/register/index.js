import { h } from "preact";
import useStoreon from "storeon/preact";

// Style
import style from "./style";

const RegisterForm = () => {
  const { dispatch, email, name } = useStoreon("email", "name");

  return (
    <div class={style.form}>
      <label>Email</label>
      <input
        type="text"
        value={email}
        onInput={e => {
          dispatch("register/pre-save/email", e.target.value);
        }}
      />
      {email}
      <br />
      <label>Name</label>
      <input
        type="text"
        value={name}
        onInput={e => {
          dispatch("register/pre-save/name", e.target.value);
        }}
      />
      {name}
    </div>
  );
};

export default RegisterForm;
