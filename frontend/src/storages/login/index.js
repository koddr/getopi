import createStore from "storeon";
import persistState from "@storeon/localstorage";

//
let store = store => {
  store.on("@init", () => ({ loginEmail: "" }));
  store.on("login/pre-save/email", (state, email) => ({ loginEmail: email }));
};

const LoginStore = createStore([store, persistState(["loginEmail"])]);

export default LoginStore;
