// Style
import style from "./style";

// Hooks
import { useEffect } from "preact/hooks";

// Init store
import StoreContext from "storeon/preact/context";
import RegisterStore from "../../storages/register";

// Router link
import { Link } from "preact-router/match";

// UI components
import RegisterForm from "../../components/forms/register";
import Separator from "../../components/ui/separator";
import { ButtonLink } from "../../components/ui/button";

const Register = () => {
  useEffect(() => {
    document.title = "Create a new Account | getopi.";
  }, []);

  return (
    <StoreContext.Provider value={RegisterStore}>
      <div class={style.register}>
        <div class={style.image}></div>
        <div class={style.form}>
          <p>
            &larr;&nbsp;<Link href="/">Back to Home</Link>
          </p>
          <Separator />
          <h1>Create a new Account</h1>
          <Separator />
          <RegisterForm />
          <Separator text="Already have an Account?" />
          <ButtonLink
            name="Login to your Account"
            icon=";)"
            outline={true}
            fullwidth={true}
            href="/login"
          />
        </div>
      </div>
    </StoreContext.Provider>
  );
};

export default Register;
