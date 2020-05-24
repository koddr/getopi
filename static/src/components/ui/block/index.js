/** @jsx h */

import { h } from "preact";
import style from "./style";

const Block = (props) => {
  let blockType, blockContent;
  switch (props.type) {
    case "title":
      blockType = style.block_title;
      blockContent = props.content;
      break;
    case "text":
      blockType = style.block_text;
      blockContent = props.content;
      break;
    case "reward":
      blockType = style.block_reward;
      blockContent = props.content;
      break;
    default:
      blockContent = props.children; // throw all children elements into <Block/>
      break;
  }
  return (
    <div class={style.block}>
      <div class={style.block_label}>{props.label}</div>
      <div class={blockType}>{blockContent}</div>
    </div>
  );
};

export default Block;
