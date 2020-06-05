/** @jsx h */

import { h } from "preact";
import style from "./style";

const Button = (props) => {
  let buttonType;
  switch (props.type) {
    case "secondary":
      buttonType = style.button__secondary;
      break;
    case "success":
      buttonType = style.button__success;
      break;
    case "warning":
      buttonType = style.button__warning;
      break;
    case "danger":
      buttonType = style.button__danger;
      break;
  }
  return <button class={buttonType}>{props.name}</button>;
};

export default Button;
