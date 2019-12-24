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
import { Button } from "../../ui/button";
import Notify from "../../ui/notify";

const ForgetPasswordForm = () => {
  const [showNotify, setShowNotify] = useState(false);
  const { dispatch, loginEmail } = useStoreon("loginEmail");

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
        icon="mail"
        placeholder="e.g. mail@example.com"
        value={loginEmail}
      />
      <div class={style.group}>
        <div class={style.item}>
          <Button
            name="Send recovery link"
            icon="&rarr;"
            color="red"
            onClick={() => setShowNotify(true)}
          />
        </div>
        <div class={(style.item, style.right)}>
          <Link href="/login">Cancel</Link>
        </div>
      </div>
    </form>
  );
};

export default ForgetPasswordForm;
