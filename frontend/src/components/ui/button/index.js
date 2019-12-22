// Style
import style from "./style";

// Router link
import { Link } from "preact-router/match";

export const Button = props => {
  let color, fullwidth, outline;

  switch (props.color) {
    case "red":
      color = style.red;
      break;
    case "green":
      color = style.green;
      break;
    default:
      color = style.blue;
      break;
  }

  if (props.fullwidth === true) fullwidth = style.fullwidth;
  if (props.outline === true) outline = style.outline;

  return (
    <button
      class={`${style.button} ${color} ${fullwidth || ""} ${outline || ""}`}
      onClick={props.onClick}
    >
      <span class={style.name}>{props.name}</span>
      {props.icon ? <span class={style.icon}>{props.icon}</span> : ""}
    </button>
  );
};

export const ButtonLink = props => {
  let color, fullwidth, outline;

  switch (props.color) {
    case "red":
      color = style.red;
      break;
    case "green":
      color = style.green;
      break;
    default:
      color = style.blue;
      break;
  }

  if (props.fullwidth === true) fullwidth = style.fullwidth;
  if (props.outline === true) outline = style.outline;

  return (
    <Link
      class={`${style.button} ${color} ${fullwidth || ""} ${outline || ""}`}
      href={props.href}
      redirect={props.redirect}
    >
      <span class={style.name}>{props.name}</span>
      {props.icon ? <span class={style.icon}>{props.icon}</span> : ""}
    </Link>
  );
};
