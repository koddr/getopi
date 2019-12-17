import createStore from "storeon";
import persistState from "@storeon/localstorage";

//
let store = store => {
  store.on("@init", () => ({ email: "" }));
  store.on("login/pre-save/email", (state, email) => ({ email: email }));
};

const LoginStore = createStore([store, persistState(["email"])]);

export default LoginStore;
