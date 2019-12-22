// Style
import style from "./style";

const Button = props => {
  let design, fullwidth;

  switch (props.design) {
    case "secondary":
      design = style.secondary;
      break;
    case "success":
      design = style.success;
      break;
    default:
      design = style.primary;
      break;
  }

  if (props.fullwidth === true) fullwidth = style.fullwidth;

  return (
    <button
      class={`${style.button} ${design} ${fullwidth}`}
      onClick={props.onClick}
    >
      <span class={style.name}>{props.name}</span>
      {props.icon ? <span class={style.icon}>{props.icon}</span> : ""}
    </button>
  );
};

export default Button;
