// Style
import style from "./style";

const Input = props => (
  <div class={style.field}>
    <p>
      <label for={props.id} class={style.label}>
        {props.label}&nbsp;&darr;
      </label>
    </p>
    <input
      id={props.id}
      class={style.input}
      type={props.type}
      placeholder={props.placeholder}
      value={props.value}
      onInput={props.onInput}
    />
  </div>
);

export default Input;
