import createStore from "storeon";
import persistState from "@storeon/localstorage";

//
let store = store => {
  store.on("@init", () => ({ registerEmail: "", registerName: "" }));
  store.on("register/pre-save/email", (state, email) => ({
    registerEmail: email
  }));
  store.on("register/pre-save/name", (state, name) => ({
    registerName: name
  }));
};

const RegisterStore = createStore([
  store,
  persistState(["registerEmail", "registerName"])
]);

export default RegisterStore;
