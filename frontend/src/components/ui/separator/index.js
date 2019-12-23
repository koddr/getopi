// Style
import style from './style';

const Separator = props => (
	<div class={style.separator}>
		{props.text ? <div class={style.text}>{props.text}</div> : ''}
		{props.text ? <div class={style.line} /> : ''}
	</div>
);

export default Separator;
