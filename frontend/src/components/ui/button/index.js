// Style
import style from "./style";

const Button = props => {
  let design;

  switch (props.design) {
    case "secondary":
      design = style.secondary;
      break;
    default:
      design = style.primary;
      break;
  }

  return (
    <button class={`${style.button} ${design}`} onClick={props.onClick}>
      {props.icon ? <span class={style.icon}>{props.icon}</span> : ""}
      <span class={style.name}>{props.name}</span>
    </button>
  );
};

export default Button;
