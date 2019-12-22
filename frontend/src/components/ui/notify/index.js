// Style
import style from "./style";

const Notify = props => {
  let design;

  switch (props.design) {
    case "success":
      design = style.success;
      break;
    case "error":
      design = style.error;
      break;
    default:
      design = style.info;
      break;
  }

  if (props.autoCloseTime) setTimeout(props.onClose, props.autoCloseTime);

  return (
    <div class={`${style.notify} ${design}`}>
      <div class={style.text}>{props.text}</div>
      <button class={style.close} onClick={props.onClose}>
        <div class={style.line_1}></div>
        <div class={style.line_2}></div>
      </button>
    </div>
  );
};

export default Notify;
