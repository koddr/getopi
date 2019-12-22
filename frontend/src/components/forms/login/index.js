// Style
import style from "./style";

// Use state
import { useState } from "preact/hooks";

// Use store
import useStoreon from "storeon/preact";

// Router link
import { Link } from "preact-router/match";

// UI elements
import Input from "../../ui/input";
import Button from "../../ui/button";
import Notify from "../../ui/notify";

const LoginForm = () => {
  const [showNotify, setShowNotify] = useState(false);
  const { dispatch, loginEmail } = useStoreon("loginEmail");

  return (
    <div class={style.form}>
      {showNotify && (
        <Notify
          text="OK! This is info message!"
          onClose={() => setShowNotify(false)}
        />
      )}
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
        placeholder="○ ○ ○ ○ ○"
      />
      <div class={style.group}>
        <div class={style.item}>
          <Button
            name="Go to Account"
            icon="&rarr;"
            onClick={() => setShowNotify(true)}
          />
        </div>
        <div class={(style.item, style.right)}>
          <Link href="/">Forgot password?</Link>
        </div>
      </div>
    </div>
  );
};

export default LoginForm;
