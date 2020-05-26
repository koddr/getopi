/** @jsx h */

import { h } from "preact";
import style from "./style";

const Button = (props) => {
  let buttonType;
  switch (props.type) {
    case "secondary":
      buttonType = style.secondary;
      break;
    case "success":
      buttonType = style.success;
      break;
    case "warning":
      buttonType = style.warning;
      break;
    case "danger":
      buttonType = style.danger;
      break;
  }
  return <button class={buttonType}>{props.name}</button>;
};

export default Button;
