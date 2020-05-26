/** @jsx h */

import { h } from "preact";
import { useState } from "preact/hooks";
import style from "./style";

const Checkbox = (props) => {
  // Connect to local state
  const [checkboxState, setCheckboxState] = useState({
    checked: props.default_state,
  });

  // Define style
  let checkboxType;
  switch (props.type) {
    case "success":
      checkboxType = style.success;
      break;
    case "warning":
      checkboxType = style.warning;
      break;
    case "danger":
      checkboxType = style.danger;
      break;
  }

  // Toggle checked state
  const toggleCheckboxState = () => {
    let checked = !checkboxState.checked;
    setCheckboxState({ checked });
    if (props.callback) props.callback();
  };

  // Render checkbox
  return (
    <div
      onClick={toggleCheckboxState}
      class={[
        style.checkbox,
        checkboxType,
        checkboxState.checked ? style.checkbox_checked : null,
      ].join(" ")}
    >
      <label class={style.checkbox_label}>
        <input
          type="checkbox"
          name={props.name}
          class={style.checkbox_input}
          checked={checkboxState.checked}
        />
        <div onClick={toggleCheckboxState} class={style.checkbox_circle} />
      </label>
      <div class={style.checkbox_text}>{props.text}</div>
    </div>
  );
};

export default Checkbox;
