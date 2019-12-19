// Style
import style from './style';

const Input = props => (
	<div class={style.field}>
		<input
			class={style.input}
			type={props.type}
			placeholder={props.placeholder}
			value={props.value}
			onInput={props.onInput}
		/>
		<label class={style.label}>{props.label}</label>
	</div>
);

export default Input;
