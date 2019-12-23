// Style
import style from "./style";

// Use state
import { useState } from "preact/hooks";

// Use store
import useStoreon from "storeon/preact";

// UI elements
import Input from "../../ui/input";
import { Button } from "../../ui/button";
import Notify from "../../ui/notify";

const RegisterForm = () => {
  const [showNotify, setShowNotify] = useState(false);
  const { dispatch, registerEmail, registerName } = useStoreon(
    "registerEmail",
    "registerName"
  );

  return (
    <form class={style.form}>
      {showNotify && (
        <Notify
          text="OK! This is info message!"
          onClose={() => setShowNotify(false)}
        />
      )}
      <Input
        id="email"
        label="Enter your E-mail"
        type="email"
        placeholder="e.g. mail@example.com"
        value={registerEmail}
        icon="mail"
        onInput={e => {
          dispatch("register/pre-save/email", e.target.value);
        }}
      />
      <Input
        id="name"
        label="Enter your Name"
        type="text"
        placeholder="e.g. John Smith"
        value={registerName}
        icon="contact"
        onInput={e => {
          dispatch("register/pre-save/name", e.target.value);
        }}
      />
      <div class={style.group}>
        <div class={style.item}>
          <Button
            name="Create account"
            color="green"
            icon="&rarr;"
            onClick={() => setShowNotify(true)}
          />
        </div>
      </div>
    </form>
  );
};

export default RegisterForm;
