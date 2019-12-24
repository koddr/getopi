// Style
import style from "./style";

const Checkbox = props => {
  return (
    <label class={style.checkbox}>
      <input
        id={props.id}
        name={props.name}
        type="checkbox"
        class={style.input}
        checked={props.checked && "checked"}
      />
      <div class={style.box}>
        <div class={style.line_1}></div>
        <div class={style.line_2}></div>
      </div>
      <div class={style.text}>{props.text}</div>
    </label>
  );
};

export default Checkbox;
