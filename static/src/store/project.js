// Project Store

const ProjectStore = (store) => {
  store.on("@init", () => ({ showCompletedTasks: true })); // Initial state
  store.on("show completed tasks", (state, status) => ({
    showCompletedTasks: status,
  }));
};

export default ProjectStore;
