// @flow

import { createStoreon } from "storeon";
import { persistState } from "@storeon/localstorage";

// Stores
import ProjectStore from "./project";

const store = createStoreon([
  ProjectStore,
  persistState(["showCompletedTasks"]),
]);

export default store;
