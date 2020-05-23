/** @jsx h */

import { h } from "preact";
import style from "./style";

const Button = (props) => {
  let buttonType;
  switch (props.type) {
    case "primary":
      buttonType = style.primary;
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
    default:
      buttonType = style.default;
      break;
  }
  return <button class={buttonType}>{props.name}</button>;
};

export default Button;
