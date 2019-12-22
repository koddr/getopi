// Style
import style from "./style";

// Hooks
import { useEffect } from "preact/hooks";

// Router link
import { Link } from "preact-router/match";

// Init store
import StoreContext from "storeon/preact/context";
import LoginStore from "../../storages/login";

// UI Component
import Separator from "../../components/ui/separator";
import ForgetPasswordForm from "../../components/forms/forget-pass";

const ForgetPassword = () => {
  useEffect(() => {
    document.title = "Forget password | getopi.";
  }, []);

  return (
    <StoreContext.Provider value={LoginStore}>
      <div class={style.forget_password}>
        <div class={style.image}></div>
        <div class={style.form}>
          <p>
            &larr;&nbsp;<Link href="/">Back to Home</Link>
          </p>
          <Separator />
          <h1>Forget Password?</h1>
          <Separator />
          <ForgetPasswordForm />
        </div>
      </div>
    </StoreContext.Provider>
  );
};

export default ForgetPassword;
