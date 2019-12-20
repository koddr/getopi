// Style
import style from "./style";

// Use store
import useStoreon from "storeon/preact";

// UI elements
import Input from "../../ui/input";
import Button from "../../ui/button";

const LoginForm = () => {
  const { dispatch, loginEmail } = useStoreon("loginEmail");

  return (
    <div class={style.form}>
      <Input
        id="email"
        label="Enter your e-mail"
        type="email"
        placeholder="mail@example.com"
        value={loginEmail}
        onInput={e => {
          dispatch("login/pre-save/email", e.target.value);
        }}
      />
      <Input
        id="password"
        label="Enter your password"
        type="password"
        placeholder="● ● ● ● ● ● ●"
      />
      <Button name="Login" design="secondary" onClick={() => alert("OK!")} />
    </div>
  );
};

export default LoginForm;
