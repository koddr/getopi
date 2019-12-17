import createStore from "storeon";
import persistState from "@storeon/localstorage";

//
let store = store => {
  store.on("@init", () => ({ email: "", name: "" }));
  store.on("register/pre-save/email", (state, email) => ({
    email: email
  }));
  store.on("register/pre-save/name", (state, name) => ({
    name: name
  }));
};

const RegisterStore = createStore([store, persistState(["email", "name"])]);

export default RegisterStore;
