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
      checkboxType = style.checkbox__success;
      break;
    case "warning":
      checkboxType = style.checkbox__warning;
      break;
    case "danger":
      checkboxType = style.checkbox__danger;
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
        checkboxState.checked ? style.checkbox__checked : null,
      ].join(" ")}
    >
      <label class={style.checkbox__label}>
        <input
          type="checkbox"
          name={props.name}
          class={style.checkbox__input}
          checked={checkboxState.checked}
        />
        <div onClick={toggleCheckboxState} class={style.checkbox__circle} />
      </label>
      <div class={style.checkbox__text}>{props.text}</div>
    </div>
  );
};

export default Checkbox;
